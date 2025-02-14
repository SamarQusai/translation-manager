package utils

import (
	logger "github.com/sirupsen/logrus"
	"regexp"
)

var arabicRegex = `^[\x{0600}-\x{06FF}\x{0750}-\x{077F}\x{08A0}-\x{08FF}\x{FB50}-\x{FDFF}\x{FE70}-\x{FEFF}\s]+$`

func IsTextInArabic(text string) bool {
	result, err := regexp.MatchString(arabicRegex, text)
	if err != nil {
		logger.Error("IsTextInArabic - matching regex error. error: ", err.Error(), " text: ", text, " regex:", arabicRegex)
		return false
	}
	return result
}
