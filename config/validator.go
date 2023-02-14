package config

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// RequestValidator request validator Struct
type RequestValidator struct {
	validator *validator.Validate
}

// NewPrequestValidator 생성자
func NewRequestValidator(validator *validator.Validate) *RequestValidator {
	return &RequestValidator{validator: validator}
}

// Validate validate 메서드
func (rv *RequestValidator) Validate(i interface{}) error {
	if err := rv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

// ValidatorModule validator 모듈
var ValidatorModule = fx.Module(
	"config/validator",
	fx.Provide(NewRequestValidator),
	fx.Provide(validator.New),
)
