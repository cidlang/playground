package syntactical

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"compiler/internal/lexical"
)

func TestAnaSyn_Run(t *testing.T) {
	// 1 + 2 * 3 div foo - bar
	num1 := lexical.NewSymbol(lexical.NUM, "", 1)
	addition := lexical.NewSymbol('+', "", 0)
	num2 := lexical.NewSymbol(lexical.NUM, "", 2)
	multiplication := lexical.NewSymbol('*', "", 0)
	num3 := lexical.NewSymbol(lexical.NUM, "", 3)
	intDivision := lexical.NewSymbol(lexical.DIV, "", 0)
	id1 := lexical.NewSymbol(lexical.ID, "foo", 0)
	subtraction := lexical.NewSymbol('-', "", 0)
	id2 := lexical.NewSymbol(lexical.ID, "bar", 0)
	end := lexical.NewSymbol(lexical.END, "", 0)
	a := newAnaSynTest(num1, addition, num2, multiplication, num3, intDivision, id1, subtraction, id2, end)

	a.expr()

	expected := `1
2
3
*
foo
DIV
+
bar
-
`

	assert.Equal(t, expected, a.output.String())
}
