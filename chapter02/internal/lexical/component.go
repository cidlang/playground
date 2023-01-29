package lexical

const (
	NUM Component = iota + 256
	DIV
	MOD
	ID
	END
)

type Component rune

func (c Component) ToString() string {
	return string(c)
}
