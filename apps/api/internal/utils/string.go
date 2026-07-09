package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CapitalizeEachWord(s string) string {
	// Split into words
	words := strings.Fields(s)

	// Create a title-casing transformer for the desired language
	caser := cases.Title(language.Und) // Und = undefined, works for general Unicode

	for i, word := range words {
		words[i] = caser.String(word)
	}

	// Join back into a single string
	return strings.Join(words, " ")
}
