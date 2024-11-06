package fetch

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch/mock"
)

func Test_Get(t *testing.T) {
	testServer := httptest.NewServer(mock.App())
	defer testServer.Close()

	instance := Create(&Config{
		BaseUrl: testServer.URL,
	})

	type Response struct {
		Data []mock.User `json:"data"`
	}
	res, err := instance.Schema(&Response{}).Get("/api/users")
	require.Nil(t, err)
	require.Equal(t, 200, res.Status)
	fmt.Println(res.Data)
}
