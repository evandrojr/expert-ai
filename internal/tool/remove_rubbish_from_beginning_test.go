package tool

import (
	"fmt"
	"os"
	"testing"

	"github.com/evandrojr/expert-ai/filesystem"
)

func TestRemoveRubbishFromBeginning(t *testing.T) {

	prompt := `Em qual lugar do mundo surgiram os ritmos samba de roda e o samba de caboclo?`
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
	fmt.Println(RemoveRubbishFromBeginning(dirt, prompt))
}
