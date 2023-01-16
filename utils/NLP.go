package utils

import (
	"github.com/rylans/getlang"
)

type NLP struct {
	Text       string
	Language   string
	Confidence float64
}

func Proc(text string) (m NLP) {
	info := getlang.FromString(text)
	m.Text = text
	m.Language = info.LanguageCode()
	m.Confidence = info.Confidence()
	return
}
