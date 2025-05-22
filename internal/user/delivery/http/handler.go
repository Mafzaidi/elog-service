package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/internal/server/middleware"
	"github.com/mafzaidi/elog/internal/user"
	"github.com/mafzaidi/elog/pkg/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userUC user.UseCase
}

func NewUserHandler(uc user.UseCase) user.Handler {
	return &UserHandler{
		userUC: uc,
	}
}

func (h *UserHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {

		IDStr := c.Param("id")
		claims := middleware.GetUserFromContext(c)
		if claims.Group != "admin" && claims.UserID != IDStr {
			return response.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized", "you don't have access to this route")
		}

		ID, err := primitive.ObjectIDFromHex(IDStr)
		if err != nil {
			return response.ErrorHandler(c, http.StatusBadRequest, "BadRequest", err.Error())
		}

		data, err := h.userUC.User(ID)
		if err != nil {
			return response.ErrorHandler(c, http.StatusNotFound, "NotFound", err.Error())
		}

		resp := &user.GetResponse{
			ID:          data.ID,
			Username:    data.Username,
			Fullname:    data.Fullname,
			PhoneNumber: data.PhoneNumber,
			Email:       data.Email,
			Group:       data.Group,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		}

		return response.SuccesHandler(c, &response.Response{
			Message: "get user data successfully",
			Data:    resp,
		})
	}
}
