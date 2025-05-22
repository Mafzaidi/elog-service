package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/config"
)

type Handler interface {
	Register() echo.HandlerFunc
	Login(cfg *config.Config) echo.HandlerFunc
}
