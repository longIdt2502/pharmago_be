package utils

import (
	"github.com/rainycape/unidecode"
	"regexp"
	"strings"
)

func ExtractFirstLetters(input string) string {
	stringRaw := unidecode.Unidecode(input)
	stringRaw = removeNonAlphanumeric(stringRaw)
	words := strings.Fields(stringRaw)
	result := ""

	for _, word := range words {
		if len(word) > 0 {
			result += string(word[0])
		}
	}

	return strings.ToUpper(result)
}

func removeNonAlphanumeric(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9 ]+")

	cleaned := reg.ReplaceAllString(input, "")

	return cleaned
}
