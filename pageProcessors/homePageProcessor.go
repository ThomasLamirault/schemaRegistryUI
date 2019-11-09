package pageProcessors

import (
	"github.com/tlamirault/schemaRegistryUi/entities"
	"github.com/tlamirault/schemaRegistryUi/htmltpl"
	"html/template"
	"net/http"
)

var tplate *template.Template

func ProcessHome(w http.ResponseWriter, baseURL string, appConfig entities.AppConfig) error {
	var home entities.Home
	var compatibility entities.Config
	var err error
	var subjects []string
	if compatibility, err = GetSchemaRegistryConfig(baseURL, appConfig.Registry); err != nil {
		return err
	}

	if subjects, err = GetSubjectsList(baseURL, appConfig.Registry); err != nil {
		return err
	}

	home.Address = appConfig.App.Address
	home.Port = appConfig.App.Port
	home.GlobalCompat = compatibility.CompatibilityLevel
	home.ListSubjects = subjects
	if home.HomeAddresse, err = GetHomeAddress(appConfig); err != nil {
		return err
	}

	if tplate, err = template.New("home").Parse(htmltpl.HomePage); err != nil {
		return nil
	}

	if err = tplate.Execute(w, home); err != nil {
		return nil
	}
	return nil
}
