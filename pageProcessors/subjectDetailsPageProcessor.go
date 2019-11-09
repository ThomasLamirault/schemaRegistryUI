package pageProcessors

import (
	"encoding/json"
	"github.com/tlamirault/schemaRegistryUi/entities"
	"github.com/tlamirault/schemaRegistryUi/htmltpl"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ProcessSubjectDetails(w http.ResponseWriter, path string, baseURL string, appConfig entities.AppConfig) error {
	var subjectVersion entities.SubjectVersion
	tab := strings.Split(path, "/")
	if id, err := strconv.Atoi(tab[4]); err != nil {
		return err
	} else {
		subjectVersion.Id = id
	}
	subjectVersion.Subject = tab[2]
	subjectVersion.SubjectAddress = "http://" + appConfig.App.Address + ":" + strconv.Itoa(appConfig.App.Port) + "/subjects/" + subjectVersion.Subject +"/"

	if subjectVersion, err := getSubjectSchema(subjectVersion, baseURL, appConfig.Registry); err != nil {
		return err
	} else {
		if tplate, err := template.New("subjectVersion").Parse(htmltpl.SubjectDetailsPage); err != nil {
			return err
		} else {
			if err = tplate.Execute(w, subjectVersion); err != nil {
				return err
			}
		}
	}
	return nil
}

func getSubjectSchema(subjectVersion entities.SubjectVersion, baseURL string, registry entities.Registry) (entities.SubjectVersion, error) {

	if resp, err := http.Get(baseURL + registry.Address + ":" + registry.Port + "/subjects/" + subjectVersion.Subject + "/versions/" + strconv.Itoa(subjectVersion.Id)); err != nil {
		return subjectVersion, err
	} else {
		defer resp.Body.Close()
		if body, err := GetResponseBodyByte(resp); err != nil {
			return subjectVersion, err
		} else {
			err = json.Unmarshal(body, &subjectVersion)
			return subjectVersion, err
		}
	}
}
