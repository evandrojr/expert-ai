package tool

import (
	"fmt"
	"testing"
)

func TestRemoveRubbishFromBeginning(t *testing.T) {
	fullText := `Olá sou o começo
				 quem nasceu primeiro o ovo
				 ou a galinha?`

	fmt.Println(RemoveRubbishFromBeginning(fullText, `quem nasceu primeiro`))
	t.Log("oi")
}
