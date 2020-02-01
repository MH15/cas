package tokenizer

import (
	"strings"
	"unicode"
)

type Reader struct {
	*strings.Reader
}

const DELIMITERS = "+-*/^%()"

func Tokens(in string) []string {
	tokens := make([]string, 0)

	stringReader := strings.NewReader(in)
	reader := Reader{stringReader}

	for reader.hasNext() {
		reader.consumeWhitespace()
		// fmt.Println("eeee")
		token, _ := reader.Peek()

		if strings.ContainsRune(DELIMITERS, token) {
			tokens = append(tokens, string(token))
			reader.Consume()
			continue
		} else {
			if unicode.IsLetter(token) {
				tokens = append(tokens, reader.consumeIdentifier())
				continue
			} else if unicode.IsDigit(token) {
				tokens = append(tokens, reader.consumeNumber())
				continue
			}
		}

		reader.Consume()

		tokens = append(tokens, string(token))
	}

	// fmt.Println(strings.Join(tokens, ", "))

	return tokens
}
