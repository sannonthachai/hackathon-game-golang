package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subosito/gotenv"
	"gitlab.com/sannonthachai/find-the-hidden-backend/config"
	"gitlab.com/sannonthachai/find-the-hidden-backend/model"

	vocabRepo "gitlab.com/sannonthachai/find-the-hidden-backend/business/vocab/repository"
	vocabService "gitlab.com/sannonthachai/find-the-hidden-backend/business/vocab/service"
	vocabTransport "gitlab.com/sannonthachai/find-the-hidden-backend/business/vocab/transport"
	vocabHTTPRoute "gitlab.com/sannonthachai/find-the-hidden-backend/business/vocab/transport/http"

	userRepo "gitlab.com/sannonthachai/find-the-hidden-backend/business/user/repository"
	userService "gitlab.com/sannonthachai/find-the-hidden-backend/business/user/service"
	userTransport "gitlab.com/sannonthachai/find-the-hidden-backend/business/user/transport"
	userHTTPRoute "gitlab.com/sannonthachai/find-the-hidden-backend/business/user/transport/http"
)

func init() {
	gotenv.Load()
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	appConfig := model.Config{
		Port:      os.Getenv("PORT"),
		MySQLConn: os.Getenv("MYSQL_CONN"),
		Secret:    os.Getenv("SECRET"),
	}

	config.InitDB(appConfig.MySQLConn)
	mysqlDB := config.GetDB()
	defer config.CloseDB()

	initEndpoint(e, mysqlDB, appConfig)

	e.GET("/healthz", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})

	go func() {
		if err := e.Start(":" + appConfig.Port); err != nil {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func initEndpoint(e *echo.Echo, mysqlDB *gorm.DB, appConfig model.Config) {
	routePublic := e.Group("/public")
	routePrivate := e.Group("/private")

	vocabRepository := vocabRepo.NewVocabRepository(mysqlDB)
	vocabService := vocabService.NewVocabService(vocabRepository)
	vocabHandler := vocabTransport.NewVocabHandler(vocabService)

	userRepository := userRepo.NewUserRepository(mysqlDB)
	userService := userService.NewUserService(userRepository, appConfig)
	userHandler := userTransport.NewUserHandler(userService)

	vocabHTTPRoute := vocabHTTPRoute.NewVocabHTTPRoute(vocabHandler, appConfig)
	userHTTPRoute := userHTTPRoute.NewUserHTTPRoute(userHandler, appConfig)

	vocabHTTPRoute.RoutePublic(routePublic)
	userHTTPRoute.RoutePrivate(routePrivate)
	userHTTPRoute.RoutePublic(routePublic)

}
