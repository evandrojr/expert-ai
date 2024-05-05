package tool

import (
	"fmt"
	"os"
	"testing"

	"github.com/evandrojr/expert-ai/internal/filesystem"
)

func TestCleanPoeAnswerTest(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return
	}
	fmt.Println("Diretório de execução:", dir)

	dirt, err := filesystem.ReadFile(dir + "/texts/poe_dirt_answer.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(CleanPoeAnswer(dirt))
	t.Log("oi")
}
