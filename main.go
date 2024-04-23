package main

import (
	"fmt"
	"os"
	"regexp"

	"io/fs"

	"github.com/evandrojr/expert-ai/poe"
)

func main() {
	fmt.Println("google-chrome --remote-debugging-port=9222")

	prompt, err := readFile("prompt.txt")
	if err != nil {
		panic(err)
	}

	answer, err := poe.Prompt(prompt)
	if err != nil {
		panic(err)
	}

	err = writeFile("answer.txt", answer)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
	// fmt.Println("google-chrome --remote-debugging-port=9222
}

func readFile(filename string) (string, error) {
	// Read the entire file into a byte slice.
	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string.
	str := string(b)

	return str, nil
}

func writeFile(filename string, data string) error {
	// Create the file.
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	// Write the data to the file.
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}

	// Close the file.
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func ExtractPlainText(inputFile, outputFile string) error {
	// Ler o arquivo HTML
	htmlContent, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("erro ao ler o arquivo de entrada: %v", err)
	}

	// Definir expressão regular para encontrar tags HTML e símbolos HTML
	htmlTagsRegex := regexp.MustCompile("<[^>]*>|&[^;]+;")

	// Remover tags HTML e símbolos HTML do conteúdo HTML
	cleanText := htmlTagsRegex.ReplaceAllString(string(htmlContent), "")

	// Escrever o texto limpo no arquivo de saída
	err = os.WriteFile(outputFile, []byte(cleanText), fs.FileMode(0644))
	if err != nil {
		return fmt.Errorf("rro ao escrever no arquivo de saída: %v", err)
	}

	fmt.Println("Texto limpo foi salvo em", outputFile)
	return nil
}
