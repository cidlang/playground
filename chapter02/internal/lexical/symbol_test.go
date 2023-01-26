package lexical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetComponent(t *testing.T) {
	s := Symbol{component: END}

	assert.Equal(t, END, s.GetComponent())
}

func TestGetLexeme(t *testing.T) {
	s := Symbol{lexeme: "foo-bar"}

	assert.Equal(t, "foo-bar", s.GetLexeme())
}

func TestGetValue(t *testing.T) {
	s := Symbol{value: 1234}

	assert.Equal(t, 1234, s.GetValue())
}
