package config

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator(validator *validator.Validate) *RequestValidator {
	return &RequestValidator{validator: validator}
}

func (rv *RequestValidator) Validate(i interface{}) error {
	if err := rv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

var ValidatorModule = fx.Module(
	"config/validator",
	fx.Provide(NewRequestValidator),
	fx.Provide(validator.New),
)
