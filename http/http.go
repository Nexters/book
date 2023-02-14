package http

import (
	"context"
	"log"
	"net/http"

	"github.com/nexters/book/app"
	"github.com/nexters/book/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nexters/book/http/auth"

	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/config"
	"github.com/nexters/book/config/environment"
	_ "github.com/nexters/book/docs"
	"github.com/nexters/book/external/search"
	"go.uber.org/fx"
)

// RegisterHooks 라이프사이클 훅 등록
func RegisterHooks(
	lifecycle fx.Lifecycle,
	e *echo.Echo,
	settings *config.Settings,
	db config.Database,
	validator *config.RequestValidator,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// hooks는 blocking으로 동작하므로 separate goroutine으로 실행 필요
			// https://github.com/uber-go/fx/issues/627#issuecomment-399235227

			go func() {
				e.Validator = validator
				e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
					AllowOrigins:     []string{"http://localhost:3030", "http://localhost:3000", "https://pieceofbook.com", "https://www.pieceofbook.com"},
					AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowCredentials, echo.HeaderAuthorization},
					AllowCredentials: true,
					AllowMethods:     []string{http.MethodGet, http.MethodOptions, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
				}))
				// e.Use(middleware.CORS())

				configureSwagger(settings)

				if err := db.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Memo{}); err != nil {
					log.Fatal(err)
				}

				if err := e.Start(settings.BindAddress()); err != nil {
					log.Fatal(err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}

// configureSwagger 스웨거 설정
func configureSwagger(settings *config.Settings) {
	docs.SwaggerInfo.Title = "Book API 문서"
	docs.SwaggerInfo.Description = "독서기록 작성 서비스 API 문서"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = settings.App.API_HOST
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	if settings.App.ENV == environment.DEV {
		docs.SwaggerInfo.Schemes = []string{"http", "http"}
	}
}

// Modules 메인 모듈
var Modules = fx.Module(
	"app",
	fx.Provide(config.NewSettings, echo.New, search.NewBookSearch),
	config.DBModule,
	ControllerModule,
	app.RepositoryModule,
	app.ServiceModule,
	auth.BearerAuthModuole,
	config.ValidatorModule,
	fx.Invoke(RegisterHooks),
)
