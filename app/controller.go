package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/book"
	"github.com/nexters/book/app/common/auth"
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
	b := e.Group("/books")
	u := e.Group("/users")
	m := e.Group("/memos")
	b.GET("", c.Book.FetchAll, ba.ValidateBearerHeader)
	b.GET("/:bookId", c.Book.FindBookAndAllMemosByBookID, ba.ValidateBearerHeader)
	// b.GET("/:isbn", c.Book.FindBookByISBN, ba.ValidateBearerHeader)
	b.GET("/search", c.Book.Search)
	b.POST("", c.Book.CreateBook, ba.ValidateBearerHeader)
	b.PATCH("/:bookId", c.Book.UpdateBook, ba.ValidateBearerHeader)
	b.DELETE("/:bookId", c.Book.DeleteBook, ba.ValidateBearerHeader)
	u.GET("/token", c.User.CreateUserAndToken)
	u.GET("", c.User.FindUser, ba.ValidateBearerHeader)
	m.POST("", c.Memo.CreateMemo, ba.ValidateBearerHeader)
	m.PATCH("/:memoId", c.Memo.UpdateMemo, ba.ValidateBearerHeader)
	m.DELETE("/:memoId", c.Memo.DeleteMemo, ba.ValidateBearerHeader)
}

// ControllerModule 컨트롤러
var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		book.NewBookController,
		user.NewUserController,
		memo.NewMemoController,
		NewController,
	),
	fx.Invoke(bindRoute),
)
