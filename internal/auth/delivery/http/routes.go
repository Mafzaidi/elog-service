package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/config"
	"github.com/mafzaidi/elog/internal/auth"
)

func MapPublicRoutes(g *echo.Group, h auth.Handler, cfg *config.Config) {
	g.POST("/register", h.Register())
	g.POST("/login", h.Login(cfg))
}

func MapPrivateRoutes(g *echo.Group, h auth.Handler) {
	g.POST("/logout", h.Logout())
	g.GET("/me", h.GetCurrentUser())
}
