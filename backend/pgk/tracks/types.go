package tracks

type SuccessResponse struct {
	Data []Track `json:"data"`
}

type Track struct {
	Url       *string `json:"url"`
	Artist    *string `json:"artist"`
	Album     *string `json:"album"`
	Name      *string `json:"name"`
	PlayCount int     `json:"playCount"`
}
