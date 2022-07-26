package v1

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) Init() *echo.Echo {
	e := echo.New()

	v1 := e.Group("/v1")
	{
		h.InitUserRoutes(v1)
	}

	return e
}
