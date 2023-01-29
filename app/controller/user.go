package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/repository"
	_ "github.com/nexters/book/docs"
)

type (
	UserController interface {
		CreateUser(c echo.Context) error
	}
	userController struct {
		repo repository.UserRepository
	}
)

func NewUserController(r repository.UserRepository) UserController {
	return userController{r}
}

// @Tags         user
// @Summary 사용자 추가 API
// @Description API를 호출하면 UUID를 발급함. local storage에 저장해두고 userId로 사용하면 됨.
// @Accept json
// @Produce json
// @Success 200 {object} entity.User
// @Router /user [get]
func (u userController) CreateUser(c echo.Context) error {
	user, err := u.repo.CreateUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}
