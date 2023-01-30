package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/service"
)

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

// @Tags         memo
// @Summary 특정 사용자가 특정 책에 대해 작성한 모든 메모를 가져오는 API
// @Description 특정 사용자가 특정 책에 대해 작성한 모든 메모를 가져오는 API. query string으로 userId와 bookId를 넘기면 됨.
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Param userId query string true "aaaa-bbbb-cccc"
// @Param bookId query string true "2"
// @Success 200 {object} []entity.Memo
// @Router /memos [get]
func (m memoController) FindAllMemoByUserAndBookID(c echo.Context) error {
	userID := c.QueryParam("userId")
	bookID := c.QueryParam("bookId")

	bookIDUint, err := strconv.ParseUint(bookID, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	memos, err := m.memoService.FindAllMemoByUserAndBookID(userID, uint(bookIDUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, memos)
}

// @Tags         memo
// @Summary 메모 추가 API
// @Description 특정 사용자가 특정 책에 대해 메모를 작성하는 API
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Param body body service.CreateMemoParam true "service.CreateMemoParam{}"
// @Success 201 {object} entity.Memo
// @Router /memos [post]
func (m memoController) CreateMemo(c echo.Context) error {
	param := service.CreateMemoParam{}
	if err := c.Bind(&param); err != nil {
		return c.String(http.StatusBadRequest, "Bad request, check request body")
	}

	memo, err := m.memoService.CreateMemo(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, memo)
}
