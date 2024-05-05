package tool

import (
	"fmt"
	"html"
	"io/fs"
	"os"
	"regexp"
)

func ExtractPlainText(inputFile, outputFile string) error {
	// Ler o arquivo HTML
	htmlContent, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo de entrada: %v", err)
	}

	// Definir expressão regular para encontrar tags HTML
	htmlTagsRegex := regexp.MustCompile("<[^>]*>")
	// Converter entidades HTML para caracteres legíveis
	cleanText := html.UnescapeString(string(htmlContent))

	// Remover tags HTML do conteúdo HTML
	cleanText = htmlTagsRegex.ReplaceAllString(cleanText, " ")

	// Escrever o texto limpo no arquivo de saída
	err = os.WriteFile(outputFile, []byte(cleanText), fs.FileMode(0644))
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo de saída: %v", err)
	}

	fmt.Println("Texto limpo foi salvo em", outputFile)
	return nil
}
