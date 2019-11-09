package entities

type Subject struct {
	Name         string     `json:"subject"`
	Versions     []Versions `json:"versions"`
	HomeAddresse string     `json:"homeAddresse"` // TODO virer ca
}

type Versions struct {
	Version int    `json:"version"`
	Link    string `json:"link"`
}
