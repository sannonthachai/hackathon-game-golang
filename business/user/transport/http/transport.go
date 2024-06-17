package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/sannonthachai/find-the-hidden-backend/business/user/transport"
	"gitlab.com/sannonthachai/find-the-hidden-backend/model"
	modelUser "gitlab.com/sannonthachai/find-the-hidden-backend/model/user"
)

type httpRoute struct {
	handler   transport.Handler
	appConfig model.Config
}

func NewUserHTTPRoute(userHandler transport.Handler, appConfig model.Config) httpRoute {
	return httpRoute{
		handler:   userHandler,
		appConfig: appConfig,
	}
}

func (h httpRoute) RoutePublic(e *echo.Group) {
	api := e.Group("/api/v1")

	api.POST("/register", h.handler.Register)
	api.POST("/login", h.handler.Login)
}

func (h httpRoute) RoutePrivate(e *echo.Group) {
	api := e.Group("/api/v1")

	config := middleware.JWTConfig{
		Claims:     &modelUser.JwtCustomClaims{},
		SigningKey: []byte(h.appConfig.Secret),
	}
	api.Use(middleware.JWTWithConfig(config))

	user := api.Group("/user")
	user.PUT("/point", h.handler.UpdateUserPoint)
	user.GET("/point", h.handler.GetUserPoint)
	user.GET("/chapter-point", h.handler.GetUserPointByChapter)
	user.GET("/chapter-leaderboard", h.handler.GetLeaderBoardByChapter)
	user.GET("/leaderboard", h.handler.GetLeaderBoard)
}
