package tracks

import (
	"sort"

	"github.com/optimuscrime/lastfm-on-this-day/pgk/lastfm"
)

func mapTracks(tracks []lastfm.RecentTracksData) []Track {
	data := make([]Track, len(tracks))

	for i, track := range tracks {
		data[i] = Track{
			Url:       track.Url,
			Artist:    track.Artist,
			Album:     track.Album,
			Name:      track.Name,
			PlayCount: track.PlayCount,
		}
	}

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].PlayCount > data[j].PlayCount
	})

	return data
}
