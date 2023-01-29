package lexical

type symbolTable map[string]Symbol

func (st symbolTable) insert(lexeme string, component Component) Symbol {
	s := NewSymbol(component, lexeme, 0)

	st[lexeme] = s
	return s
}

func (st symbolTable) search(lexeme string) (Symbol, bool) {
	s, found := st[lexeme]
	return s, found
}

func newSymbolTable(ss []Symbol) symbolTable {
	var st symbolTable
	st = make(symbolTable, 0)

	for _, s := range ss {
		st.insert(s.GetLexeme(), s.GetComponent())
	}

	return st
}
