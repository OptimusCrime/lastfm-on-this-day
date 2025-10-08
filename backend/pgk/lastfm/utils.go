package lastfm

import (
	"strconv"
	"time"
)

type tracksSearchDuration struct {
	from string
	to   string
}

func getTracksSearchDuration(dateStr string) (*tracksSearchDuration, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}

	start := t.UTC()
	end := start.AddDate(0, 0, 1).Add(-time.Second)

	return &tracksSearchDuration{
		from: strconv.FormatInt(start.Unix(), 10),
		to:   strconv.FormatInt(end.Unix(), 10),
	}, nil

}

func createTrackId(item recentTrackResponse) string {
	return getStrValue(item.Artist.Text) + "---" + getStrValue(item.Album.Text) + "---" + getStrValue(item.Name)
}

func getStrValue(val *string) string {
	if val == nil {
		return ""
	}

	return *val
}
