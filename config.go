package fetch

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	// Url is the server url that will be used for the request
	Url string
	// BaseUrl is the base url that will be used for the request
	BaseUrl string
	// Headers are the headers that will be used for the request
	Headers http.Header
	// Params are the params that will be used for the request
	Params map[string]interface{}
	// Data is the data that will be used for the request
	Data map[string]interface{}
	// Timeout is the timeout that will be used for the request
	Timeout time.Duration
	// WithCredentials is the with credentials that will be used for the request
	WithCredentials bool
	// ResponseType is the response type that will be used for the request
	ResponseType string
}

// GetConfig returns a new *http.Request with the given method, uri and input.
// The request will have the base url, headers, params and with credentials set
// from the config. If the response type is json, the request will have the
// content type set to application/json. If the request has cookies, they will be
// added to the request.
func (f *Fetch) GetConfig(method string, uri string, input io.Reader) (*http.Request, error) {
	var formatUrl string
	if f.Config != nil {
		formatUrl = f.Config.BaseUrl
	}

	formatUrl += IfSlashPrefixString(uri)
	if len(f.Config.Params) > 0 {
		formatUrl += "?" + ParseQuery(f.Config.Params)
	}

	fullUrl, err := url.ParseRequestURI(formatUrl)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	if input != nil {
		req, err = http.NewRequest(method, fullUrl.String(), input)
	} else {
		req, err = http.NewRequest(method, fullUrl.String(), nil)
	}

	if f.Config.ResponseType == "json" {
		req.Header.Set("Content-Type", "application/json")
	}

	if f.Config.Headers != nil {
		for k, v := range f.Config.Headers {
			for _, vv := range v {
				if req.Header.Get(k) == "" {
					req.Header.Set(k, vv)
				} else {
					req.Header.Add(k, vv)
				}
			}
		}
	}

	if f.Config.WithCredentials && f.cookies != nil {
		for _, cookie := range f.cookies {
			req.AddCookie(cookie)
		}
	}

	return req, err
}
