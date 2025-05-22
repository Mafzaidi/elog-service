package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/menu"
	"github.com/mafzaidi/elog/pkg/response"
)

type MenuHandler struct {
	menuUC menu.UseCase
}

func NewMenuHandler(uc menu.UseCase) menu.Handler {
	return &MenuHandler{
		menuUC: uc,
	}
}

func (h *MenuHandler) Filter() echo.HandlerFunc {
	return func(c echo.Context) error {
		pl := menu.FilterPayload{}
		if err := c.Bind(&pl); err != nil {
			return response.ErrorHandler(c, http.StatusBadRequest, "BadRequest", err.Error())
		}

		data, err := h.menuUC.ActiveUserMenus(pl.IsActive, pl.Group)
		if err != nil {
			return response.ErrorHandler(c, http.StatusNotFound, "NotFound", err.Error())
		}

		var resp []menu.FilterResponse
		for _, d := range data {
			resp = append(resp, menu.FilterResponse{
				Title: d.Title,
				Url:   d.Url,
				Icon:  d.Icon,
				Group: d.Group,
			})
		}

		return response.SuccesHandler(c, &response.Response{
			Message: "OK",
			Data:    resp,
		})
	}
}
