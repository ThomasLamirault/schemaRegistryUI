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

func ProcessSubjects(w http.ResponseWriter, path string, baseURL string, appConfig entities.AppConfig) error {
	var subject entities.Subject
	tab := strings.Split(path, "/")
	if versions, err := getSubjectsVersion(tab[2], baseURL, appConfig.Registry); err != nil {
		return err
	} else {
		subject.Name = tab[2]
		subject.Versions = GetVersionLinkList(tab[2], versions)
		if subject.HomeAddresse, err = GetHomeAddress(appConfig); err != nil {
			return err
		}
	}

	if tplate, err := template.New("subject").Parse(htmltpl.SubjectsPage); err != nil {
		return err
	} else {
		if err = tplate.Execute(w, subject); err != nil {
			return err
		}
	}
	return nil
}

func GetVersionLinkList(subjectName string, versionList []int) []entities.Versions {
	versionsTab := make([]entities.Versions, len(versionList))
	for index, version := range versionList {
		versionsTab[index].Version = version
		versionsTab[index].Link = "/subjects/" + subjectName + "/versions/" + strconv.Itoa(version)
	}
	return versionsTab
}

func getSubjectsVersion(subjectName string, baseURL string, registry entities.Registry) ([]int, error) {
	var listVersions []int
	resp, err := http.Get(baseURL + registry.Address + ":" + registry.Port + "/subjects/" + subjectName + "/versions")
	if err != nil {
		return listVersions, err
	}
	defer resp.Body.Close()
	body, err := GetResponseBodyByte(resp)
	if err != nil {
		return listVersions, err
	}
	err = json.Unmarshal(body, &listVersions)
	return listVersions, nil
}
