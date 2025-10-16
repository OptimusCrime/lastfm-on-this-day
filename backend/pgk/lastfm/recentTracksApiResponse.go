package lastfm

type recentTracksApiResponse struct {
	RecentTracks struct {
		Track []recentTrackResponse `json:"track"`
	} `json:"recenttracks"`
}

type recentTrackApiResponse struct {
	RecentTracks struct {
		Track recentTrackResponse `json:"track"`
	} `json:"recenttracks"`
}

type recentTrackResponse struct {
	Artist *struct {
		Text *string `json:"#text"`
	} `json:"artist"`

	Album *struct {
		Text *string `json:"#text"`
	} `json:"album"`

	Meta *struct {
		NowPlaying *string `json:"nowplaying"`
	} `json:"@attr"`

	Streamable *string `json:"streamable"`
	Mbid       *string `json:"mbid"`

	Name *string `json:"name"`
	Url  *string `json:"url"`
}
