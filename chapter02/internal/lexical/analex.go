package lexical

import "bytes"

const (
	charEOF char = 0
)

type AnaLex interface {
	Next() Component
}

type anaLex struct {
	input      *bytes.Buffer
	position   int
	numLine    int
	valcomplex int
}

func NewAnaLex(input *bytes.Buffer) AnaLex {
	return &anaLex{input: input}
}

func (a *anaLex) Next() Component {
	panic("not implemented")
}

func (a *anaLex) getChar() char {
	c, _ := a.input.ReadByte()
	return char(c)
}

func (a *anaLex) unGetChar() {
	_ = a.input.UnreadByte()
}

func (a *anaLex) readNumber() int {
	var result int

	c := a.getChar()
	for c.isDigit() {
		result = result*10 + int(c-'0')
		c = a.getChar()
	}
	a.unGetChar()
	return result
}

func (a *anaLex) readAlNum() string {
	var result string

	c := a.getChar()
	for c.isAlNum() {
		result = result + string(c)
		c = a.getChar()
	}
	a.unGetChar()
	return result
}
