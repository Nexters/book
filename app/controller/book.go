package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nexters/book/app/service"
	"github.com/nexters/book/external/search"
)

type CreateBookParam struct {
	ISBN  string `json:"ISBN"`
	Title string `json:"title"`
}

type (
	BookController interface {
		FetchAll(ctx echo.Context) error
		Search(c echo.Context) error
		CreateBook(c echo.Context) error
	}
	bookController struct {
		bookSearch  search.BookSearch
		bookService service.BookService
	}
)

func NewBookController(s search.BookSearch, svc service.BookService) BookController {
	return bookController{s, svc}
}

func (b bookController) FetchAll(c echo.Context) error {
	return c.String(http.StatusOK, "book")
}

func (b bookController) Search(c echo.Context) error {
	title := c.QueryParam("title")
	res, err := b.bookSearch.SearchBookByTitle(title)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, res.Items)
}

func (b bookController) CreateBook(c echo.Context) error {
	bookParam := CreateBookParam{}
	if err := c.Bind(&bookParam); err != nil {
		return c.String(http.StatusBadRequest, "Provide IBSN and book title correctly")
	}

	res, err := b.bookService.CreateBook(bookParam.Title, bookParam.ISBN)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, res)
}
