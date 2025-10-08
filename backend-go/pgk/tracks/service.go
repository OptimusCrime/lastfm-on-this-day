package tracks

import "github.com/optimuscrime/lastfm-on-this-day/pgk/lastfm"

type service struct {
	lastfmService *lastfm.Service
}

func (s *service) getTracks(sessionKey string, date string) (*[]lastfm.RecentTracksData, error) {
	return s.lastfmService.GetTracks(sessionKey, date)
}
