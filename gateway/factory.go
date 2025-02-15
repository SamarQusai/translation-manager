package gateway

import (
	logger "github.com/sirupsen/logrus"
	"translationManager/gateway/open_ai"
	"translationManager/internal/context"
)

const OpenAI = "open_ai"

func NewTranslationGateway(configuration *context.Configuration, translationService string) Interface {
	var currentTranslationService Interface
	if translationService == OpenAI {
		currentTranslationService = open_ai.NewOpenAi(configuration)
	} else {
		logger.Panic("NewTranslationGateway - translation service type not handled. type: ", translationService)
	}

	return currentTranslationService
}
