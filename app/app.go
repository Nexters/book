package app

import (
	"context"
	"log"
	"net/http"

	"github.com/chaewonkong/go-template/app/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func bindRoute(e *echo.Echo, c Controller) {
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	b := e.Group("/books")
	b.GET("", c.Book.FetchAll)
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
				// db.MakeMigration(&entity.Book{})
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
	fx.Provide(config.NewSettings, echo.New),
	fx.Options(config.DBModule),
	fx.Options(ControllerModule),
	fx.Invoke(RegisterHooks),
)
