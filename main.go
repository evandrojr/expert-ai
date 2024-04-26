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
	doublePrompt("prompt.txt", "answer_poe.txt", "answer_chatgpt.txt", "combined_prompt.txt")
	// tool.JoinFiles("prompt.txt", "prompts/compare_answers.txt", "answers/prompts_combined_answers.txt")
	doublePrompt("prompts/combined_prompt.txt", "2ndAnswer_poe.txt", "2ndAnswer_chatgpt.txt", "2nCombined_prompt.txt")
}

func doublePrompt(promptFile, answer1File, answer2File, combinedPrompt string) {

	prompt, err := tool.ReadFile(promptFile)
	if err != nil {
		panic(err)
	}
	headerForAnswer1 := "Resposta 1:\n\n"
	headerForAnswer2 := "\n\nResposta 2:\n\n"

	var claude3 artificialintelligence.Claude3
	answerClaude := sendPrompt(claude3, prompt)

	err = tool.WriteFile("answers/"+answer1File, answerClaude)
	if err != nil {
		panic(err)
	}

	var chatgpt artificialintelligence.Chatgpt
	answerChatgpt := sendPrompt(chatgpt, prompt)

	err = tool.WriteFile("answers/"+answer2File, answerChatgpt)
	if err != nil {
		panic(err)
	}

	basedOntheQuestion := `Baseado na pergunta:
	
	`

	compareAnswers, err := tool.ReadFile(promptFile)
	if err != nil {
		panic(err)
	}

	combinedPromptText := basedOntheQuestion + prompt + "\n\n" +
		compareAnswers +
		headerForAnswer1 + answerClaude +
		headerForAnswer2 + answerChatgpt

	err = tool.WriteFile("prompts/"+combinedPrompt, combinedPromptText)
	if err != nil {
		panic(err)
	}
	// tool.JoinFiles("answers/"+answer1File, "answers/"+answer2File, "answers/"+combinedAnswers)

}
