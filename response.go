package fetch

import "encoding/json"

type Response struct {
	Data       []byte
	Status     int
	StatusText string
	Error      error
}

func (r *Response) Format(model interface{}) *Response {
	if model != nil {
		err := json.Unmarshal(r.Data, model)
		if err != nil {
			r.Error = err
		}
	}

	return r
}
