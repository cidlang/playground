package syntactical

import (
	"bytes"
	"fmt"

	"compiler/internal/lexical"
)

type AnaSyn interface {
	Run()
}

func NewAnaSyn(anaLex lexical.AnaLex, output *bytes.Buffer) AnaSyn {
	return &anaSyn{anaLex: anaLex, output: output}
}

type anaSyn struct {
	anaLex      lexical.AnaLex
	preAnalysis lexical.Symbol
	output      *bytes.Buffer
}

func (a *anaSyn) Run() {
	a.preAnalysis = a.anaLex.Next()

	for a.preAnalysis.GetComponent() != lexical.END {
		a.expr()
		a.match(';')
	}
}

func (a *anaSyn) expr() {
	a.term()
	for {
		switch a.preAnalysis.GetComponent() {
		case '+', '-':
			p := a.preAnalysis
			a.match(p.GetComponent())
			a.term()
			a.emit(p)
		default:
			return
		}
	}
}

func (a *anaSyn) term() {
	a.factor()
	for {
		switch a.preAnalysis.GetComponent() {
		case '*', '/', lexical.DIV, lexical.MOD:
			p := a.preAnalysis
			a.match(p.GetComponent())
			a.factor()
			a.emit(p)
		default:
			return
		}
	}
}

func (a *anaSyn) factor() {
	switch a.preAnalysis.GetComponent() {
	case '(':
		a.match('(')
		a.expr()
		a.match(')')
	case lexical.NUM:
		a.emit(a.preAnalysis)
		a.match(lexical.NUM)
	case lexical.ID:
		a.emit(a.preAnalysis)
		a.match(lexical.ID)
	default:
		panic("syntax error")
	}
}

func (a *anaSyn) match(c lexical.Component) {
	if a.preAnalysis.GetComponent() == c {
		a.preAnalysis = a.anaLex.Next()
		return
	}
	panic("syntax error")
}

func (a *anaSyn) emit(s lexical.Symbol) {
	switch s.GetComponent() {
	case '+', '-', '*', '/':
		fmt.Fprintln(a.output, s.GetComponent().ToString())
	case lexical.DIV:
		fmt.Fprintln(a.output, "DIV")
	case lexical.MOD:
		fmt.Fprintln(a.output, "MOD")
	case lexical.NUM:
		fmt.Fprintln(a.output, s.GetValue())
	case lexical.ID:
		fmt.Fprintln(a.output, s.GetLexeme())
	default:
		fmt.Fprintf(a.output, "component %v, value %v\n", s.GetComponent().ToString(), s.GetValue())
	}
}
