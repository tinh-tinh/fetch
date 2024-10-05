package fetch

import (
	"testing"

	"github.com/stretchr/testify/require"
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

	str := ParseQuery([]interface{}{query})
	require.Equal(t, "name=Abc&age=13&fromDate=2024-01-01&toDate=2024-12-12&isAdmin=true", str)
}
