package activities

import (
	"context"
)

var hardCodedHelloTranslations = map[string]string{
	"English":  "Hello",
	"Spanish":  "Hola",
	"French":   "Bonjour",
	"German":   "Hallo",
	"Italian":  "Ciao",
	"Japanese": "Konnichiwa",
	"Russian":  "Privet",
	"Chinese":  "Nǐ hǎo",
	"Arabic":   "Marhaba",
	"Hindi":    "Namaste",
}

type LanguageNotFound struct{}

func (e *LanguageNotFound) Error() string {
	return "Language not found"
}

func TranslateThisTo(_ context.Context, language string) (string, error) {
	translation, ok := hardCodedHelloTranslations[language]
	if !ok {
		return "", &LanguageNotFound{}
	}
	return translation, nil
}
