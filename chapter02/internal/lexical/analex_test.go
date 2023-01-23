package lexical

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChar(t *testing.T) {
	var c char

	buffer := bytes.NewBufferString("bar")
	analyser := &anaLex{input: buffer}

	c = analyser.getChar()
	assert.Equal(t, char('b'), c)

	c = analyser.getChar()
	assert.Equal(t, char('a'), c)

	c = analyser.getChar()
	assert.Equal(t, char('r'), c)
}

func TestGetCharCases(t *testing.T) {
	var c char

	buffer := bytes.NewBufferString("FOo")
	analyser := &anaLex{input: buffer}

	c = analyser.getChar()
	assert.Equal(t, char('F'), c)

	c = analyser.getChar()
	assert.Equal(t, char('O'), c)

	c = analyser.getChar()
	assert.Equal(t, char('o'), c)
}

func TestGetCharSpecial(t *testing.T) {
	var c char

	buffer := bytes.NewBufferString(" \n\t")
	analyser := &anaLex{input: buffer}

	c = analyser.getChar()
	assert.Equal(t, char(' '), c)

	c = analyser.getChar()
	assert.Equal(t, char('\n'), c)

	c = analyser.getChar()
	assert.Equal(t, char('\t'), c)

	c = analyser.getChar()
	assert.Equal(t, charEOF, c)
}

func TestUnGetChar(t *testing.T) {
	var c char

	buffer := bytes.NewBufferString("foo")
	analyser := &anaLex{input: buffer}

	analyser.getChar()
	analyser.unGetChar()

	c = analyser.getChar()
	assert.Equal(t, char('f'), c)
}

func TestReadNumber(t *testing.T) {
	tests := []struct {
		input    string
		output   int
		nextChar char
	}{
		{input: "123foo", output: 123, nextChar: 'f'},
		{input: "foo123", output: 0, nextChar: 'f'},
		{input: "321 bar", output: 321, nextChar: ' '},
		{input: "", output: 0, nextChar: charEOF},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			buffer := bytes.NewBufferString(test.input)
			analyser := &anaLex{input: buffer}

			assert.Equal(t, test.output, analyser.readNumber())
			assert.Equal(t, test.nextChar, analyser.getChar())
		})
	}
}

func TestReadAlNum(t *testing.T) {
	tests := []struct {
		input    string
		output   string
		nextChar char
	}{
		{input: "123foo>", output: "123foo", nextChar: '>'},
		{input: "foo123>", output: "foo123", nextChar: '>'},
		{input: "321 bar", output: "321", nextChar: ' '},
		{input: "", output: "", nextChar: charEOF},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			buffer := bytes.NewBufferString(test.input)
			analyser := &anaLex{input: buffer}

			assert.Equal(t, test.output, analyser.readAlNum())
			assert.Equal(t, test.nextChar, analyser.getChar())
		})
	}
}
