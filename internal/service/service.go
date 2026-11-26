package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

const alphabet = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

func ConvertText(text string) string {
	isText := strings.ContainsAny(strings.ToLower(text), alphabet)
	if isText {
		return morse.DefaultConverter.ToMorse(text)
	} else {
		return morse.DefaultConverter.ToText(text)
	}
}
