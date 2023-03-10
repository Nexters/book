package auth

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type (
	// BearerAuth Bearer 기반 Authentication Interface
	BearerAuth interface {
		GetToken(c echo.Context) (string, error)
		ParseToken(auth string) (string, error)
		ValidateBearerHeader(next echo.HandlerFunc) echo.HandlerFunc
	}

	// bearerAuth Bearer 기반 Authentication Struct
	bearerAuth struct {
		Authorization string `header:"Authorization" validate:"required,startsWith=Bearer"`
	}
)

// NewBearerAuth 생성자
func NewBearerAuth() BearerAuth {
	return bearerAuth{}
}

// ValidateBearerHeader middleware
func (b bearerAuth) ValidateBearerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		binder := new(echo.DefaultBinder)
		err := binder.BindHeaders(c, &b)

		_, err = b.ParseToken(b.Authorization)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Provide bearer user token")
		}
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

// GetToken header에서 토큰 추출
func (b bearerAuth) GetToken(c echo.Context) (token string, err error) {
	binder := new(echo.DefaultBinder)
	err = binder.BindHeaders(c, &b)
	if err != nil {
		return
	}
	token, err = b.ParseToken(b.Authorization)
	if err != nil {
		return
	}

	return
}

// ParseToken 토큰 파싱
func (b bearerAuth) ParseToken(auth string) (token string, err error) {
	// Validate Bearer token with uuid
	if matched, matchErr := regexp.MatchString("^Bearer\\s.+", auth); matchErr != nil || !matched {
		if matchErr == nil {
			err = errors.New("RegExp '^Bearer\\s.+' match failed")
			return
		}
		err = matchErr
		return
	}
	token = strings.Split(auth, " ")[1]

	return
}

// BearerAuthModule bearer auth 모듈
var BearerAuthModuole = fx.Module("auth/bearer", fx.Provide(NewBearerAuth))
