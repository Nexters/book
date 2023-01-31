package app

import (
	"github.com/nexters/book/app/service"
	"go.uber.org/fx"
)

// ServiceModule 서비스 모듈
var ServiceModule = fx.Module(
	"service",
	fx.Provide(service.NewBookService, service.NewMemoService),
)
