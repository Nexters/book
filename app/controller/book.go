package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nexters/book/app/auth"
	"github.com/nexters/book/app/service"
	"github.com/nexters/book/external/search"
)

// CreateBookParam 책 생성 parameters
type CreateBookParam struct {
	ISBN  string `json:"ISBN" validate:"required,isbn"`
	Title string `json:"title" validate:"required"`
}

type (
	// BookController BookController Interface
	BookController interface {
		FetchAll(ctx echo.Context) error
		Search(c echo.Context) error
		CreateBook(c echo.Context) error
		FindBookByISBN(c echo.Context) error
	}

	// bookController bookController Struct
	bookController struct {
		bookSearch  search.BookSearch
		bookService service.BookService
		auth        auth.BearerAuth
	}
)

// NewBookController 생성자
func NewBookController(s search.BookSearch, svc service.BookService, auth auth.BearerAuth) BookController {
	return bookController{s, svc, auth}
}

// @Tags         book
// @Summary 사용자가 등록한 모든 책을 조회하는 API
// @Description 사용자가 등록한 모든 책을 조회하는 API. TODO: 읽을책/완독 구분해 가져오게 할 예정
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Success 200 {object} []entity.Book
// @Router /books [get]
func (b bookController) FetchAll(c echo.Context) error {
	token, err := b.auth.GetToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	books, err := b.bookService.FindAllBooks(token)

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
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
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
// @Router /books/search [get]
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
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Param request body CreateBookParam true "CreateBookParam{}"
// @Success 201 {object} entity.Book
// @Router /books [post]
func (b bookController) CreateBook(c echo.Context) error {
	token, err := b.auth.GetToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	bookParam := CreateBookParam{}
	if err := c.Bind(&bookParam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Provide IBSN and book title correctly")
	}

	if err := c.Validate(bookParam); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := b.bookService.CreateBook(bookParam.Title, bookParam.ISBN, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, res)
}
