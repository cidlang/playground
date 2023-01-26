package lexical

type Symbol struct {
	component Component
	lexeme    string
	value     int
}

func (s Symbol) GetComponent() Component {
	return s.component
}

func (s Symbol) GetLexeme() string {
	return s.lexeme
}

func (s Symbol) GetValue() int {
	return s.value
}
