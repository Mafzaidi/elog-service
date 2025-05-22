package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/event"
)

type EventHandler struct {
	eventUC event.UseCase
}

func NewEventHandler(uc event.UseCase) event.Handler {
	return &EventHandler{
		eventUC: uc,
	}
}

func (e *EventHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, struct {
				Message string `json:"message"`
				Status  int    `json:"status"`
			}{
				Message: "Bad Request",
				Status:  http.StatusBadRequest,
			})
		}

		event, err := e.eventUC.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, struct {
				Message string `json:"message"`
				Status  int    `json:"status"`
			}{
				Message: "Not Found",
				Status:  http.StatusNotFound,
			})
		}

		return c.JSON(http.StatusOK, event)
	}
}
