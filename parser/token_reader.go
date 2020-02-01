package parser

type Tokens struct {
	tokens []string
	index  int
}

func (t *Tokens) Consume() string {
	token := t.tokens[t.index]
	t.index++
	return token
}

func (t *Tokens) Peek() string {
	token := t.tokens[t.index]
	return token
}

func (t *Tokens) HasNext() bool {
	return t.index < len(t.tokens)
}
