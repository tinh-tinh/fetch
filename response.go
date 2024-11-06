package fetch

type Response struct {
	Data       interface{}
	Status     int
	StatusText string
	Error      error
}
