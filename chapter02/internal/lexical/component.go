package lexical

const (
	NUM Component = iota + 256
	DIV
	MOD
	ID
	END
)

type Component int
