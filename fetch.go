package fetch

import (
	"io"
	"net/http"
)

type Fetch struct {
	Config   *Config
	Response Response
	cookies  []*http.Cookie
}

func Create(config *Config) *Fetch {
	return &Fetch{
		Config: config,
	}
}

func (f *Fetch) Get(url string, params ...interface{}) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}
	return f.do("GET", url, nil)
}

func (f *Fetch) Post(url string, data interface{}, params ...interface{}) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("POST", url, ParseData(data))
}

func (f *Fetch) Patch(url string, data interface{}, params ...interface{}) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("PATCH", url, ParseData(data))
}

func (f *Fetch) Put(url string, data interface{}, params ...interface{}) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("PUT", url, ParseData(data))
}

func (f *Fetch) Delete(url string, params ...interface{}) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}
	return f.do("GET", url, nil)
}

func (f *Fetch) do(method string, uri string, input io.Reader) *Response {
	req, err := f.GetConfig(method, uri, input)
	if err != nil {
		panic(err)
	}

	client := http.Client{}

	if f.Config.Timeout > 0 {
		client.Timeout = f.Config.Timeout
	}

	response := &Response{}
	resp, err := client.Do(req)
	if err != nil {
		response.Error = err
		return response
	}

	if resp.Cookies() != nil && f.Config.WithCredentials {
		f.cookies = resp.Cookies()
	}

	response.Status = resp.StatusCode
	response.StatusText = resp.Status

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.Error = err
	}

	response.Data = body
	return response
}
