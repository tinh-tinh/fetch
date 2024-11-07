package fetch_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/fetch"
	"github.com/tinh-tinh/tinhtinh/core"
)

func Test_AppModule(t *testing.T) {
	controller := func(module *core.DynamicModule) *core.DynamicController {
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

	module := func() *core.DynamicModule {
		appModule := core.NewModule(core.NewModuleOptions{
			Imports: []core.Module{
				fetch.Register(&fetch.Config{
					BaseUrl: "https://jsonplaceholder.typicode.com",
				}),
			},
			Controllers: []core.Controller{
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
}
