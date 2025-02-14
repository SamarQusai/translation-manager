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
	//engine.Use(c.Options)
	apis := engine.Group("/api/v1/translation")
	{
		//apis.Use(c.checkHeaders)
		apis.POST("", receiver.httpRequestHandler.Translate)
	}
}
