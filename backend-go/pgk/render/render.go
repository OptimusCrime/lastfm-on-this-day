package render

import (
	"encoding/json"
	"net/http"

	"github.com/optimuscrime/lastfm-on-this-day/pgk/logger"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/resterr"
)

func JSON(w http.ResponseWriter, r *http.Request, v any) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	log := logger.FromContext(ctx)

	switch t := v.(type) {
	case resterr.Resterr:
		if t.StatusCode == 500 {
			log.Error(t.Err.Error())
		} else {
			log.Debug(t.Err.Error())
		}

		w.WriteHeader(t.StatusCode)

		return
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		fail(w, r, err)
		return
	}
}

func fail(w http.ResponseWriter, r *http.Request, err error) {
	log := logger.FromContext(r.Context())

	log.Error("Failed to parse JSON", err)

	w.WriteHeader(http.StatusInternalServerError)
}
