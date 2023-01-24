package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/service"
)

type CreateMemoParam struct {
	UserID   string `json:"userId"`
	BookID   uint64 `json:"bookId"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

type (
	MemoController interface {
		FindAllMemoByUserAndBookID(c echo.Context) error
		CreateMemo(c echo.Context) error
	}
	memoController struct {
		memoService service.MemoService
	}
)

func NewMemoController(ms service.MemoService) MemoController {
	return memoController{ms}
}

func (m memoController) FindAllMemoByUserAndBookID(c echo.Context) error {
	userID := c.QueryParam("userId")
	bookID := c.QueryParam("bookId")

	bookIDUint64, err := strconv.ParseUint(bookID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	memos, err := m.memoService.FindAllMemoByUserAndBookID(userID, bookIDUint64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, memos)
}

func (m memoController) CreateMemo(c echo.Context) error {
	param := CreateMemoParam{}
	if err := c.Bind(&param); err != nil {
		return c.String(http.StatusBadRequest, "Bad request, check request body")
	}

	memo, err := m.memoService.CreateMemo(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, memo)
}
