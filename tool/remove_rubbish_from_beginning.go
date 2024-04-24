package tool

import (
	s "strings"
)

func RemoveRubbishFromBeginning(fullText string, prompt string) string {
	index := s.Index(fullText, prompt)
	if index == -1 {
		return fullText
	}
	return fullText[index:]
}
