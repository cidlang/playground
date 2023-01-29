package syntactical

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"compiler/internal/lexical"
)

func TestAnaSyn_Emit(t *testing.T) {
	tests := []struct {
		input  lexical.Symbol
		output string
	}{
		{input: lexical.NewSymbol('+', "", 0), output: "+\n"},
		{input: lexical.NewSymbol('-', "", 0), output: "-\n"},
		{input: lexical.NewSymbol('*', "", 0), output: "*\n"},
		{input: lexical.NewSymbol('/', "", 0), output: "/\n"},
		{input: lexical.NewSymbol(lexical.DIV, "", 0), output: "DIV\n"},
		{input: lexical.NewSymbol(lexical.MOD, "", 0), output: "MOD\n"},
		{input: lexical.NewSymbol(lexical.NUM, "", 1234), output: "1234\n"},
		{input: lexical.NewSymbol(lexical.ID, "fooBar", 0), output: "fooBar\n"},
	}

	for _, test := range tests {
		t.Run(test.input.GetLexeme(), func(t *testing.T) {
			var o bytes.Buffer
			as := anaSyn{output: &o}

			as.emit(test.input)

			assert.Equal(t, test.output, o.String())
		})
	}
}

func TestAnaSyn_EmitUnknown(t *testing.T) {
	var o bytes.Buffer
	as := anaSyn{output: &o}
	input := lexical.NewSymbol('€', "euro", 41)

	as.emit(input)

	assert.Equal(t, "component €, value 41\n", o.String())
}
