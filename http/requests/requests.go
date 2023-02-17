package requests

import (
	"bytes"
	"io"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type (
	HttpRequest[T interface{}] interface {
		GET(url string) (T, error)
		POST(url string, body interface{}) (T, error)
		// PATCH()
		// DELETE()
	}

	httpRequest[T interface{}] struct {
		headerParams map[string]string
	}
)

func NewHttp[T interface{}](params map[string]string) HttpRequest[T] {
	return httpRequest[T]{params}
}

func (h httpRequest[T]) POST(url string, body interface{}) (res T, err error) {
	byte, err := json.Marshal(body)
	if err != nil {
		return
	}

	buff := bytes.NewBuffer(byte)

	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		return
	}

	for key, val := range h.headerParams {
		req.Header.Add(key, val)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return
	}

	return
}

func (h httpRequest[T]) GET(url string) (res T, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	for key, val := range h.headerParams {
		req.Header.Add(key, val)
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return
	}

	return
}
