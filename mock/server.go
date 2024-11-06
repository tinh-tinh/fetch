package mock

import (
	"net/http"
	"slices"

	"github.com/tinh-tinh/tinhtinh/core"
	"github.com/tinh-tinh/tinhtinh/dto/transform"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Controller(module *core.DynamicModule) *core.DynamicController {
	ctrl := module.NewController("users")

	users := []User{
		{ID: 1, Name: "Abc", Email: "abc@ab.com"},
		{ID: 2, Name: "Def", Email: "def@de.com"},
		{ID: 3, Name: "Ghi", Email: "ghi@gh.com"},
	}

	ctrl.Get("", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{
			"data": users,
		})
	})

	ctrl.Pipe(core.Body(&CreateUser{})).Post("", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{
			"data": ctx.Body(),
		})
	})

	ctrl.Get("/{id}", func(ctx core.Ctx) error {
		id := ctx.Param("id")
		userIdx := slices.IndexFunc(users, func(u User) bool {
			return u.ID == transform.ToInt(id)
		})
		if userIdx == -1 {
			return ctx.JSON(core.Map{
				"data": nil,
			})
		}
		return ctx.JSON(core.Map{
			"data": users[userIdx],
		})
	})

	ctrl.Pipe(core.Body(&CreateUser{})).Put("/{id}", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{
			"data": ctx.Body(),
		})
	})

	ctrl.Pipe(core.Body(&CreateUser{})).Patch("/{id}", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{
			"data": ctx.Body(),
		})
	})

	ctrl.Delete("/{id}", func(ctx core.Ctx) error {
		return ctx.JSON(core.Map{
			"data": ctx.Body(),
		})
	})

	return ctrl
}

func Module() *core.DynamicModule {
	module := core.NewModule(core.NewModuleOptions{
		Controllers: []core.Controller{Controller},
	})

	return module
}

func App() http.Handler {
	app := core.CreateFactory(Module)
	app.SetGlobalPrefix("/api")
	return app.PrepareBeforeListen()
}
