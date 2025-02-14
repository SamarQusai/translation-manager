package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"translationManager/internal/model"
)

func TestSuccessCase_ValidateStruct(testing *testing.T) {

	dialogues := make([]model.Dialogue, 0)
	dialogues = append(dialogues, model.Dialogue{
		Speaker:  "John",
		Time:     "00:00:04",
		Sentence: "Good morning",
	})
	request := model.DialogueRequest{Dialogue: dialogues}
	validator := Validator{}
	errors := validator.ValidateStruct(request)
	assert.Len(testing, errors, 0)
}

func TestFailureEmptySentenceCase_ValidateStruct(testing *testing.T) {

	dialogues := make([]model.Dialogue, 0)
	dialogues = append(dialogues, model.Dialogue{
		Speaker:  "John",
		Time:     "00:00:04",
		Sentence: "",
	})
	request := model.DialogueRequest{Dialogue: dialogues}
	validator := Validator{}
	errors := validator.ValidateStruct(request)
	assert.Len(testing, errors, 1)
	assert.Equal(testing, errors[0].Field, "Dialogue")
}

func TestFailureEmptyNameCase_ValidateStruct(testing *testing.T) {

	dialogues := make([]model.Dialogue, 0)
	dialogues = append(dialogues, model.Dialogue{
		Speaker:  "",
		Time:     "00:00:04",
		Sentence: "Good morning",
	})
	request := model.DialogueRequest{Dialogue: dialogues}
	validator := Validator{}
	errors := validator.ValidateStruct(request)
	assert.Len(testing, errors, 1)
	assert.Equal(testing, errors[0].Field, "Dialogue")
}
