package fetch_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch"
)

func Test_parseQuery(t *testing.T) {
	type QueryDto struct {
		Name     string `query:"name"`
		Age      int    `query:"age"`
		FromDate string `query:"fromDate"`
		ToDate   string `query:"toDate"`
		IsAdmin  bool   `query:"isAdmin"`
	}

	query := &QueryDto{
		Name:     "Abc",
		Age:      13,
		FromDate: "2024-01-01",
		ToDate:   "2024-12-12",
		IsAdmin:  true,
	}

	str := fetch.ParseQuery([]interface{}{query})
	require.Equal(t, "name=Abc&age=13&fromDate=2024-01-01&toDate=2024-12-12&isAdmin=true", str)
}

func Test_ParseData(t *testing.T) {
	type DataDto struct {
		Name string `json:"name,omitempty"`
		Age  int    `json:"age"`
	}

	data := &DataDto{
		Name: "Abc",
		Age:  13,
	}

	str := fetch.ParseData(data)
	reader := strings.NewReader("{\"name\":\"Abc\",\"age\":13}")
	require.Equal(t, reader, str)

	null := fetch.ParseData(nil)
	require.Nil(t, null)
}
