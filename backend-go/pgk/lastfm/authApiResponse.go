package lastfm

type authApiResponse struct {
	Session *sessionAuthApiResponse `json:"session"`
}

type sessionAuthApiResponse struct {
	Name       string `json:"name"`
	Key        string `json:"key"`
	Subscriber int    `json:"subscriber"`
}
