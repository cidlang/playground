package syntactical

import (
	"bytes"
	"compiler/internal/lexical"
)

type anaLexMock struct {
	symbols []lexical.Symbol
	pos     int
}

func (m *anaLexMock) Next() lexical.Symbol {
	result := m.symbols[m.pos]
	m.pos++

	return result
}

func newAnaSynTest(current lexical.Symbol, next ...lexical.Symbol) *anaSyn {
	var o bytes.Buffer
	anaLex := anaLexMock{symbols: next}
	return &anaSyn{anaLex: &anaLex, preAnalysis: current, output: &o}
}
