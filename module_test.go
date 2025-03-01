package fetch_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch/v2"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func Test_AppModule(t *testing.T) {
	controller := func(module core.Module) core.Controller {
		ctrl := module.NewController("posts")
		httpFetch := fetch.Inject(module)

		ctrl.Get("/", func(ctx core.Ctx) error {
			type Post struct {
				UserID int    `json:"userId"`
				ID     int    `json:"id"`
				Title  string `json:"title"`
				Body   string `json:"body"`
			}

			type Posts []Post
			var data Posts
			res := httpFetch.Get("posts").Format(&data)

			return ctx.Status(res.Status).JSON(core.Map{
				"data": data,
			})
		})

		return ctrl
	}

	module := func() core.Module {
		appModule := core.NewModule(core.NewModuleOptions{
			Imports: []core.Modules{
				fetch.Register(&fetch.Config{
					BaseUrl: "https://jsonplaceholder.typicode.com",
				}),
			},
			Controllers: []core.Controllers{
				controller,
			},
		})

		return appModule
	}

	app := core.CreateFactory(module)
	app.SetGlobalPrefix("/api")

	testServer := httptest.NewServer(app.PrepareBeforeListen())
	defer testServer.Close()

	testClient := testServer.Client()

	res, err := testClient.Get(testServer.URL + "/api/posts")
	require.Nil(t, err)

	require.Equal(t, 200, res.StatusCode)
	data, err := io.ReadAll(res.Body)
	require.Nil(t, err)
	fmt.Println(string(data))
}
