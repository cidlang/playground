package lexical

type symbolTable map[string]Symbol

func (st symbolTable) insert(lexeme string, component Component) Symbol {
	symbol := Symbol{
		lexeme:    lexeme,
		value:     0,
		component: component,
	}

	st[lexeme] = symbol
	return symbol
}

func (st symbolTable) search(lexeme string) (Symbol, bool) {
	symbol, found := st[lexeme]
	return symbol, found
}

func newSymbolTable(ss []Symbol) symbolTable {
	var st symbolTable
	st = make(symbolTable, 0)

	for _, s := range ss {
		st.insert(s.GetLexeme(), s.GetComponent())
	}

	return st
}
