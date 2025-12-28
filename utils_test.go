package fetch_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch/v2"
)

func Test_ParseData(t *testing.T) {
	type DataDto struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age"`
	}

	data := &DataDto{
		Name: "Abc",
		Age:  13,
	}

	str := fetch.ParseData(data, json.Marshal)
	reader := strings.NewReader("{\"name\":\"Abc\",\"age\":13}")
	require.Equal(t, reader, str)

	null := fetch.ParseData(nil, json.Marshal)
	require.Nil(t, null)
}
