package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mafzaidi/elog/config"
	"github.com/mafzaidi/elog/internal/auth"
	"github.com/mafzaidi/elog/internal/server/middleware"
	"github.com/mafzaidi/elog/pkg/response"
	"github.com/mafzaidi/elog/pkg/utils"
)

type AuthHandler struct {
	authUC auth.UseCase
}

func NewAuthHandler(uc auth.UseCase) auth.Handler {
	return &AuthHandler{
		authUC: uc,
	}
}

func (h *AuthHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		pl := &auth.RegisterPayload{}

		if err := json.NewDecoder(c.Request().Body).Decode(&pl); err != nil {
			return response.ErrorHandler(c, http.StatusBadRequest, "BadRequest", err.Error())
		}

		if err := h.authUC.Register(pl); err != nil {
			return response.ErrorHandler(c, http.StatusInternalServerError, "InternalServerError", err.Error())
		}

		return response.SuccesHandler(c, &response.Response{
			Message: "user registered successfully",
		})
	}
}

func (h *AuthHandler) Login(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		pl := &auth.LoginPayload{}
		if err := json.NewDecoder(c.Request().Body).Decode(&pl); err != nil {
			return response.ErrorHandler(c, http.StatusBadRequest, "BadRequest", err.Error())
		}

		var validToken string
		if cookie, err := c.Cookie("jwt_token"); err == nil {
			validToken = cookie.Value
		}

		data, err := h.authUC.Login(pl.Email, pl.Password, validToken, cfg)
		if err != nil {
			return response.ErrorHandler(c, http.StatusInternalServerError, "InternalServerError", err.Error())
		}

		newCookie := new(http.Cookie)
		newCookie.Name = "jwt_token"
		newCookie.Value = data.Token
		// newCookie.HttpOnly = true
		newCookie.Secure = true
		newCookie.SameSite = http.SameSiteNoneMode
		newCookie.Expires = data.Claims.ExpiresAt.Time
		newCookie.Path = "/"

		c.SetCookie(newCookie)

		resp := &auth.LoginResponse{
			ID:          data.User.ID,
			Username:    data.User.Username,
			Fullname:    data.User.Fullname,
			PhoneNumber: data.User.Fullname,
			Email:       data.User.Email,
			Group:       data.User.Group,
			CreatedAt:   data.User.CreatedAt,
			UpdatedAt:   data.User.UpdatedAt,
			AccessToken: struct {
				Type      string    `json:"type"`
				Token     string    `json:"token"`
				ExpiresAt time.Time `json:"expires_at"`
			}{
				Type:      "Bearer",
				Token:     data.Token,
				ExpiresAt: data.Claims.ExpiresAt.Time,
			},
		}

		return response.SuccesHandler(c, &response.Response{
			Message: "user login successfully",
			Data:    resp,
		})
	}
}

func (h *AuthHandler) GetCurrentUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := middleware.GetUserFromContext(c)
		authenticatedUserID, err := utils.ToObjectID(claims.UserID)
		if err != nil {
			msg := fmt.Errorf("invalid token user ID: %v", err)
			return response.ErrorHandler(c, http.StatusBadRequest, "Unauthorized", msg.Error())
		}

		data, err := h.authUC.User(authenticatedUserID)
		if err != nil {
			return response.ErrorHandler(c, http.StatusNotFound, "DataNotFound", err.Error())
		}

		resp := &auth.GetCurrentUserResponse{
			ID:          data.ID,
			Username:    data.Username,
			Fullname:    data.Fullname,
			PhoneNumber: data.PhoneNumber,
			Email:       data.Email,
			Group:       data.Group,
		}

		return response.SuccesHandler(c, &response.Response{
			Message: "fetch user data successfully",
			Data:    resp,
		})
	}
}

func (h *AuthHandler) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		expiredCookie := &http.Cookie{
			Name:     "jwt_token",
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			MaxAge:   -1,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteNoneMode,
		}
		c.SetCookie(expiredCookie)

		return response.SuccesHandler(c, &response.Response{
			Message: "user logged out successfully",
		})
	}
}
