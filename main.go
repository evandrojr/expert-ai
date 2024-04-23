package main

import (
	"fmt"

	"github.com/evandrojr/expert-ai/drivers/chatgpt"
	"github.com/evandrojr/expert-ai/drivers/poe"
	"github.com/evandrojr/expert-ai/tools"
)

func main() {
	fmt.Println("run:")
	fmt.Println("google-chrome --remote-debugging-port=9222")

	prompt, err := tools.ReadFile("prompt.txt")
	if err != nil {
		panic(err)
	}

	answer_poe, err := poe.Prompt(prompt)
	if err != nil {
		panic(err)
	}

	resposta1Cabecalho := "Resposta 1:\n\n"
	resposta2Cabecalho := "\n\nResposta 2:\n\n"

	err = tools.WriteFile("answers/answer_poe.txt", resposta1Cabecalho+answer_poe)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer_poe)

	answer_chatgpt, err := chatgpt.Prompt(prompt)
	if err != nil {
		panic(err)
	}

	err = tools.WriteFile("answers/answer_chatgpt.txt", resposta2Cabecalho+answer_chatgpt)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer_chatgpt)

	tools.JoinFiles("answers/answer_poe.txt", "answers/answer_chatgpt.txt", "answers/combined_answers.txt")

}
