package syntactical

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"compiler/internal/lexical"
)

func TestAnaSyn_FactorParenthesis(t *testing.T) {
	openParent := lexical.NewSymbol('(', "", 0)
	num := lexical.NewSymbol(lexical.NUM, "", 1234)
	closeParent := lexical.NewSymbol(')', "", 0)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(openParent, num, closeParent, end)

	a.factor()

	assert.Equal(t, "1234\n", a.output.String())
}

func TestAnaSyn_FactorNumber(t *testing.T) {
	num := lexical.NewSymbol(lexical.NUM, "", 1234)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num, end)

	a.factor()

	assert.Equal(t, "1234\n", a.output.String())
}

func TestAnaSyn_FactorIdentifier(t *testing.T) {
	id := lexical.NewSymbol(lexical.ID, "foo-bar", 0)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(id, end)

	a.factor()

	assert.Equal(t, "foo-bar\n", a.output.String())
}

func TestAnaSyn_FactorError(t *testing.T) {
	unknown := lexical.NewSymbol('€', "", 0)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(unknown, end)

	assert.Panics(t, func() { a.factor() })
}
