package entities

type SubjectVersion struct {
	Subject        string `json:"subject"`
	Version        int    `json:"version"`
	Id             int    `json:"id"`
	Schema         string `json:"schema"`
	SubjectAddress string `json:"subjectAddress"`
}
