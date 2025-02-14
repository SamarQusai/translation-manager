package service

import (
	"translationManager/internal/model"
	"translationManager/internal/utils"
)

func (receiver *Service) Translate(request model.DialogueRequest) (*model.DialogueResponse, error) {
	texts := prepareToBeTranslatedArray(request)
	if len(texts) == 0 {
		return nil, model.ErrNoTextsToBeTranslated
	}

	translatedTexts, translationError := receiver.translatorGateway.Translate(texts, model.EnglishLanguage)
	if translationError != nil {
		return nil, translationError
	}

	response := buildResponse(request, translatedTexts)
	return response, nil
}

func prepareToBeTranslatedArray(request model.DialogueRequest) []string {
	toBeTranslated := make([]string, 0)
	for _, dialogue := range request.Dialogue {
		if utils.IsTextInArabic(dialogue.Sentence) {
			toBeTranslated = append(toBeTranslated, dialogue.Sentence)
		}
	}

	return toBeTranslated
}

func buildResponse(request model.DialogueRequest, translatedTexts []string) *model.DialogueResponse {
	var response model.DialogueResponse
	responseDialogues := make([]model.DialogueView, 0)
	for index, dialogue := range request.Dialogue {
		responseDialogues = append(responseDialogues, model.DialogueView{
			Speaker:  dialogue.Speaker,
			Time:     dialogue.Time,
			Sentence: translatedTexts[index],
		})
	}
	response.Dialogue = responseDialogues
	return &response
}
