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
	fmt.Println("killall chrome")
	fmt.Println("google-chrome --remote-debugging-port=9222")
	doublePrompt("prompt.txt", "answer_poe.txt", "answer_chatgpt.txt", "combined_answers.txt")
	tool.JoinFiles("prompt.txt", "prompts/compare_answers.txt", "answers/prompts_combined_answers.txt")

	doublePrompt("answers/"+"prompts_combined_answers.txt", "2ndAnswer_poe.txt", "2ndAnswer_chatgpt.txt", "2nCombined_answers.txt")
}

func doublePrompt(promptFile, answer1File, answer2File, combinedAnswers string) {

	prompt, err := tool.ReadFile(promptFile)
	if err != nil {
		panic(err)
	}
	headerForAnswer1 := "Resposta 1:\n\n"
	headerForAnswer2 := "\n\nResposta 2:\n\n"

	var claude3 artificialintelligence.Claude3
	answerClaude := sendPrompt(claude3, prompt)

	err = tool.WriteFile("answers/"+answer1File, headerForAnswer1+answerClaude)
	if err != nil {
		panic(err)
	}
	fmt.Println(answerClaude)

	var chatgpt artificialintelligence.Chatgpt
	answerChatgpt := sendPrompt(chatgpt, prompt)

	err = tool.WriteFile("answers/"+answer2File, headerForAnswer2+answerChatgpt)
	if err != nil {
		panic(err)
	}
	fmt.Println(answerChatgpt)

	// err = tool.WriteFile("answers/answer_poe_sem_cabecalho.txt", answerClaude)

	if err != nil {
		panic(err)
	}
	fmt.Println(answerClaude)

	tool.JoinFiles("answers/"+answer1File, "answers/"+answer2File, "answers/"+combinedAnswers)

}
