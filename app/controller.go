package app

import (
	c "github.com/nexters/book/app/controller"
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

var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		c.NewBookController,
		c.NewUserController,
		c.NewMemoController,
		NewController,
	),
)
