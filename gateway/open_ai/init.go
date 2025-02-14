package open_ai

import (
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"translationManager/internal/context"
)

type OpenAI struct {
	client *openai.Client
}

func NewOpenAi(configuration *context.Configuration) *OpenAI {
	client := openai.NewClient(option.WithAPIKey(configuration.OpenAiApiKey()))
	return &OpenAI{
		client: client,
	}
}
