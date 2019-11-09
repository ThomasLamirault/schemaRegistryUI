package pageProcessors

import (
	"compress/gzip"
	"encoding/json"
	"github.com/tlamirault/schemaRegistryUi/entities"
	ce "github.com/tlamirault/schemaRegistryUi/errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	contentEncodingHeaderKey = "Content-Encoding"
	gzipEncodingHeaderValue  = "gzip"
)

func GetSchemaRegistryConfig(baseURL string, registry entities.Registry) (entities.Config, error) {
	var compatibility entities.Config
	resp, err := http.Get(baseURL + registry.Address + ":" + registry.Port + "/config")
	if err != nil {
		return compatibility, err
	}
	defer resp.Body.Close()
	body, err := GetResponseBodyByte(resp)
	if err != nil {
		return compatibility, err
	}
	err = json.Unmarshal(body, &compatibility)
	return compatibility, nil
}

func GetSubjectsList(baseURL string, registry entities.Registry) ([]string, error) {
	var listSubjects []string
	resp, err := http.Get(baseURL + registry.Address + ":" + registry.Port + "/subjects")
	if err != nil {
		return listSubjects, err
	}
	defer resp.Body.Close()
	body, err := GetResponseBodyByte(resp)
	if err != nil {
		return listSubjects, err
	}
	err = json.Unmarshal(body, &listSubjects)
	return listSubjects, nil
}

func GetResponseBodyByte(response *http.Response) ([]byte, error) {
	var reader io.Reader = response.Body
	var err error
	encoding := response.Header.Get(contentEncodingHeaderKey)
	if encoding == gzipEncodingHeaderValue {
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
	} else {
		reader = response.Body
	}
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetHomeAddress(appConfig entities.AppConfig) (string, error) {
	address := appConfig.App.Address
	port := strconv.Itoa(appConfig.App.Port)
	if address == "" || port == "" {
		return "", &ce.AppConfigError{Address: address, Port: port}
	} else {
		return "http://" + address + ":" + port, nil
	}
}
