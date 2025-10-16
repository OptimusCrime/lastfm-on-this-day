package lastfm

import (
	"encoding/json"

	"github.com/optimuscrime/lastfm-on-this-day/pgk/config"
)

type Service struct {
	requester *requester
}

func New(c *config.Config) *Service {
	return &Service{
		requester: &requester{
			config: c,
		},
	}
}

func (s *Service) Authenticate(token string) (string, error) {
	params := map[string]string{
		"token": token,
	}

	res, err := s.requester.makeLastFmRequest("auth.getSession", params)
	if err != nil {
		return "", err
	}

	var authApiResponse authApiResponse
	if err := json.Unmarshal(res, &authApiResponse); err != nil {
		return "", err
	}

	return authApiResponse.Session.Key, nil
}

func (s *Service) GetTracks(sessionKey string, date string) (*[]RecentTracksData, error) {
	duration, err := getTracksSearchDuration(date)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"sk":    sessionKey,
		"from":  duration.from,
		"to":    duration.to,
		"limit": "200",
	}

	res, err := s.requester.makeLastFmRequest("user.getRecentTracks", params)
	if err != nil {
		return nil, err
	}

	var recentTracksApiResponse recentTracksApiResponse
	if err := json.Unmarshal(res, &recentTracksApiResponse); err != nil {
		// The last.fm API is a weird. It should return an array of "recently played" tracks. If there is only one
		// track in the response, the array is instead an object. This breaks the unmarshalling of course, so we
		// have to do this dumb thing.
		var recentTrackApiResponse recentTrackApiResponse
		if innerErr := json.Unmarshal(res, &recentTrackApiResponse); innerErr != nil {
			// Return the outer error
			return nil, err
		}

		return mapRecentTracks([]recentTrackResponse{recentTrackApiResponse.RecentTracks.Track})
	}

	return mapRecentTracks(recentTracksApiResponse.RecentTracks.Track)
}

func mapRecentTracks(tracks []recentTrackResponse) (*[]RecentTracksData, error) {
	data := map[string]*RecentTracksData{}
	for _, v := range tracks {
		if v.Meta != nil && v.Meta.NowPlaying != nil && *v.Meta.NowPlaying == "true" {
			continue
		}

		id := createTrackId(v)

		if value, ok := data[id]; ok {
			value.PlayCount += 1
			continue
		}

		data[id] = &RecentTracksData{
			Url:       v.Url,
			Artist:    v.Artist.Text,
			Album:     v.Album.Text,
			Name:      v.Name,
			PlayCount: 1,
		}
	}

	resp := make([]RecentTracksData, 0, len(data))
	for _, v := range data {
		resp = append(resp, *v)
	}

	return &resp, nil
}
