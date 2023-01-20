package app

import (
	"github.com/chaewonkong/go-template/app/controller"
	"go.uber.org/fx"
)

type Controller struct {
	Book controller.BookController
}

func NewController(book controller.BookController) Controller {
	return Controller{book}
}

var ControllerModule = fx.Module(
	"controller",
	fx.Provide(controller.NewBookController, NewController),
)
