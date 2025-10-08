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
	"github.com/optimuscrime/lastfm-on-this-day/pgk/tracks"
)

const port = "8113"

func main() {
	sLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	slog.SetDefault(sLogger)

	lastfmService := lastfm.New(config.CreateConfig())

	sLogger.Debug("boot complete")

	r := mux.NewRouter()
	r.Use(middleware.CreateCorsMiddleware)
	r.Use(middleware.CreateLoggerMiddleware(sLogger))
	r.Use(middleware.CreateAuthMiddleware())

	auth.RegisterHandlers(r, lastfmService)
	tracks.RegisterHandlers(r, lastfmService)

	sLogger.Debug("starting server on port " + port)
	http.ListenAndServe(":"+port, r)
}
