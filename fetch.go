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

// Create initializes and returns a new Fetch instance with the provided configuration.
// The configuration includes details like BaseUrl, Headers, Params, and more,
// which are used to customize HTTP requests made by the Fetch instance.
func Create(config *Config) *Fetch {
	return &Fetch{
		Config: config,
	}
}

// Get makes a GET request to the specified URL and returns a Response
// object containing the response status, status text, and body.
//
// If the params argument is provided, it is used to construct the
// query string of the URL.
//
// The Response object returned by this function can be used to access
// the status code and body of the HTTP response.
func (f *Fetch) Get(url string, params ...interface{}) *Response {
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

	return f.do("POST", url, ParseData(data))
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

	return f.do("PATCH", url, ParseData(data))
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

	return f.do("PUT", url, ParseData(data))
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
	return f.do("GET", url, nil)
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
