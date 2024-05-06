package artificialintelligence

import (
	"fmt"
	"testing"
)

func TestTimeout(t *testing.T) {
	question := "Qual é o carro mais rápido do mundo"
	answer, err := PromptBrowserChatGpt(question)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(answer)
}
