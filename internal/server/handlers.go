package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	emdw "github.com/labstack/echo/v4/middleware"
	accounthttp "github.com/mafzaidi/elog/internal/account/delivery/http"
	accountrepo "github.com/mafzaidi/elog/internal/account/repository"
	accountuc "github.com/mafzaidi/elog/internal/account/usecase"
	authhttp "github.com/mafzaidi/elog/internal/auth/delivery/http"
	authrepo "github.com/mafzaidi/elog/internal/auth/repository"
	authuc "github.com/mafzaidi/elog/internal/auth/usecase"
	eventhttp "github.com/mafzaidi/elog/internal/event/delivery/http"
	eventrepo "github.com/mafzaidi/elog/internal/event/repository"
	eventuc "github.com/mafzaidi/elog/internal/event/usecase"
	menuhttp "github.com/mafzaidi/elog/internal/menu/delivery/http"
	menurepo "github.com/mafzaidi/elog/internal/menu/repository"
	menuuc "github.com/mafzaidi/elog/internal/menu/usecase"
	"github.com/mafzaidi/elog/internal/server/middleware"
	servicehttp "github.com/mafzaidi/elog/internal/service/delivery/http"
	servicerepo "github.com/mafzaidi/elog/internal/service/repository"
	serviceuc "github.com/mafzaidi/elog/internal/service/usecase"
	userhttp "github.com/mafzaidi/elog/internal/user/delivery/http"
	userrepo "github.com/mafzaidi/elog/internal/user/repository"
	useruc "github.com/mafzaidi/elog/internal/user/usecase"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	e.Use(emdw.CORSWithConfig(emdw.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
		AllowMethods:     []string{echo.POST, echo.GET, echo.OPTIONS},
	}))

	v1 := e.Group("/api/v1")
	health := v1.Group("/health")
	events := v1.Group("/events")
	authPublic := v1.Group("/auth")

	private := e.Group("/private/api/v1")
	private.Use(middleware.JWTAuthMiddleware)
	menus := private.Group("/menus")
	users := private.Group("/users")
	services := private.Group("/services")
	accounts := private.Group("/accounts")
	authPrivate := private.Group("/auth")

	eventRepo := eventrepo.NewEventRepository(s.db.DB)
	authRepo := authrepo.NewAuthRepository(s.db.DB)
	userRepo := userrepo.NewUserRepository(s.db.DB)
	menuRepo := menurepo.NewMenuRepository(s.db.DB)
	serviceRepo := servicerepo.NewServiceRespository(s.db.DB)
	accountRepo := accountrepo.NewAccountRepository(s.db.DB)

	eventUC := eventuc.NewEventUseCase(eventRepo)
	authUC := authuc.NewAuthUseCase(authRepo)
	menuUC := menuuc.NewMenuUseCase(menuRepo)
	userUC := useruc.NewUserUseCase(userRepo)
	serviceUC := serviceuc.NewServiceUseCase(
		serviceRepo,
		accountRepo,
	)
	accountUC := accountuc.NewAccountUseCase(
		accountRepo,
		serviceRepo,
		userRepo,
	)

	eventHandler := eventhttp.NewEventHandler(eventUC)
	eventhttp.MapRoutes(events, eventHandler)

	authHandler := authhttp.NewAuthHandler(authUC)
	authhttp.MapPublicRoutes(authPublic, authHandler, s.cfg)
	authhttp.MapPrivateRoutes(authPrivate, authHandler)

	menuHandler := menuhttp.NewMenuHandler(menuUC)
	menuhttp.MapRoutes(menus, menuHandler)

	userHandler := userhttp.NewUserHandler(userUC)
	userhttp.MapRoutes(users, userHandler)

	serviceHandler := servicehttp.NewServiceHandler(serviceUC)
	servicehttp.MapRoutes(services, serviceHandler)

	accountHandler := accounthttp.NewAccountHandler(accountUC)
	accounthttp.MapRoutes(accounts, accountHandler)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}{
			Message: "Healthy",
			Status:  http.StatusOK,
		})
	})

	return nil
}
