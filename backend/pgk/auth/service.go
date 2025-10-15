package auth

import (
	"github.com/optimuscrime/lastfm-on-this-day/pgk/lastfm"
	"github.com/optimuscrime/lastfm-on-this-day/pgk/token"
)

type service struct {
	lastfm *lastfm.Service
	token  *token.Service
}

func (s *service) authenticate(lastfmToken string) (string, error) {
	accessToken, err := s.lastfm.Authenticate(lastfmToken)
	if err != nil {
		return "", err
	}

	return s.token.EncryptToken(accessToken)
}
