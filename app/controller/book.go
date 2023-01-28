package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nexters/book/app/service"
	"github.com/nexters/book/external/search"
)

type CreateBookParam struct {
	ISBN   string `json:"ISBN"`
	Title  string `json:"title"`
	UserID string `json:"userId"`
}

type (
	BookController interface {
		FetchAll(ctx echo.Context) error
		Search(c echo.Context) error
		CreateBook(c echo.Context) error
		FindBookByISBN(c echo.Context) error
	}
	bookController struct {
		bookSearch  search.BookSearch
		bookService service.BookService
	}
)

func NewBookController(s search.BookSearch, svc service.BookService) BookController {
	return bookController{s, svc}
}

// @Tags         book
// @Summary Find all books
// @Description Find all books by userId
// @Accept json
// @Produce json
// @Param userId query string true "abcd-efgh-1234"
// @Success 200 {object} []entity.Book
// @Router /books [get]
func (b bookController) FetchAll(c echo.Context) error {

	userId := c.QueryParam("userId")

	books, err := b.bookService.FindAllBooks(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, books)
}

// @Tags         book
// @Summary ISBN으로 책 조회 API
// @Description ISBN으로 책의 상세 내용을 조회하는 API
// @Accept json
// @Produce json
// @Param bookId path string true "12345678"
// @Success 200 {object} entity.Book
// @Router /books/{bookId} [get]
func (b bookController) FindBookByISBN(c echo.Context) error {
	ISBN := c.Param("isbn")
	book, err := b.bookService.FindBookByISBN(ISBN)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, book)
}

// @Tags         book
// @Summary 책 검색 API
// @Description Naver API를 이용해 책을 검색하게 하는 API query string으로 title을 넘기면 검색 결과를 반환.
// @Accept json
// @Produce json
// @Param title query string true "미움받을 용기"
// @Success 200 {object} []search.SearchItem
// @Router /books [get]
func (b bookController) Search(c echo.Context) error {
	title := c.QueryParam("title")
	res, err := b.bookSearch.SearchBookByTitle(title)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, res.Items)
}

// @Tags         book
// @Summary 읽을 책을 등록하는 API
// @Description 책의 ISBN, 제목, userId를 body로 제공하면 읽을책으로 등록하는 API
// @Accept json
// @Produce json
// @Param request body CreateBookParam true "CreateBookParam{}"
// @Success 201 {object} entity.Book
// @Router /books/search [post]
func (b bookController) CreateBook(c echo.Context) error {
	bookParam := CreateBookParam{}
	if err := c.Bind(&bookParam); err != nil {
		return c.String(http.StatusBadRequest, "Provide IBSN and book title correctly")
	}

	res, err := b.bookService.CreateBook(bookParam.Title, bookParam.ISBN, bookParam.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, res)
}
