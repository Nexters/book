package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/auth"
	c "github.com/nexters/book/app/controller"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
)

type Controller struct {
	Book c.BookController
	User c.UserController
	Memo c.MemoController
}

func NewController(
	book c.BookController,
	user c.UserController,
	memo c.MemoController,
) Controller {
	return Controller{book, user, memo}
}
func bindRoute(e *echo.Echo, c Controller, ba auth.BearerAuth) {
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	b := e.Group("/books")
	u := e.Group("/users")
	m := e.Group("/memos")
	b.GET("", c.Book.FetchAll, ba.ValidateBearerHeader)
	b.GET("/:isbn", c.Book.FindBookByISBN, ba.ValidateBearerHeader)
	b.GET("/search", c.Book.Search)
	b.POST("", c.Book.CreateBook, ba.ValidateBearerHeader)
	u.GET("/token", c.User.CreateUserAndToken)
	u.GET("", c.User.FindUser, ba.ValidateBearerHeader)
	m.GET("", c.Memo.FindAllMemoByUserAndBookID, ba.ValidateBearerHeader)
	m.POST("", c.Memo.CreateMemo, ba.ValidateBearerHeader)
}

var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		c.NewBookController,
		c.NewUserController,
		c.NewMemoController,
		NewController,
	),
	fx.Invoke(bindRoute),
)
