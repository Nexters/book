package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"github.com/nexters/book/http/auth"

	_ "github.com/nexters/book/docs"
)

type (
	// UserController UserController Interface
	UserController interface {
		CreateUserAndToken(c echo.Context) error
		FindUser(c echo.Context) error
	}

	// userController userController Struct
	userController struct {
		repo    UserRepository
		auth    auth.BearerAuth
		service UserService
	}
)

// NewUserController 생성자
func NewUserController(r UserRepository, auth auth.BearerAuth, s UserService) UserController {
	return userController{r, auth, s}
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
// @Summary 사용자 통계 조회 API
// @Description Authorization header의 bearer token을 이용해 사용자 통계를 조회함
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 570d33ca-bd5c-4019-9192-5ee89229e8ec"
// @Success 200 {object} payloads.UserStatPayload
// @Router /users [get]
func (u userController) FindUser(c echo.Context) error {
	token, err := u.auth.GetToken(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	user, err := u.service.FindUserStat(token)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

// userRoute user route 등록 함수
func userRoute(e *echo.Echo, c UserController, auth auth.BearerAuth) {
	u := e.Group("/users")
	u.GET("/token", c.CreateUserAndToken)
	u.GET("", c.FindUser, auth.ValidateBearerHeader)
}

// UserControllerModule user controller를 등록하는 module
var UserControllerModule = fx.Module(
	"github.com/nexters/book/app/user/user_controller",
	fx.Provide(NewUserController),
	fx.Invoke(userRoute),
)
