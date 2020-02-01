package tokenizer

import (
	"bufio"
	"io"
	"unicode"
)

// Peek the next rune
func (r *Reader) Peek() (rune, error) {
	char, _, err := r.ReadRune()
	if err != nil {
		if err == io.EOF {
			panic("EOF")
		}
		panic(err)
	}
	r.UnreadRune()
	return char, err
}

// Consume the next rune
func (r *Reader) Consume() rune {
	char, _, err := r.ReadRune()
	if err != nil {
		panic(err)
	}
	return char
}

// Length of the input remaining
func (r *Reader) hasNext() bool {
	_, _, err := r.ReadRune()
	if err == io.EOF {
		return false
	} else if err != nil {
		// panic(err)
	}
	r.UnreadRune()
	return true
}

// Consume runes until condition is false.
func (r *Reader) consumeWhile(condition ConsumeCondition) string {
	result := ""

	for r.hasNext() {
		char, _ := r.Peek()
		if condition(char) {
			r.Consume()
			result += string(char)
		} else {
			break
		}
	}
	return result
}

// Consume and discard zero or more whitespace characters.
func (r *Reader) consumeWhitespace() {
	r.consumeWhile(func(ru rune) bool {
		return unicode.IsSpace(ru) || unicode.IsControl(ru)
	})
}

func (r *Reader) consumeIdentifier() string {
	s := ""
	c, err := r.Peek()
	if err != nil {
		if err == io.EOF {
			panic("EOF in consumeIdentifier")
		}
		panic(err)
	}

	// first index of an identifier must be a letter
	if unicode.IsLetter(c) {

		s += r.consumeWhile(func(ru rune) bool {
			return unicode.IsLetter(ru) || unicode.IsNumber(ru) || ru == '_'
		})
	}

	return s

}

func (r *Reader) consumeNumber() string {
	return r.consumeWhile(func(ru rune) bool {
		return unicode.IsDigit(ru)
	})
}

func (r *Reader) nextCharAdv() (rune, error) {
	char, _, err := r.ReadRune()
	// fmt.Println("r char: " + string(char))
	// check(err)
	r.UnreadRune()

	return char, err
}

func nextChar(buf *bufio.Reader) rune {
	char, _, err := buf.ReadRune()
	// fmt.Println("reading char: " + string(char))
	check(err)
	buf.UnreadRune()

	return char
}

// ConsumeCondition is a wrapper type for a function used as Comparators in Java
// are used to fulfill a condition requirement in the function consumeWhile()
type ConsumeCondition func(rune) bool

func check(e error) {
	if e != nil {
		panic(e)
	}
}
