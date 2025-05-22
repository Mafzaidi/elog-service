package menu

import "github.com/labstack/echo/v4"

type Handler interface {
	Filter() echo.HandlerFunc
}
