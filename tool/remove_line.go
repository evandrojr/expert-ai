package tool

import "strings"

func RemoveLine(str string, lineToRemove int) string {
	lines := strings.Split(str, "\n")
	if lineToRemove < 0 || lineToRemove >= len(lines) {
		return str // Linha inválida, retorna a string original
	}

	newLines := make([]string, 0, len(lines)-1)
	for i, line := range lines {
		if i != lineToRemove {
			newLines = append(newLines, line)
		}
	}

	return strings.Join(newLines, "\n")
}
