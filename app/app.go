package app

import (
	"context"
	"log"
	"net/http"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
	_ "github.com/nexters/book/docs"
	"github.com/nexters/book/external/search"
	"go.uber.org/fx"
)

func bindRoute(e *echo.Echo, c Controller) {
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	b := e.Group("/books")
	u := e.Group("/users")
	m := e.Group("/memos")
	b.GET("", c.Book.FetchAll)
	b.GET("/:isbn", c.Book.FindBookByISBN)
	b.GET("/search", c.Book.Search)
	b.POST("", c.Book.CreateBook)
	u.POST("", c.User.CreateUser)
	m.GET("", c.Memo.FindAllMemoByUserAndBookID)
	m.POST("", c.Memo.CreateMemo)
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
				e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
					AllowOrigins: []string{"http://localhost:3030", "http://localhost:3000"},
					AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
				}))
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
	fx.Options(ServiceModule),
	fx.Invoke(RegisterHooks),
)
