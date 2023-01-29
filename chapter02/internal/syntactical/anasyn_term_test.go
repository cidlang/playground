package syntactical

import (
	"compiler/internal/lexical"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnaSyn_TermMultiplication(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	multiplication := lexical.NewSymbol('*', "", 0)
	num2 := lexical.NewSymbol(lexical.NUM, "", 9876)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, multiplication, num2, end)

	a.term()

	assert.Equal(t, "1234\n9876\n*\n", a.output.String())
}

func TestAnaSyn_TermDivision(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	multiplication := lexical.NewSymbol('*', "", 0)
	num2 := lexical.NewSymbol(lexical.NUM, "", 9876)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, multiplication, num2, end)

	a.term()

	assert.Equal(t, "1234\n9876\n*\n", a.output.String())
}

func TestAnaSyn_TermIntegerDivision(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	intDivision := lexical.NewSymbol(lexical.DIV, "", 0)
	num2 := lexical.NewSymbol(lexical.NUM, "", 9876)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, intDivision, num2, end)

	a.term()

	assert.Equal(t, "1234\n9876\nDIV\n", a.output.String())
}

func TestAnaSyn_TermRemainder(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	remainder := lexical.NewSymbol(lexical.MOD, "", 0)
	num2 := lexical.NewSymbol(lexical.NUM, "", 9876)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, remainder, num2, end)

	a.term()

	assert.Equal(t, "1234\n9876\nMOD\n", a.output.String())
}

func TestAnaSyn_TermFactor(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, end)

	a.term()

	assert.Equal(t, "1234\n", a.output.String())
}

func TestAnaSyn_TermWhatever(t *testing.T) {
	num1 := lexical.NewSymbol(lexical.NUM, "", 1234)
	whatever := lexical.NewSymbol('€', "", 0)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, whatever, end)

	a.term()

	assert.Equal(t, "1234\n", a.output.String())
}
