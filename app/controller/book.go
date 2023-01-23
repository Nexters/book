package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nexters/book/external/search"
)

type (
	BookController interface {
		FetchAll(ctx echo.Context) error
		Search(c echo.Context) error
	}
	bookController struct {
		s search.BookSearch
	}
)

func NewBookController(s search.BookSearch) BookController {
	return bookController{s}
}

func (b bookController) FetchAll(c echo.Context) error {
	return c.String(http.StatusOK, "book")
}

func (b bookController) Search(c echo.Context) error {
	res, err := b.s.SearchBookByTitle("색채가")
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, res.Items)
}
