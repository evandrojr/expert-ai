package tool

import (
	"testing"
)

func TestExtractPlainText2(t *testing.T) {
	// Criar um arquivo de entrada de teste com conteúdo HTML
	// inputFileName := "test_input.html"
	// inputContent := "<html><body><h1>Título</h1><p>Este é um <b>teste</b> de HTML.</p></body></html>"
	// err := os.WriteFile(inputFileName, []byte(inputContent), fs.FileMode(0644))
	// if err != nil {
	// 	t.Fatalf("Erro ao criar arquivo de entrada de teste: %v", err)
	// }
	// defer os.Remove(inputFileName)

	// // Definir o nome do arquivo de saída de teste
	// outputFileName := "test_output.txt"
	// defer os.Remove(outputFileName)

	// Executar o método ExtractPlainText
	err := ExtractPlainText("answer.txt", "saida.txt")
	if err != nil {
		t.Fatalf("Erro ao executar ExtractPlainText: %v", err)
	}

	// // Ler o conteúdo do arquivo de saída gerado
	// outputContent, err := os.ReadFile(outputFileName)
	// if err != nil {
	// 	t.Fatalf("Erro ao ler o arquivo de saída gerado: %v", err)
	// }

	// // Verificar se o texto limpo está correto
	// expectedOutput := "Título\nEste é um teste de HTML."
	// if string(outputContent) != expectedOutput {
	// 	t.Errorf("Texto limpo não corresponde ao esperado. Esperado: %s, Obtido: %s", expectedOutput, string(outputContent))
	// }
}
