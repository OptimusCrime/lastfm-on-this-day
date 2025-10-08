package lastfm

type recentTracksApiResponse struct {
	RecentTracks struct {
		Track []recentTrackResponse `json:"track"`
	} `json:"recenttracks"`
}

type recentTrackResponse struct {
	Artist struct {
		Text *string `json:"#text"`
	} `json:"artist"`

	Album struct {
		Text *string `json:"#text"`
	} `json:"album"`

	Streamable *string `json:"streamable"`
	Mbid       *string `json:"mbid"`

	Name *string `json:"name"`
	Url  *string `json:"url"`
}
