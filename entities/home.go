package entities

type Home struct {
	Address      string   `json:"address"`
	Port         int      `json:"port"`
	GlobalCompat string   `json:"globalCompat"`
	ListSubjects []string `json:"name"`
	HomeAddresse string   `json:"homeAddresse"`
}
