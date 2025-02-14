package http

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"translationManager/internal/model"
	"translationManager/internal/pkg"
)

func (receiver *Handler) Translate(c *gin.Context) {
	var request model.DialogueRequest
	bindingError := c.ShouldBindJSON(&request)
	if bindingError != nil {
		logger.Warning("Invalid request. request: ", request, " error: ", bindingError.Error())
		model.BuildHttpResponse(c, http.StatusBadRequest, model.InvalidRequestMessage, bindingError.Error())
		return
	}

	validator := pkg.Validator{}
	errors := validator.ValidateStruct(request)
	if len(errors) > 0 {
		model.BuildHttpResponse(c, http.StatusBadRequest, model.InvalidRequestMessage, errors)
		return
	}

	if len(request.Dialogue) == 0 {
		model.BuildHttpResponse(c, http.StatusBadRequest, model.EmptyConversationMessage, nil)
		return
	}

	response, translationError := receiver.translatorService.Translate(request)
	if translationError != nil {
		if translationError.Error() == model.ErrNoTextsToBeTranslated.Error() {
			model.BuildHttpResponse(c, http.StatusOK, model.SuccessNoTextsNeedsTranslation, request)
			return
		}

		model.BuildHttpResponse(c, http.StatusInternalServerError, model.ServerError, nil)
		return
	}

	model.BuildHttpResponse(c, http.StatusOK, model.SuccessMessage, response)
	return
}
