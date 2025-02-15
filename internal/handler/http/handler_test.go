package http

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"translationManager/internal/model"
)

func TestInvalidRequest_Translate(testing *testing.T) {
	dialogues := make([]model.Dialogue, 0)
	request := model.DialogueRequest{Dialogue: dialogues}
	response, translationError := translationService.Translate(request)
	assert.NotNil(testing, translationError)
	assert.EqualError(testing, translationError, model.ErrNoTextsToBeTranslated.Error())
	assert.Nil(testing, response)
}

func TestUnauthorizedRequest_Translate(testing *testing.T) {
	os.Setenv(OpenAiApiKey, EmptyOpenAiApiKeyValue)
	dialogues := make([]model.Dialogue, 0)
	dialogues = append(dialogues, model.Dialogue{
		Speaker:  "John",
		Time:     "00:01:01",
		Sentence: "آهلا",
	})
	request := model.DialogueRequest{Dialogue: dialogues}
	response, translationError := translationService.Translate(request)
	assert.NotNil(testing, translationError)
	assert.Contains(testing, translationError.Error(), "You didn't provide an API key")
	assert.Nil(testing, response)
}
