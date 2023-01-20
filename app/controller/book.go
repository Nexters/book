package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	BookController interface {
		FetchAll(ctx echo.Context) error
	}
	bookController struct{}
)

func NewBookController() BookController {
	return bookController{}
}

func (b bookController) FetchAll(c echo.Context) error {
	return c.String(http.StatusOK, "book")
}
