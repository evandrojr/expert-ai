package tool

import (
	"bytes"
	"html"
	"regexp"
)

func HTMLToText(htmlContent string) string {

	// Substituir as tags <br> e </br> por \n
	re := regexp.MustCompile(`</?br\s*/?>`)
	text := re.ReplaceAllString(htmlContent, "\n")
	// Remover tags HTML
	re = regexp.MustCompile(`<[^>]+>`)
	text = re.ReplaceAllString(text, " ")

	// Traduzir entidades HTML
	text = html.UnescapeString(text)

	// Remover espaços em branco extras
	text = string(bytes.TrimSpace([]byte(text)))

	return text
}
