package fetch

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/tinh-tinh/tinhtinh/core"
)

type Fetch struct {
	schema   interface{}
	Config   *Config
	Response Response
}

func Create(config *Config) *Fetch {
	return &Fetch{
		Config: config,
	}
}

func (f *Fetch) Schema(data interface{}) *Fetch {
	f.schema = data
	return f
}

func (f *Fetch) Get(url string, params ...interface{}) (*Response, error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}
	return f.do("GET", url, nil)
}

func (f *Fetch) Post(url string, data interface{}, params ...interface{}) (*Response, error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("POST", url, ParseData(data))
}

func (f *Fetch) Patch(url string, data interface{}, params ...interface{}) (*Response, error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("PATCH", url, ParseData(data))
}

func (f *Fetch) Put(url string, data interface{}, params ...interface{}) (*Response, error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("PUT", url, ParseData(data))
}

func (f *Fetch) Delete(url string, params ...interface{}) (*Response, error) {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}
	return f.do("GET", url, nil)
}

func (f *Fetch) do(method string, uri string, input io.Reader) (*Response, error) {
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

	response := &Response{}
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

	if f.schema != nil {
		data := f.schema
		err = json.Unmarshal(body, data)
		if err != nil {
			response.Error = err
			return response, nil
		}
		f.schema = nil
		response.Data = data
		response.Error = nil
		return response, nil
	}

	response.Data = body
	response.Error = nil
	return response, nil
}
