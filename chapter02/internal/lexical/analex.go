package lexical

import (
	"bytes"
)

type AnaLex interface {
	Next() Symbol
}

func NewAnaLex(input *bytes.Buffer) AnaLex {
	return newAnaLex(input, defaultKeywords)
}

func newAnaLex(input *bytes.Buffer, ss []Symbol) AnaLex {
	st := newSymbolTable(ss)
	return &anaLex{input: input, st: st}
}

type anaLex struct {
	input   *bytes.Buffer
	st      symbolTable
	numLine int
}

func (a *anaLex) Next() Symbol {
	var c char

	for {
		c = a.getChar()

		if c.isBlank() {

		} else if c.isNewLine() {
			a.numLine++
		} else if c.isDigit() {
			a.unGetChar()
			num := a.readNumber()
			return Symbol{
				lexeme:    "",
				value:     num,
				component: NUM,
			}
		} else if c.isAlpha() {
			var lexeme string
			for c.isAlNum() {
				lexeme = lexeme + string(c)
				c = a.getChar()
			}

			if c != charEOF {
				a.unGetChar()
			}

			symbol, found := a.st.search(lexeme)
			if !found {
				symbol = a.st.insert(lexeme, ID)
			}
			return symbol
		} else if c.isEOF() {
			return Symbol{component: END}
		} else {
			return Symbol{component: Component(c)}
		}
	}
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
