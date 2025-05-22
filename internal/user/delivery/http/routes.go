package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/user"
)

func MapRoutes(g *echo.Group, h user.Handler) {
	g.GET("/:id", h.Get())
}
