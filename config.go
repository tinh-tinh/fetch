package fetch

import (
	"net/http"
	"time"
)

type Config struct {
	// Url is the server url that will be used for the request
	Url string
	// Method is the request method to be used when making the request
	BaseUrl         string
	Headers         http.Header
	Params          map[string]interface{}
	Data            map[string]interface{}
	Timeout         time.Duration
	WithCredentials bool
	ResponseType    string
}
