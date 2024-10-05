package fetch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Get(t *testing.T) {
	instance := Create(&Config{
		BaseUrl: "https://pokeapi.co",
	})

	res, err := instance.Get("/api/v2/pokemon/ditto")
	require.Nil(t, err)
	require.Equal(t, 200, res.Status)
}
