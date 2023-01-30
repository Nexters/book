package controller

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nexters/book/app/repository"
	_ "github.com/nexters/book/docs"
)

type (
	UserController interface {
		CreateUserAndToken(c echo.Context) error
		FindUser(c echo.Context) error
	}
	userController struct {
		repo repository.UserRepository
	}
)

func NewUserController(r repository.UserRepository) UserController {
	return userController{r}
}

type AuthHeader struct {
	Authorization string `header:"Authorization"`
}

func getToken(auth string) (token string, err error) {
	//validation
	token = strings.Split(auth, " ")[1]

	return
}

// @Tags         user
// @Summary 사용자 추가 API
// @Description API를 호출하면 UUID를 token으로 발급함. local storage에 저장해두고 userId로 사용하면 됨.
// @Accept json
// @Produce json
// @Success 200 {object} entity.User
// @Router /users/token [get]
func (u userController) CreateUserAndToken(c echo.Context) error {
	user, err := u.repo.CreateUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

// @Tags         user
// @Summary 사용자 조회 API
// @Description Authorization header의 bearer token을 이용해 사용자를 조회함
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Success 200 {object} entity.User
// @Router /users [get]
func (u userController) FindUser(c echo.Context) error {
	// h := AuthHeader{}
	// binder := new(echo.DefaultBinder)
	// binder.BindHeaders(c, &h)

	// token, err := getToken(h.Authorization)
	// if err != nil {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, err)
	// }

	// fmt.Println(token)

	return c.JSON(http.StatusOK, "ok")
}
