package app

import (
	"github.com/nexters/book/app/service"
	"go.uber.org/fx"
)

var ServiceModule = fx.Module(
	"service",
	fx.Provide(service.NewBookService),
)
