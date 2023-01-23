package app

import (
	c "github.com/chaewonkong/go-template/app/controller"
	"go.uber.org/fx"
)

type Controller struct {
	Book c.BookController
	User c.UserController
}

func NewController(
	book c.BookController,
	user c.UserController,
) Controller {
	return Controller{book, user}
}

var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		c.NewBookController,
		c.NewUserController,
		NewController,
	),
)
