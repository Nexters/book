package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/book"

	"github.com/nexters/book/http/auth"

	"github.com/nexters/book/app/memo"
	"github.com/nexters/book/app/user"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
)

// Controller 컨트롤러 구조체
type Controller struct {
	Book book.BookController
	User user.UserController
	Memo memo.MemoController
}

// NewController 생성자
func NewController(
	book book.BookController,
	user user.UserController,
	memo memo.MemoController,
) Controller {
	return Controller{book, user, memo}
}

// bindRoute 라우팅
func bindRoute(e *echo.Echo, c Controller, ba auth.BearerAuth) {
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	m := e.Group("/memos")
	m.POST("", c.Memo.CreateMemo, ba.ValidateBearerHeader)
	m.PATCH("/:memoId", c.Memo.UpdateMemo, ba.ValidateBearerHeader)
	m.DELETE("/:memoId", c.Memo.DeleteMemo, ba.ValidateBearerHeader)
}

// ControllerModule 컨트롤러
var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		memo.NewMemoController,
		NewController,
	),
	book.BookControllerModule,
	user.UserControllerModule,
	fx.Invoke(bindRoute),
)
