package account

import "github.com/labstack/echo/v4"

type Handler interface {
	Create() echo.HandlerFunc
	FilterUsersAccounts() echo.HandlerFunc
}
