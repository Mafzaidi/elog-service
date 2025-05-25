package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/account"
	"github.com/mafzaidi/elog/internal/server/middleware"
	"github.com/mafzaidi/elog/pkg/response"
	"github.com/mafzaidi/elog/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountHandler struct {
	accountUC account.UseCase
}

func NewAccountHandler(uc account.UseCase) account.Handler {
	return &AccountHandler{
		accountUC: uc,
	}
}

func (h *AccountHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		pl := account.CreatePayload{}
		if err := json.NewDecoder(c.Request().Body).Decode(&pl); err != nil {
			return response.ErrorHandler(c, http.StatusBadRequest, "BadRequest", err.Error())
		}

		claims := middleware.GetUserFromContext(c)
		if claims == nil {
			return response.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized", "token is invalid")
		}

		userIDObj, err := utils.ToObjectID(claims.UserID)
		if err != nil || userIDObj == primitive.NilObjectID {
			msg := fmt.Errorf("objectID format is not valid: %v", userIDObj)
			return response.ErrorHandler(c, http.StatusBadRequest, "BadRequest", msg.Error())
		}

		params := &account.CreateParams{
			UserID:      userIDObj,
			PasswordApp: pl.PasswordApp,
			Username:    pl.Username,
			Password:    pl.Password,
			Host:        pl.Host,
			Notes:       pl.Notes,
			Service:     pl.Service,
			IsActive:    pl.IsActive,
		}

		if err := h.accountUC.Store(params); err != nil {
			return response.ErrorHandler(c, http.StatusInternalServerError, "InternalServerError", err.Error())
		}

		return response.SuccesHandler(c, &response.Response{
			Message: "account created successfully",
		})
	}
}
