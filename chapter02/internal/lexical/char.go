package lexical

type char byte

func (c char) isBlank() bool {
	return c == ' ' || c == '\t'
}

func (c char) isNewLine() bool {
	return c == '\n'
}

func (c char) isDigit() bool {
	return c >= '0' && c <= '9'
}

func (c char) isAlpha() bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func (c char) isAlNum() bool {
	return c.isDigit() || c.isAlpha()
}

func (c char) isEOF() bool {
	return c == 0
}
