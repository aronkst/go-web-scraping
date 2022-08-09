package data

type Value struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ValueList struct {
	Name   string    `json:"name"`
	Values [][]Value `json:"values"`
}
