package data

type Body struct {
	Url        string           `json:"url"`
	Javascript bool             `json:"javascript"`
	HTML       string           `json:"html"`
	Find       []map[string]any `json:"find"`
}
