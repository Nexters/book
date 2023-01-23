package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/repository"
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

func (u userController) CreateUser(c echo.Context) error {
	user, err := u.repo.CreateUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}
