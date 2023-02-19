package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Gzip() echo.MiddlewareFunc {
	return middleware.Gzip()
}
