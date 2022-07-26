package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/kwarabei/Discoverish/api"
	v1 "github.com/kwarabei/Discoverish/internal/handlers/v1"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Mount(db *gorm.DB) *echo.Echo {
	h := v1.NewHandler(db)
	APIv1 := h.Init()

	APIv1.Validator = &CustomValidator{validator: validator.New()}

	APIv1.GET("/swagger/*", echoSwagger.WrapHandler)

	return APIv1
}
