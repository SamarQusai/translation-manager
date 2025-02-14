package main

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"translationManager/internal/context"
	"translationManager/router"
)

func main() {

	router := router.InitRouter()
	translationServiceContext := context.NewTranslationServiceContext()
	engine := gin.New()
	router.Install(engine)
	if err := engine.Run(":" + translationServiceContext.Configuration().HttpPort()); err != nil {
		logger.Panic("error when running gin engine. error: ", err)
	}
}
