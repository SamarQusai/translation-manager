package utilities

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"time"
	"translationManager/internal/model"
)

type Validator struct {
	Config       map[string]interface{}
	ErrorDetails []Error
}
type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v *Validator) ValidateStruct(s interface{}) []Error {
	validate := validator.New()
	validate.RegisterValidation("validateConversations", validateConversations)

	err := validate.Struct(s)
	if err != nil {
		for _, validationError := range err.(validator.ValidationErrors) {
			errorDetails := Error{}
			errorDetails.Field = validationError.Field()
			errorDetails.Message = "Field " + validationError.Field() + " should be " + validationError.Tag()
			if validationError.Param() != "" {
				errorDetails.Message += " " + validationError.Param()
			}
			v.ErrorDetails = append(v.ErrorDetails, Error{
				Message: errorDetails.Message,
				Field:   errorDetails.Field,
			})
		}
	}

	return v.ErrorDetails
}

func validateConversations(fl validator.FieldLevel) bool {

	data := fl.Field().Interface()
	var conversations []model.Dialogue
	conversations = data.([]model.Dialogue)

	if len(conversations) == 0 {
		return false
	}

	for _, conversation := range conversations {
		if strings.TrimSpace(conversation.Speaker) == "" {
			return false
		}

		if !isValidTimeFormat(conversation.Time) {
			return false
		}

		if strings.TrimSpace(conversation.Sentence) == "" {
			return false
		}
	}

	return true
}

func isValidTimeFormat(timeString string) bool {
	const layout = "15:04:05"
	_, err := time.Parse(layout, timeString)
	if err != nil {
		return false
	}
	return true
}
