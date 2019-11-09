package routing

import (
	"github.com/tlamirault/schemaRegistryUi/entities"
	pageProcessor "github.com/tlamirault/schemaRegistryUi/pageProcessors"
	"net/http"
	"regexp"
)

var appConfig entities.AppConfig
var baseURL string

func Routing(baseURLParam string, appConfigParam entities.AppConfig) error {

	baseURL = baseURLParam
	appConfig = appConfigParam
	mux := http.NewServeMux()
	mux.HandleFunc("/", SchemaRegistryHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}
	return nil
}

func SchemaRegistryHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var subjectPattern = regexp.MustCompile(`/\w+/[\w\W\D]+/$`)
	var subjectDetailPattern = regexp.MustCompile(`/\w+/[\w\W\D]+/\w+/[0-9]+`)

	switch {
	case subjectPattern.MatchString(r.URL.Path):
		err = pageProcessor.ProcessSubjects(w, r.URL.Path, baseURL, appConfig)
	case subjectDetailPattern.MatchString(r.URL.Path):
		err = pageProcessor.ProcessSubjectDetails(w, r.URL.Path, baseURL, appConfig)
	default:
		err = pageProcessor.ProcessHome(w, baseURL, appConfig)
	}

	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func GetRegistryBaseURL(appConfig entities.AppConfig) string {
	if appConfig.Registry.Port == "443" {
		return "https://"
	} else {
		return "http://"
	}
}
