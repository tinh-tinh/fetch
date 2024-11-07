package fetch_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch"
)

func Test_Config(t *testing.T) {
	instance := fetch.Create(&fetch.Config{
		BaseUrl:      "https://jsonplaceholder.typicode.com",
		Headers:      http.Header{"x-api-key": []string{"abcd", "efgh"}},
		ResponseType: "json",
	})

	req, err := instance.GetConfig("GET", "", nil)
	require.Nil(t, err)
	require.Equal(t, 2, len(req.Header.Values("x-api-key")))
	require.Equal(t, "abcd", req.Header.Values("x-api-key")[0])
	require.Equal(t, "efgh", req.Header.Values("x-api-key")[1])
}
