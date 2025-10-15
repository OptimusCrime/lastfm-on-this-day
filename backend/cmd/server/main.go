package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/auth"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/config"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/lastfm"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/middleware"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/token"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/tracks"
)

const port = "8113"

func main() {
	sLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(sLogger)

	c := config.CreateConfig()

	tokenService := token.New(c)
	lastfmService := lastfm.New(c)

	sLogger.Debug("boot complete")

	r := mux.NewRouter()
	r.Use(middleware.CreateCorsMiddleware)
	r.Use(middleware.CreateLoggerMiddleware(sLogger))
	r.Use(middleware.CreateAuthMiddleware(tokenService))

	auth.RegisterHandlers(r, c, lastfmService, tokenService)
	tracks.RegisterHandlers(r, lastfmService)

	sLogger.Debug("starting server on port " + port)
	http.ListenAndServe(":"+port, r)
}
