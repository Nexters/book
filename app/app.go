package app

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/external/search"
	"go.uber.org/fx"
)

func bindRoute(e *echo.Echo, c Controller) {
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	b := e.Group("/books")
	u := e.Group("/users")
	b.GET("", c.Book.FetchAll)
	b.GET("/search", c.Book.Search)
	u.POST("", c.User.CreateUser)
}

func RegisterHooks(
	lifecycle fx.Lifecycle,
	e *echo.Echo,
	settings *config.Settings,
	db config.Database,
	c Controller,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// hooks는 blocking으로 동작하므로 separate goroutine으로 실행 필요
			// https://github.com/uber-go/fx/issues/627#issuecomment-399235227
			go func() {
				bindRoute(e, c)

				if err := db.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.UserBooks{}, &entity.Memo{}); err != nil {
					log.Fatal(err)
				}

				if err := e.Start(settings.BindAddress()); err != nil {
					log.Fatal(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}

var Modules = fx.Module(
	"app",
	fx.Provide(config.NewSettings, echo.New, search.NewBookSearch),
	fx.Options(config.DBModule),
	fx.Options(ControllerModule),
	fx.Options(RepositoryModule),
	fx.Invoke(RegisterHooks),
)
