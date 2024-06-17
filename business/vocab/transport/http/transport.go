package http

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/sannonthachai/find-the-hidden-backend/business/vocab/transport"
	"gitlab.com/sannonthachai/find-the-hidden-backend/model"
)

type httpRoute struct {
	handler   transport.Handler
	appConfig model.Config
}

func NewVocabHTTPRoute(vocabHandler transport.Handler, appConfig model.Config) httpRoute {
	return httpRoute{
		handler:   vocabHandler,
		appConfig: appConfig,
	}
}

func (h httpRoute) RoutePublic(e *echo.Group) {
	api := e.Group("/api/v1")

	api.GET("/vocab", h.handler.GetVocabByChapter)
	api.GET("/text-to-speech", h.handler.TextToSpeech)
}
