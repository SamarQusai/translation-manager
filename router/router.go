package router

import (
	"github.com/gin-gonic/gin"
	"translationManager/internal/handler/http"
)

type Router struct {
	httpRequestHandler *http.Handler
}

func InitRouter() *Router {
	httpRequestHandler := http.NewHttpRequestHandler()
	return &Router{
		httpRequestHandler: httpRequestHandler,
	}
}

func (receiver *Router) Install(engine *gin.Engine) {
	apis := engine.Group("/api/v1/translation")
	{
		apis.POST("", receiver.httpRequestHandler.Translate)
	}
}
