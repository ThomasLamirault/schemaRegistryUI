package routing

import (
	"github.com/tlamirault/schemaRegistryUi/entities"
	"testing"
)

func TestGetRegistryBaseURL(t *testing.T) {
	var appConfig entities.AppConfig
	appConfig.Registry.Port = "443"
	url := GetRegistryBaseURL(appConfig)
	if url != "https://" {
		t.Errorf("Want %s, got %s", "https://", url)
	}
}

func TestGetRegistryBaseURL2(t *testing.T) {
	var appConfig entities.AppConfig
	appConfig.Registry.Port = "1234"
	url := GetRegistryBaseURL(appConfig)
	if url != "http://" {
		t.Errorf("Want %s, got %s", "https://", url)
	}
}
