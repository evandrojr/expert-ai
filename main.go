package main

import (
	"fmt"

	"github.com/evandrojr/expert-ai/driver/chatgpt"
	"github.com/evandrojr/expert-ai/driver/poe"
	"github.com/evandrojr/expert-ai/tool"
)

func main() {
	fmt.Println("run:")
	fmt.Println("google-chrome --remote-debugging-port=9222")

	prompt, err := tool.ReadFile("prompt.txt")
	if err != nil {
		panic(err)
	}

	answer_poe, err := poe.Prompt(prompt)
	if err != nil {
		panic(err)
	}

	resposta1Cabecalho := "Resposta 1:\n\n"
	resposta2Cabecalho := "\n\nResposta 2:\n\n"

	err = tool.WriteFile("answers/answer_poe.txt", resposta1Cabecalho+answer_poe)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer_poe)

	err = tool.WriteFile("answers/answer_poe_sem_cabecalho.txt", answer_poe)

	if err != nil {
		panic(err)
	}
	fmt.Println(answer_poe)

	answer_chatgpt, err := chatgpt.Prompt(prompt)
	if err != nil {
		panic(err)
	}

	err = tool.WriteFile("answers/answer_chatgpt.txt", resposta2Cabecalho+answer_chatgpt)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer_chatgpt)

	tool.JoinFiles("answers/answer_poe.txt", "answers/answer_chatgpt.txt", "answers/combined_answers.txt")

}
