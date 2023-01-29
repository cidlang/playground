package lexical

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComponent_ToString(t *testing.T) {
	var c Component
	c = '+'

	s := c.ToString()

	assert.Equal(t, "+", s)
}
