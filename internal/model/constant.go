package model

import "errors"

var ErrNoTextsToBeTranslated = errors.New("no texts found to be translated")

const (
	SuccessMessage                 = "Success"
	SuccessNoTextsNeedsTranslation = "No texts needs translations"
	InvalidRequestMessage          = "Invalid request"
	EmptyConversationMessage       = "Empty conversation"
	ServerError                    = "Server error"

	EnglishLanguage = "English"
)
