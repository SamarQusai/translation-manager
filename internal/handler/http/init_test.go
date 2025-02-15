package http

import (
	"os"
	"testing"
	"translationManager/internal/context"
	"translationManager/internal/service"
)

const OpenAiApiKey = "OPEN_AI_API_KEY"
const OpenAiApiKeyValue = ""
const EmptyOpenAiApiKeyValue = ""

var serviceContext context.ServiceContext
var translationService service.Service

func TestMain(testing *testing.M) {
	os.Setenv(OpenAiApiKey, OpenAiApiKeyValue)
	serviceContext = context.NewTranslationServiceContext()
	translationService = service.NewTranslationService(serviceContext)
	testing.Run()
}
