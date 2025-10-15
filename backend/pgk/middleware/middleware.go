package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/render"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/resterr"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/token"
)

func CreateCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json;")
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-allow-methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("access-control-allow-headers", "*")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CreateLoggerMiddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId, _ := uuid.NewUUID()

			log := logger.With("requestId", requestId.String())

			log.Debug("call to endpoint",
				"method", r.Method,
				"path", r.URL.EscapedPath(),
			)

			ctx := context.WithValue(r.Context(), "logger", log)

			next.ServeHTTP(w, r.WithContext(ctx))

			log.Debug("finished call to endpoint",
				"method", r.Method,
				"path", r.URL.EscapedPath(),
			)
		})
	}
}

func CreateAuthMiddleware(tokenService *token.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.EscapedPath()
			if path == "/v1/auth" && r.Method == http.MethodPost {
				next.ServeHTTP(w, r)
				return
			}

			authorizationValue := r.Header.Get("Authorization")
			authorizationValueSplit := strings.Split(authorizationValue, " ")

			if len(authorizationValueSplit) != 2 {
				render.JSON(w, r, resterr.FromErr(errors.New("token missing"), 403))
				return
			}

			// Because why not?
			if strings.ToLower(authorizationValueSplit[0]) != "bearer" {
				render.JSON(w, r, resterr.FromErr(errors.New("token missing"), 403))
				return
			}

			accessToken, err := tokenService.ValidateToken(authorizationValueSplit[1])

			if err != nil {
				if errors.Is(err, token.ErrInvalidConfigurationProvided) || errors.Is(err, token.ErrMissingEncryptionSubstitution) {
					render.JSON(w, r, resterr.FromErr(errors.New("incorrect configuration"), 500))
					return
				}

				if errors.Is(err, token.ErrInvalidToken) {
					render.JSON(w, r, resterr.FromErr(errors.New("incorrect token"), 403))
					return
				}
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, "sessionKey", accessToken)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
