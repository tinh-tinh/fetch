package fetch_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch/v2"
)

func TestBuildQueryParams(t *testing.T) {
	// Define a struct to test
	type Query struct {
		Name  string   `query:"name"`
		Age   int      `query:"age"`
		Tags  []string `query:"tags"`
		Empty string
	}

	query := Query{
		Name:  "John",
		Age:   30,
		Tags:  []string{"go", "programming"},
		Empty: "",
	}

	// Happy path test
	values := fetch.BuildQueryParams(query)
	require.Equal(t, "age=30&name=John&tags=go&tags=programming", values.Encode())

	// Test with nil input
	values = fetch.BuildQueryParams(nil)
	require.Equal(t, "", values.Encode())

	// Test with unsupported type
	values = fetch.BuildQueryParams(123)
	require.Equal(t, "", values.Encode())

	// Test with nested struct
	type NestedQuery struct {
		Title string `query:"title"`
		Query Query  `query:"nested"`
	}

	nested := NestedQuery{
		Title: "Developer",
		Query: query,
	}

	values = fetch.BuildQueryParams(nested)
	require.Equal(t, "age=30&name=John&tags=go&tags=programming&title=Developer", values.Encode())

}

func TestParseQuery(t *testing.T) {
	// Define structs to test
	type Query1 struct {
		Name string `query:"name"`
		Age  int    `query:"age"`
	}

	type Query2 struct {
		Active bool `query:"active"`
		Score  int  `query:"score"`
	}

	query1 := Query1{
		Name: "Alice",
		Age:  25,
	}

	query2 := Query2{
		Active: true,
		Score:  100,
	}

	// Happy path test
	result := fetch.ParseQuery(query1, query2)
	require.Equal(t, "age=25&name=Alice&active=true&score=100", result)

	// Test with nil input
	result = fetch.ParseQuery(nil)
	require.Equal(t, "", result)
}
