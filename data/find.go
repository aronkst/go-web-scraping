package data

type Find struct {
	Name      string `json:"name"`
	Class     string `json:"class"`
	Attribute string `json:"attribute"`
}

type FindList struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Find  []Find `json:"find"`
}
