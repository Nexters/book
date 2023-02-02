package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/auth"
	"github.com/nexters/book/app/service"
)

type (
	// MemoController MemoController Interface
	MemoController interface {
		CreateMemo(c echo.Context) error
	}

	// memoController memoController Struct
	memoController struct {
		memoService service.MemoService
		auth        auth.BearerAuth
	}
)

// NewMemoController 생성자
func NewMemoController(ms service.MemoService, auth auth.BearerAuth) MemoController {
	return memoController{ms, auth}
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
	token, err := m.auth.GetToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	param := service.CreateMemoParam{}
	if err := c.Bind(&param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request, check request body")
	}

	if err := c.Validate(param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	memo, err := m.memoService.CreateMemo(param, token)
	if err != nil {
		if err == service.MaxLenError {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, memo)
}
