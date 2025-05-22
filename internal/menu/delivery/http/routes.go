package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/menu"
)

func MapRoutes(g *echo.Group, h menu.Handler) {
	g.GET("", h.Filter())
}
