package fetch

import (
	"io"
	"strings"

	"github.com/tinh-tinh/tinhtinh/v2/core"
)

// ParseData serializes the given data into JSON format and returns an io.Reader
// containing the JSON data. If the input data is nil, the function returns nil.
// If the serialization fails, the function panics. The returned io.Reader can
// be used to read the JSON data as a stream.
func ParseData(data interface{}, encoder core.Encode) io.Reader {
	if data == nil {
		return nil
	}
	buffer, err := encoder(data)
	if err != nil {
		panic(err)
	}
	return strings.NewReader(string(buffer))
}
