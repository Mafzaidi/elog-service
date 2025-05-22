package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/event"
)

func MapRoutes(g *echo.Group, h event.Handler) {
	g.GET("/:id", h.Get())
}
