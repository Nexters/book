package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/book"

	"github.com/nexters/book/app/memo"
	"github.com/nexters/book/app/user"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
)

// appRoute  health check, swagger 라우팅
func appRoute(e *echo.Echo) {
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

// ControllerModule 컨트롤러
var ControllerModule = fx.Module(
	"controller",
	book.BookControllerModule,
	user.UserControllerModule,
	memo.MemoControllerModule,
	fx.Invoke(appRoute),
)
