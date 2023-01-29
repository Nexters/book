package auth

import (
	"errors"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	BearerAuth interface {
		GetToken(c echo.Context, i interface{}) (string, error)
		ParseToken(auth string) (string, error)
	}

	bearerAuth struct {
		Authorization string `header:"Authorization" validator:"required,startsWith=Bearer,"`
	}
)

func NewBearerAuth() BearerAuth {
	return bearerAuth{}
}

func (b bearerAuth) GetToken(c echo.Context, i interface{}) (token string, err error) {
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
