package lexical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbol_GetComponent(t *testing.T) {
	s := symbol{component: END}

	assert.Equal(t, END, s.GetComponent())
}

func TestSymbol_GetLexeme(t *testing.T) {
	s := symbol{lexeme: "foo-bar"}

	assert.Equal(t, "foo-bar", s.GetLexeme())
}

func TestSymbol_GetValue(t *testing.T) {
	s := symbol{value: 1234}

	assert.Equal(t, 1234, s.GetValue())
}
