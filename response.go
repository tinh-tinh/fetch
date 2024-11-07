package fetch

import "encoding/json"

type Response struct {
	Data       []byte
	Status     int
	StatusText string
	Error      error
}

// Format deserializes the response data into the given model. If the model is nil, the method does nothing.
// If the deserialization fails, the error is stored in the Response instance.
func (r *Response) Format(model interface{}) *Response {
	if model != nil {
		err := json.Unmarshal(r.Data, model)
		if err != nil {
			r.Error = err
		}
	}

	return r
}
