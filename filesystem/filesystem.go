package filesystem

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadFile(filename string) (string, error) {
	// Read the entire file into a byte slice.
	filename = AdjustPathForWindows(filename)

	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to a string.
	str := string(b)

	return str, nil
}

func AdjustPathForWindows(filenameWithPath string) string {

	var pathSeparator string
	if runtime.GOOS == "windows" {
		pathSeparator = "\\"
	} else {
		pathSeparator = "/"
	}

	// Adjust the filename to use the correct path separator
	return strings.ReplaceAll(filenameWithPath, "/", pathSeparator)
}

func WriteFile(filename string, data string) error {

	filename = AdjustPathForWindows(filename)
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

	arquivo1 = AdjustPathForWindows(arquivo1)
	arquivo2 = AdjustPathForWindows(arquivo2)
	arquivoUnido = AdjustPathForWindows(arquivoUnido)

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

func CreateDirectoryIfNotExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func JoinPaths(parts ...string) string {
	separator := "/"
	if os.PathSeparator == '\\' {
		separator = "\\"
	}

	return filepath.Join(append([]string{separator}, parts...)...)
}

func GetHomeDir() (string, error) {
	// Obter o usuário atual
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	// Retornar o diretório home
	return user.HomeDir, nil
}
