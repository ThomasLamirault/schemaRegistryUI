package pageProcessors

import (
	"bytes"
	"github.com/tlamirault/schemaRegistryUi/entities"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

type ResponseMock struct {
	StatusCode int
	HeadersMap map[string]string
	Body       []byte
}

var emptyBody []byte

func (r *ResponseMock) MakeResponse() *http.Response {
	header := http.Header{}

	res := &http.Response{
		Status:           strconv.Itoa(r.StatusCode) + " " + http.StatusText(r.StatusCode),
		StatusCode:       r.StatusCode,
		Proto:            "HTTP/1.0",
		ProtoMajor:       1,
		ProtoMinor:       0,
		Header:           header,
		Body:             ioutil.NopCloser(bytes.NewReader(r.Body)),
		ContentLength:    int64(0),
		TransferEncoding: []string{},
		Close:            false,
		Uncompressed:     false,
		Trailer:          nil,
		Request:          nil,
		TLS:              nil,
	}
	return res
}

func TestGetResponseBodyByte(t *testing.T) {
	var res ResponseMock
	res.StatusCode = 200
	res.Body = []byte{}
	res.HeadersMap = map[string]string{}
	_, err := GetResponseBodyByte(res.MakeResponse())
	if err != nil {
		t.Errorf("GetResponseBodyByte return error with good response input")
	}
}

func TestGetResponseBodyByte2(t *testing.T) {
	var res ResponseMock
	res.StatusCode = 200
	res.Body = []byte{}
	res.HeadersMap = map[string]string{}
	res.HeadersMap[contentEncodingHeaderKey] = gzipEncodingHeaderValue
	_, err := GetResponseBodyByte(res.MakeResponse())
	if err != nil {
		t.Errorf("GetResponseBodyByte return error with good response input")
	}
}

func TestGetHomeAddress(t *testing.T) {
	var appConfig entities.AppConfig
	appConfig.App.Address = "localhost"
	appConfig.App.Port = 1234
	if result, err := GetHomeAddress(appConfig); err != nil {
		t.Errorf("GetHomeAddress return error with correct input address : %s, port : %d", appConfig.App.Address, appConfig.App.Port)
	} else {
		if result != "http://"+appConfig.App.Address+":"+strconv.Itoa(appConfig.App.Port) {
			t.Errorf("GetHomeAddress return bad out with input address : %s, port : %d, result %s", appConfig.App.Address, appConfig.App.Port, result)
		}
	}
}

func TestGetHomeAddress2(t *testing.T) {
	var appConfig entities.AppConfig
	appConfig.App.Address = ""
	_, err := GetHomeAddress(appConfig)
	if err == nil {
		t.Errorf("GetHomeAddress not returning expected errors with Address : %s, err : %w", appConfig.App.Address, err)
	}
}
