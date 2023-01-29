package lexical

type Symbol interface {
	GetComponent() Component
	GetLexeme() string
	GetValue() int
}

type symbol struct {
	component Component
	lexeme    string
	value     int
}

func NewSymbol(component Component, lexeme string, value int) Symbol {
	return symbol{
		component: component,
		lexeme:    lexeme,
		value:     value,
	}
}

func (s symbol) GetComponent() Component {
	return s.component
}

func (s symbol) GetLexeme() string {
	return s.lexeme
}

func (s symbol) GetValue() int {
	return s.value
}
