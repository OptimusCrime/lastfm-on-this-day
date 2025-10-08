package auth

import "github.com/optimuscrime/lastfm-on-this-day/pgk/lastfm"

type service struct {
	lastfmService *lastfm.Service
}

func (s *service) authenticate(token string) (string, error) {
	return s.lastfmService.Authenticate(token)
}
