package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/server/middleware"
	"github.com/mafzaidi/elog/internal/service"
	"github.com/mafzaidi/elog/pkg/response"
	"github.com/mafzaidi/elog/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceHandler struct {
	serviceUC service.UseCase
}

func NewServiceHandler(uc service.UseCase) service.Handler {
	return &ServiceHandler{
		serviceUC: uc,
	}
}

func (h *ServiceHandler) FilterHasAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		pl := service.FilterHasAccountPayload{}
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

		data, err := h.serviceUC.ServicesHasAccount(authenticatedUserID, pl.IsActive)
		if err != nil {
			return response.ErrorHandler(c, http.StatusNotFound, "NotFound", err.Error())
		}

		var resp []service.FilterHasAccountResponse
		for _, d := range data {
			resp = append(resp, service.FilterHasAccountResponse{
				ID:          d.ID.Hex(),
				ServiceCode: d.Attributes.Code,
				ServiceName: d.Attributes.Name,
			})
		}

		return response.SuccesHandler(c, &response.Response{
			Message: "OK",
			Data:    resp,
		})
	}
}
