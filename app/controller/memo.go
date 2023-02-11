package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/auth"
	"github.com/nexters/book/app/service"
)

type (
	// MemoController MemoController Interface
	MemoController interface {
		CreateMemo(c echo.Context) error
		UpdateMemo(c echo.Context) error
		DeleteMemo(c echo.Context) error
	}

	// memoController memoController Struct
	memoController struct {
		memoService service.MemoService
		auth        auth.BearerAuth
	}
)

type UpdateMemoPayload struct {
	Text     string `json:"text"`
	Category string `json:"category"`
}

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

// @Tags         memo
// @Summary 메모 수정 API
// @Description 특정 사용자가 특정 책에 대해 메모를 수정하는 API
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Param memoId path string true "12345678"
// @Param body body UpdateMemoPayload true "UpdateMemoPayload{}"
// @Success 201 {object} entity.Memo
// @Router /memos/{memoId} [patch]
func (m memoController) UpdateMemo(c echo.Context) error {
	_, err := m.auth.GetToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	memoId, err := strconv.ParseUint(c.Param("memoId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	param := UpdateMemoPayload{}
	if err := c.Bind(&param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request, check request body")
	}

	memo, err := m.memoService.UpdateMemo(service.UpdateMemoParam{MemoID: uint(memoId), Text: param.Text, Category: param.Category})
	if err != nil {
		if err == service.MaxLenError {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, memo)
}

// @Tags         memo
// @Summary 메모 삭제 API
// @Description 특정 사용자가 특정 책에 대해 작성한 특정 메모를 제거하는 API
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Param memoId path string true "12345678"
// @Success 202 string true "accepted"
// @Router /memos/{memoId} [delete]
func (m memoController) DeleteMemo(c echo.Context) error {
	_, err := m.auth.GetToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	memoId, err := strconv.ParseUint(c.Param("memoId"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = m.memoService.DeleteMemo(uint(memoId))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusAccepted, "delete success")
}
