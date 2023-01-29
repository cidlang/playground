package lexical

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolTableInsert(t *testing.T) {
	st := make(symbolTable, 0)
	expected := NewSymbol(NUM, "foo-bar", 0)
	actual := st.insert("foo-bar", NUM)

	assert.Equal(t, expected, actual)
}

func TestSymbolTableSearch(t *testing.T) {
	st := make(symbolTable, 0)
	expected := NewSymbol(ID, "bar-foo", 0)
	st.insert("bar-foo", ID)
	actual, found := st.search("bar-foo")

	assert.Equal(t, expected, actual)
	assert.True(t, found)
}

func TestSymbolTableSearchNotFound(t *testing.T) {
	st := make(symbolTable, 0)
	expected := Symbol(nil)

	st.insert("foo-bar", NUM)
	actual, found := st.search("bar-foo")

	assert.Equal(t, expected, actual)
	assert.False(t, found)
}

func TestNewSymbolTable(t *testing.T) {
	sFoo := NewSymbol(NUM, "foo-bar", 0)
	sBar := NewSymbol(ID, "bar-foo", 0)

	ss := []Symbol{sFoo, sBar}
	st := newSymbolTable(ss)

	actFoo, foundFoo := st.search("foo-bar")
	actBar, foundBar := st.search("bar-foo")

	assert.Equal(t, sFoo, actFoo)
	assert.True(t, foundFoo)
	assert.Equal(t, sBar, actBar)
	assert.True(t, foundBar)
}
