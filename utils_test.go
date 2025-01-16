package fetch_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch/v2"
)

func Test_parseQuery(t *testing.T) {
	type QueryDto struct {
		Name     string `query:"name"`
		Age      int    `query:"age"`
		FromDate string `query:"fromDate"`
		ToDate   string `query:"toDate"`
		IsAdmin  bool   `query:"isAdmin"`
		IsDirect bool
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

	str2 := fetch.ParseQuery(map[string]interface{}{
		"name":     "Abc",
		"age":      13,
		"fromDate": "2024-01-01",
		"toDate":   "2024-12-12",
		"isAdmin":  true,
	})
	require.NotEmpty(t, str2)

	type QueryDto2 struct {
		Name     string `query:"name"`
		FromDate string `query:"fromDate"`
		ToDate   string `query:"toDate"`
	}
	queryArr := []*QueryDto2{
		{
			FromDate: "2024-01-01",
			ToDate:   "2024-12-12",
		},
		{
			Name: "Abc",
		},
	}
	str3 := fetch.ParseQuery(queryArr)
	require.Equal(t, "fromDate=2024-01-01&toDate=2024-12-12&name=Abc", str3)

	queryStruct := QueryDto2{
		FromDate: "2024-01-01",
		ToDate:   "2024-12-12",
	}

	str4 := fetch.ParseQuery(queryStruct)
	require.Empty(t, str4)

	str5 := fetch.ParseQuery(nil)
	require.Empty(t, str5)
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
