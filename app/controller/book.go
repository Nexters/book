package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/nexters/book/app/auth"
	"github.com/nexters/book/app/service"
	"github.com/nexters/book/app/service/payloads"
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
		FindBookAndAllMemosByBookID(c echo.Context) error
		Search(c echo.Context) error
		CreateBook(c echo.Context) error
		FindBookByISBN(c echo.Context) error
		UpdateBook(c echo.Context) error
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
// @Param isReading query bool true "default = true"
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Success 200 {object} payloads.FindAllBooksPayload
// @Router /books [get]
func (b bookController) FetchAll(c echo.Context) error {
	isReading := c.QueryParam("isReading")
	token, err := b.auth.GetToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	isReadingBool, err := strconv.ParseBool(isReading)
	if err != nil {
		isReadingBool = true
	}

	books, err := b.bookService.FindAllBooks(token, isReadingBool)

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
// @Param category query string false "comment"
// @Success 200 {object} entity.Book
// @Router /books/{bookId} [get]
func (b bookController) FindBookByISBN(c echo.Context) error {
	ISBN := c.Param("isbn")
	category := c.QueryParam("category")
	book, err := b.bookService.FindBookByISBN(ISBN, category)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, book)
}

// @Tags         book
// @Summary bookID 혹은 ISBN으로 책과 모든 메모 조회 API
// @Description bookID로 유저의 책과 그에 대한 모든 메모를 조회하는 API
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Param bookId path string true "12345678"
// @Param isbn query bool false "true"
// @Param category query string false "comment"
// @Success 200 {object} entity.Book
// @Router /books/{bookId} [get]
func (b bookController) FindBookAndAllMemosByBookID(c echo.Context) error {
	bookID := c.Param("bookId")

	// isbn인 경우 대응
	isbnString := c.QueryParam("isbn")
	category := c.QueryParam("category")

	if isbnString != "" {
		isISBN, err := strconv.ParseBool(isbnString)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if isISBN {
			book, err := b.bookService.FindBookByISBN(bookID, category)

			switch err {
			case gorm.ErrRecordNotFound:
				return c.String(http.StatusOK, "{}")
			case nil:
				return c.JSON(http.StatusOK, book)
			default:
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}
	}

	bookIDUint, err := strconv.ParseUint(bookID, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	book, err := b.bookService.FindBookAndAllMemosByBookID(uint(bookIDUint), category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

// @Tags         book
// @Summary 책을 읽는 중/완독 설정하는 API
// @Description 특정 책의 읽는 중/완독 상태를 업데이트하는 API
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Param request body payloads.UpdateBookPayload true "payloads.UpdateBookPayload{}"
// @Param bookId path string true "12345678"
// @Success 200 {object} entity.Book "entity.Book{}"
// @Router /books/{bookId} [patch]
func (b bookController) UpdateBook(c echo.Context) error {
	_, err := b.auth.GetToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	param := payloads.UpdateBookPayload{}
	if err := c.Bind(&param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// strconv.ParseUint(bookID, 10, 64)
	bookID, err := strconv.ParseUint(c.Param("bookId"), 10, 64)

	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res, err := b.bookService.UpdateBook(uint(bookID), *param.IsReading)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
