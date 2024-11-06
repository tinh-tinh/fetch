package fetch

import (
	"net/http"
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
