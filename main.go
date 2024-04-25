package main

import (
	"fmt"

	artificialintelligence "github.com/evandrojr/expert-ai/artificial_intelligence"
	"github.com/evandrojr/expert-ai/tool"
)

func sendPrompt(ai artificialintelligence.ArtificialIntelligence, prompt string) string {
	var _ai = ai.Setup()
	answer, err := _ai.SubmitPrompt(prompt)
	if err != nil {
		panic(err)
	}
	return answer
}

func main() {
	fmt.Println("Run:")
	fmt.Println("killall google-chrome")
	fmt.Println("google-chrome --remote-debugging-port=9222")

	prompt, err := tool.ReadFile("prompt.txt")
	if err != nil {
		panic(err)
	}

	var claude3 artificialintelligence.Claude3
	answerClaude := sendPrompt(claude3, prompt)

	resposta1Cabecalho := "Resposta 1:\n\n"
	resposta2Cabecalho := "\n\nResposta 2:\n\n"

	var chatgpt artificialintelligence.Chatgpt

	answerChatgpt := sendPrompt(chatgpt, prompt)

	err = tool.WriteFile("answers/answer_chatgpt.txt", resposta1Cabecalho + answerChatgpt)
	if err != nil {
		panic(err)
	}
	fmt.Println(answerChatgpt)

	err = tool.WriteFile("answers/answer_poe.txt", resposta2Cabecalho + answerClaude)
	if err != nil {
		panic(err)
	}
	fmt.Println(answerClaude)

	err = tool.WriteFile("answers/answer_poe_sem_cabecalho.txt", answerClaude)

	if err != nil {
		panic(err)
	}
	fmt.Println(answerClaude)

	tool.JoinFiles("answers/answer_poe.txt", "answers/answer_chatgpt.txt", "answers/combined_answers.txt")

}
