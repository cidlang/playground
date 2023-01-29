package syntactical

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"compiler/internal/lexical"
)

func TestAnaSyn_ExprAddition(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	addition := lexical.NewSymbol('+', "", 0)
	num2 := lexical.NewSymbol(lexical.NUM, "", 9876)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, addition, num2, end)

	a.expr()

	assert.Equal(t, "1234\n9876\n+\n", a.output.String())
}

func TestAnaSyn_ExprSubtraction(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	subtraction := lexical.NewSymbol('-', "", 0)
	num2 := lexical.NewSymbol(lexical.NUM, "", 9876)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, subtraction, num2, end)

	a.expr()

	assert.Equal(t, "1234\n9876\n-\n", a.output.String())
}

func TestAnaSyn_ExprTerm(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, end)

	a.expr()

	assert.Equal(t, "1234\n", a.output.String())
}

func TestAnaSyn_ExprWhatever(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	whatever := lexical.NewSymbol('€', "", 0)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, whatever, end)

	a.expr()

	assert.Equal(t, "1234\n", a.output.String())
}
