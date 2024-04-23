package tool

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(filename string) (string, error) {
	// Read the entire file into a byte slice.
	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string.
	str := string(b)

	return str, nil
}

func WriteFile(filename string, data string) error {
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

func JoinFiles(arquivo1, arquivo2, arquivoUnido string) error {
	// Abrir o primeiro arquivo
	f1, err := os.Open(arquivo1)
	if err != nil {
		return err
	}
	defer f1.Close()

	// Abrir o segundo arquivo
	f2, err := os.Open(arquivo2)
	if err != nil {
		return err
	}
	defer f2.Close()

	// Criar o arquivo unido
	outFile, err := os.Create(arquivoUnido)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Copiar o conteúdo do primeiro arquivo para o arquivo unido
	if _, err := io.Copy(outFile, f1); err != nil {
		return err
	}

	// Copiar o conteúdo do segundo arquivo para o arquivo unido
	if _, err := io.Copy(outFile, f2); err != nil {
		return err
	}

	fmt.Println("Arquivos unidos com sucesso!")
	return nil
}
