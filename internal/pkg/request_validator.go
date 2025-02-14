package pkg

import (
	"github.com/go-playground/validator/v10"
	"strings"
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
	validate.RegisterValidation("validateTime", validateTime)

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

func validateTime(fl validator.FieldLevel) bool {
	return false
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

		if strings.TrimSpace(conversation.Sentence) == "" {
			return false
		}
	}

	return true
}
