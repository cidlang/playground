package syntactical

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"compiler/internal/lexical"
)

func TestAnaSyn_Match(t *testing.T) {
	current := lexical.NewSymbol('a', "b", 'c')
	next := lexical.NewSymbol('1', "2", '3')
	a := newAnaSynTest(current, next)

	a.match('a')

	assert.Equal(t, next.GetComponent(), a.preAnalysis.GetComponent())
	assert.Equal(t, next.GetLexeme(), a.preAnalysis.GetLexeme())
	assert.Equal(t, next.GetValue(), a.preAnalysis.GetValue())
}

func TestAnaSyn_MatchFails(t *testing.T) {
	current := lexical.NewSymbol('a', "b", 'c')
	next := lexical.NewSymbol('1', "2", '3')
	a := newAnaSynTest(current, next)

	assert.Panics(t, func() { a.match('1') })
}
