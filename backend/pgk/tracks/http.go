package tracks

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/lastfm"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/logger"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/render"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/resterr"
)

type httpHandler struct {
	service *service
}

func RegisterHandlers(
	r *mux.Router,
	lastfmService *lastfm.Service,
) {
	h := &httpHandler{
		service: &service{
			lastfmService: lastfmService,
		},
	}

	r.HandleFunc("/v1/tracks", h.getTracks).Methods(http.MethodGet, http.MethodOptions)
}

func (h *httpHandler) getTracks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sessionKey := ctx.Value("sessionKey").(string)

	date := r.URL.Query().Get("date")
	if date == "" {
		render.JSON(w, r, resterr.FromErr(errors.New("missing date"), 400))
		return
	}

	log := logger.FromContext(ctx)

	tracksData, err := h.service.getTracks(sessionKey, date)
	if err != nil {
		render.JSON(w, r, resterr.FromErr(err, 500))
		return
	}

	log.Debug("successfully returned tracks for date")

	render.JSON(w, r, SuccessResponse{
		Data: mapTracks(*tracksData),
	})
}
