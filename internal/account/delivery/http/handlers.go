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

func (h *AccountHandler) FilterUsersAccounts() echo.HandlerFunc {
	return func(c echo.Context) error {
		pl := account.FilterUsersAccountsPayload{}
		requestedUserID, err := utils.ToObjectID(c.Param("user_id"))

		if err != nil || requestedUserID == primitive.NilObjectID {
			msg := fmt.Errorf("invalid user_id param: %v", err)
			return response.ErrorHandler(c, http.StatusBadRequest, "BadRequest", msg.Error())
		}

		if err := c.Bind(&pl); err != nil {
			return response.ErrorHandler(c, http.StatusBadRequest, "BadRequest", err.Error())
		}

		claims := middleware.GetUserFromContext(c)
		if claims == nil {
			return response.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized", "token is invalid")
		}

		authenticatedUserID, err := utils.ToObjectID(claims.UserID)
		if err != nil || authenticatedUserID == primitive.NilObjectID {
			msg := fmt.Errorf("invalid token user ID: %v", err)
			return response.ErrorHandler(c, http.StatusBadRequest, "Unauthorized", msg.Error())
		}

		if requestedUserID != authenticatedUserID {
			return response.ErrorHandler(c, http.StatusForbidden, "Forbidden", "you are not allowed to access this resource")
		}

		data, err := h.accountUC.UserAccounts(authenticatedUserID, pl.IsActive)
		if err != nil {
			return response.ErrorHandler(c, http.StatusNotFound, "NotFound", err.Error())
		}

		var resp []account.FilterUsersAccountsResponse
		for _, d := range data {
			resp = append(resp, account.FilterUsersAccountsResponse{
				ID:     d.ID.Hex(),
				UserID: d.UserID.Hex(),
				Service: struct {
					ID   string `json:"id"`
					Code string `json:"code"`
					Key  string `json:"key"`
					Name string `json:"name"`
				}{
					ID:   d.Service.ID.Hex(),
					Code: d.Service.Code,
					Key:  d.Service.Key,
					Name: d.Service.Name,
				},
				Username:          d.Username,
				PasswordEncrypted: d.PasswordEncrypted,
				Salt:              d.Salt,
				Host:              d.Host,
				Notes:             d.Notes,
				IsActive:          d.IsActive,
				CreatedAt:         d.CreatedAt,
				UpdatedAt:         d.UpdatedAt,
			})
		}

		return response.SuccesHandler(c, &response.Response{
			Message: "OK",
			Data:    resp,
		})

	}
}
