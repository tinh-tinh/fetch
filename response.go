package fetch

type Response[M any] struct {
	Data       M
	DataRaw    []byte
	Status     int
	StatusText string
	Error      error
}
