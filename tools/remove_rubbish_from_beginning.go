package tools

import (
	s "strings"
)

func RemoveRubbishFromBeginning(fullText string, prompt string) string {
	index := s.Index(fullText, prompt)
	if index == -1 {
		return fullText
	}
	index = index + len(prompt)
	return fullText[index:]
}
