package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/service"
)

func MapRoutes(g *echo.Group, h service.Handler) {
	g.GET("/:user_id", h.FilterHasAccount())
}
