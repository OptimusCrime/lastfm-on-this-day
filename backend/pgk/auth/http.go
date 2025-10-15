package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/config"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/lastfm"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/logger"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/render"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/resterr"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/token"
)

type PostBody struct {
	Token string `json:"token"`
}

type SuccessPostResponse struct {
	AccessToken string `json:"accessToken"`
}

type SuccessGetResponse struct {
	Url string `json:"url"`
}

type httpHandler struct {
	config  *config.Config
	service *service
}

func RegisterHandlers(
	r *mux.Router,
	c *config.Config,
	lastfmService *lastfm.Service,
	tokenService *token.Service,
) {
	h := &httpHandler{
		config: c,
		service: &service{
			lastfm: lastfmService,
			token:  tokenService,
		},
	}

	r.HandleFunc("/v1/auth", h.getAuth).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/v1/auth", h.postAuth).Methods(http.MethodPost, http.MethodOptions)
}

func (h *httpHandler) getAuth(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, SuccessGetResponse{
		Url: "https://www.last.fm/api/auth/?api_key=" + h.config.LastFmApiKey,
	})
}

func (h *httpHandler) postAuth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log := logger.FromContext(ctx)

	var body PostBody

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

	render.JSON(w, r, SuccessPostResponse{
		AccessToken: encryptedAccessToken,
	})
}
