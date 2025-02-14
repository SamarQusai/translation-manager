package open_ai

import (
	"context"
	"errors"
	"fmt"
	"github.com/openai/openai-go"
	logger "github.com/sirupsen/logrus"
)

func (receiver *OpenAI) Translate(texts []string, targetLanguage string) ([]string, error) {

	messages := make([]openai.ChatCompletionMessageParamUnion, 0)
	for _, text := range texts {
		prompt := fmt.Sprintf("Translate this text to %s: %s", targetLanguage, text)
		messages = append(messages, openai.UserMessage(prompt))

	}
	request := openai.ChatCompletionNewParams{
		Messages: openai.F(messages),
		Model:    openai.F(openai.ChatModelGPT4o),
	}

	chatCompletion, err := receiver.client.Chat.Completions.New(context.TODO(), request)

	if err != nil {
		logger.Error("Error while sending translation request to open ai. Error: ", err.Error(), " request: ", request)
		return nil, err
	}

	if len(chatCompletion.Choices) == 0 {
		logger.Error("Received no translation options from open ai. request: ", request, " response: ", chatCompletion.Choices)
		return nil, errors.New("no translation received")
	}

	translatedTexts := make([]string, 0)
	for _, choice := range chatCompletion.Choices {
		translatedTexts = append(translatedTexts, choice.Message.Content)
	}

	return translatedTexts, nil
}
