package fetch

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/sony/gobreaker/v2"
)

type Fetch struct {
	Config   *Config
	Response Response
	cookies  []*http.Cookie
	cb       *gobreaker.CircuitBreaker[*http.Response]
}

// Create initializes and returns a new Fetch instance with the provided configuration.
// The configuration includes details like BaseUrl, Headers, Params, and more,
// which are used to customize HTTP requests made by the Fetch instance.
func Create(config *Config) *Fetch {
	if config.Encoder == nil {
		config.Encoder = json.Marshal
	}
	if config.Decoder == nil {
		config.Decoder = json.Unmarshal
	}

	instance := &Fetch{
		Config: config,
	}
	if config.CBSettings != nil {
		cb := gobreaker.NewCircuitBreaker[*http.Response](*config.CBSettings)
		instance.cb = cb
	}
	return instance
}

// Get makes a GET request to the specified URL and returns a Response
// object containing the response status, status text, and body.
//
// If the params argument is provided, it is used to construct the
// query string of the URL.
//
// The Response object returned by this function can be used to access
// the status code and body of the HTTP response.
func (f *Fetch) Get(url string, params ...any) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}
	return f.do("GET", url, nil)
}

// Post makes a POST request to the specified URL and returns a Response
// object containing the response status, status text, and body.
//
// If the data argument is provided, it is serialized to JSON and
// included in the request body.
//
// If the params argument is provided, it is used to construct the
// query string of the URL.
//
// The Response object returned by this function can be used to access
// the status code and body of the HTTP response.
func (f *Fetch) Post(url string, data interface{}, params ...interface{}) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("POST", url, ParseData(data, f.Config.Encoder))
}

// Patch makes a PATCH request to the specified URL and returns a Response
// object containing the response status, status text, and body.
//
// If the data argument is provided, it is serialized to JSON and
// included in the request body.
//
// If the params argument is provided, it is used to construct the
// query string of the URL.
//
// The Response object returned by this function can be used to access
// the status code and body of the HTTP response.
func (f *Fetch) Patch(url string, data interface{}, params ...interface{}) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("PATCH", url, ParseData(data, f.Config.Encoder))
}

// Put makes a PUT request to the specified URL and returns a Response
// object containing the response status, status text, and body.
//
// If the data argument is provided, it is serialized to JSON and
// included in the request body.
//
// If the params argument is provided, it is used to construct the
// query string of the URL.
//
// The Response object returned by this function can be used to access
// the status code and body of the HTTP response.
func (f *Fetch) Put(url string, data interface{}, params ...interface{}) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}

	return f.do("PUT", url, ParseData(data, f.Config.Encoder))
}

// Delete makes a DELETE request to the specified URL and returns a Response
// object containing the response status, status text, and body.
//
// If the params argument is provided, it is used to construct the
// query string of the URL.
//
// The Response object returned by this function can be used to access
// the status code and body of the HTTP response.
func (f *Fetch) Delete(url string, params ...interface{}) *Response {
	if len(params) > 0 {
		url += "?" + ParseQuery(params)
	}
	return f.do("DELETE", url, nil)
}

// do executes an HTTP request with the specified method, uri, and input,
// and returns a Response object containing the response status, status text,
// body, and any error that occurred during the request.
//
// The method parameter specifies the HTTP method to use (e.g., "GET", "POST").
// The uri parameter is the target URL for the request.
// The input parameter, if provided, is used as the request body.
//
// The function uses the configuration from the Fetch instance, such as timeout
// and credential settings. If the Config's WithCredentials is true, cookies from
// the response are stored for future requests.
//
// If an error occurs during the request or while reading the response body,
// the error is stored in the Response object.
func (f *Fetch) do(method string, uri string, input io.Reader) *Response {
	response := &Response{}

	req, err := f.GetConfig(method, uri, input)
	if err != nil {
		response.Error = err
		return response
	}

	var resp *http.Response
	if f.cb != nil {
		resp, err = f.cb.Execute(func() (*http.Response, error) {
			r, err := f.call(req)
			if err != nil {
				return nil, err
			}
			return r, nil
		})
	} else {
		resp, err = f.call(req)
	}

	if err != nil {
		response.Error = err
		return response
	}
	defer resp.Body.Close()

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
	response.decoder = f.Config.Decoder
	return response
}

func (f *Fetch) call(req *http.Request) (*http.Response, error) {
	client := http.Client{}

	if f.Config.Timeout > 0 {
		client.Timeout = f.Config.Timeout
	}

	return client.Do(req)
}
