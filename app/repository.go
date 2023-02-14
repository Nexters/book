package app

import (
	"github.com/nexters/book/app/book"
	"github.com/nexters/book/app/memo"
	"github.com/nexters/book/app/user"
	"go.uber.org/fx"
)

// RepositoryModule repository 모듈
var RepositoryModule = fx.Module(
	"repository module",
	fx.Provide(user.NewUserRepository, memo.NewMemoRepository, book.NewBookRepository),
)
