package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/account"
)

func MapRoutes(g *echo.Group, h account.Handler) {
	g.POST("/store", h.Create())
	g.GET("/:user_id", h.FilterUsersAccounts())
}
