package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/lastfm"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/logger"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/render"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/resterr"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/token"
)

type Body struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	AccessToken string `json:"accessToken"`
}

type httpHandler struct {
	service *service
}

func RegisterHandlers(
	r *mux.Router,
	lastfmService *lastfm.Service,
	tokenService *token.Service,
) {
	h := &httpHandler{
		service: &service{
			lastfm: lastfmService,
			token:  tokenService,
		},
	}

	r.HandleFunc("/v1/auth", h.authenticate).Methods(http.MethodPost, http.MethodOptions)
}

func (h *httpHandler) authenticate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := logger.FromContext(ctx)

	var body Body

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		render.JSON(w, r, resterr.FromErr(err, 400))
		return
	}

	encryptedAccessToken, err := h.service.authenticate(body.Token)
	if err != nil {
		render.JSON(w, r, resterr.FromErr(err, 500))
		return
	}

	log.Debug("successfully obtained session key for user")

	render.JSON(w, r, SuccessResponse{
		AccessToken: encryptedAccessToken,
	})
}
