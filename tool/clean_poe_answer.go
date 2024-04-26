package tool

import (
	"regexp"
	"strings"
)

func CleanPoeAnswer(dirt string) string {

	clean := dirt
	regexesString := `Novo\s+chat
O Claude 3 Haiku da Anthropic supera
modelos em sua categoria de
inteligência em desempenho\,
velocidade e custo sem a necessidade de ajuste fino especializado\.
A janela de contexto foi reduzida para otimizar a velocidade e o custo\.
Para mensagens de contexto mais longas\,
experimente o Claude-3-Haiku-200k
experimente o Claude\-3\-Haiku\-200k\.
 ·  OFFICIAL
O valor dos pontos de computação está sujeito a alterações\.
Compartilhar
Solte arquivos aqui
Claude-3-Haiku
Operado por\s+@poe
\d{1,5} mil seguidores
O valor dos pontos de computação está sujeito a alterações
Ver mais   \d{1,4}
por mensagem
Ver detalhes
oficial
Informações do bot
· OFFICIAL`

	regexes := strings.Split(regexesString, "\n")

	for _, regex := range regexes {
		regexCompiled := regexp.MustCompile(`(?i)` + regex + `\s+`)
		clean = regexCompiled.ReplaceAllString(clean, "")
	}

	return clean
}
