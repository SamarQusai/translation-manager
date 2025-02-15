package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessCase_IsTextInArabic(testing *testing.T) {
	examples := []string{
		"مرحبا",
		"اللغة العربية",
		"مَرْحَبًا",
		"٠١٢٣٤٥٦٧٨٩",
	}
	for _, example := range examples {
		result := IsTextInArabic(example)
		assert.Equal(testing, result, true)
	}
}

func TestNotArabicCase_IsTextInArabic(testing *testing.T) {
	examples := []string{
		"Hey",
		"Hello",
		"",
		"Good morning",
		"01939839",
	}
	for _, example := range examples {
		result := IsTextInArabic(example)
		assert.Equal(testing, result, false)
	}
}
