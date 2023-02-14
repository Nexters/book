package app

import (
	"github.com/nexters/book/app/book"
	"github.com/nexters/book/app/memo"
	"github.com/nexters/book/app/user"
	"go.uber.org/fx"
)

// ServiceModule 서비스 모듈
var ServiceModule = fx.Module(
	"service",
	fx.Provide(book.NewBookService, memo.NewMemoService, user.NewUserService),
)
