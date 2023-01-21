package lexical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBlank(t *testing.T) {
	tests := []struct {
		input  char
		output bool
	}{
		{input: ' ', output: true},
		{input: '\t', output: true},
		{input: charEOF, output: false},
		{input: 'a', output: false},
		{input: '\n', output: false},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			assert.Equal(t, test.output, test.input.isBlank())
		})
	}
}

func TestIsNewLine(t *testing.T) {
	tests := []struct {
		input  char
		output bool
	}{
		{input: '\n', output: true},
		{input: '\t', output: false},
		{input: charEOF, output: false},
		{input: 'a', output: false},
		{input: ' ', output: false},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			assert.Equal(t, test.output, test.input.isNewLine())
		})
	}
}

func TestIsDigit(t *testing.T) {
	tests := []struct {
		input  char
		output bool
	}{
		{input: '0', output: true},
		{input: '5', output: true},
		{input: '9', output: true},
		{input: charEOF, output: false},
		{input: 'a', output: false},
		{input: ' ', output: false},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			assert.Equal(t, test.output, test.input.isDigit())
		})
	}
}

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input  char
		output bool
	}{
		{input: 'a', output: true},
		{input: 'j', output: true},
		{input: 'z', output: true},
		{input: 'A', output: true},
		{input: 'J', output: true},
		{input: 'Z', output: true},
		{input: charEOF, output: false},
		{input: '0', output: false},
		{input: ' ', output: false},
		{input: '.', output: false},
		{input: '<', output: false},
		{input: '?', output: false},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			assert.Equal(t, test.output, test.input.isAlpha())
		})
	}
}

func TestIsAlNum(t *testing.T) {
	tests := []struct {
		input  char
		output bool
	}{
		{input: 'a', output: true},
		{input: 'j', output: true},
		{input: 'z', output: true},
		{input: 'A', output: true},
		{input: 'J', output: true},
		{input: 'Z', output: true},
		{input: '0', output: true},
		{input: '5', output: true},
		{input: '9', output: true},
		{input: charEOF, output: false},
		{input: ' ', output: false},
		{input: '.', output: false},
		{input: '<', output: false},
		{input: '?', output: false},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			assert.Equal(t, test.output, test.input.isAlNum())
		})
	}
}

func TestIsEOF(t *testing.T) {
	tests := []struct {
		input  char
		output bool
	}{
		{input: charEOF, output: true},
		{input: '\n', output: false},
		{input: 'a', output: false},
		{input: ' ', output: false},
	}

	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			assert.Equal(t, test.output, test.input.isEOF())
		})
	}
}
