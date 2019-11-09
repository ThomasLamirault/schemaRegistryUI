package main

import (
	"github.com/spf13/viper"
	"github.com/tlamirault/schemaRegistryUi/entities"
	routing "github.com/tlamirault/schemaRegistryUi/routing"
	"html/template"
	"log"
)

var tplate *template.Template
var appConfig entities.AppConfig
var baseURL string
var homeAddress string

func main() {

	var err error

	appConfig, err = initConfiguration()
	if err != nil {
		log.Fatal("unable to init configuration", err)
	}

	err = routing.Routing(routing.GetRegistryBaseURL(appConfig), appConfig)
	if err != nil {
		log.Fatal("unable to route application", err)
	}
}

func initConfiguration() (entities.AppConfig, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetDefault("Address", "192.168.99.100")
	viper.SetDefault("Port", "8081")

	err := viper.ReadInConfig()
	if err != nil {
		return appConfig, err
	}
	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return appConfig, err
	}
	return appConfig, nil
}
