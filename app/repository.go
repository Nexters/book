package app

import (
	r "github.com/nexters/book/app/repository"
	"go.uber.org/fx"
)

var RepositoryModule = fx.Module(
	"repository module",
	fx.Provide(r.NewUserRepository, r.NewMemoRepository, r.NewBookRepository),
)
