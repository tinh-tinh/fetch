package fetch

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/tinh-tinh/tinhtinh/core"
)

type Fetch[M any] struct {
	Config   *Config
	Response Response[M]
}

func Create(config *Config) *Fetch[any] {
	return &Fetch[any]{
		Config: config,
	}
}

func CreateSchema[M any](config *Config) *Fetch[M] {
	return &Fetch[M]{
		Config: config,
	}
}

func (f *Fetch[M]) Get(url string, params ...interface{}) (*Response[M], error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}
	return f.do("GET", url, nil)
}

func (f *Fetch[M]) Post(url string, data interface{}, params ...interface{}) (*Response[M], error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("POST", url, ParseData(data))
}

func (f *Fetch[M]) Patch(url string, data interface{}, params ...interface{}) (*Response[M], error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("PATCH", url, ParseData(data))
}

func (f *Fetch[M]) Put(url string, data interface{}, params ...interface{}) (*Response[M], error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("PUT", url, ParseData(data))
}

func (f *Fetch[M]) Delete(url string, params ...interface{}) (*Response[M], error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}
	return f.do("GET", url, nil)
}

func (f *Fetch[M]) do(method string, uri string, input io.Reader) (*Response[M], error) {
	var fullUrl string
	if f.Config != nil {
		fullUrl = f.Config.BaseUrl
	}

	fullUrl += core.IfSlashPrefixString(uri)

	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	if input != nil {
		req, err = http.NewRequest(method, fullUrl, input)
	} else {
		req, err = http.NewRequest(method, fullUrl, nil)
	}

	if err != nil {
		return nil, err
	}

	if f.Config.ResponseType == "json" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := http.Client{}

	response := &Response[M]{}
	resp, err := client.Do(req)
	if err != nil {
		response.Error = err
		return response, nil
	}

	response.Status = resp.StatusCode
	response.StatusText = resp.Status

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.Error = err
		return response, nil
	}

	var data M

	err = json.Unmarshal(body, &data)
	if err != nil {
		response.Error = err
		return response, nil
	}

	response.Data = data
	response.Error = nil
	return response, nil
}
