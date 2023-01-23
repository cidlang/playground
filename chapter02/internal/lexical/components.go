package lexical

const (
	EOF  byte = 01
	NONE int  = -1
)

const (
	NUM Component = iota + 256
	DIV
	MOD
	ID
	END
)

type Component int
