package automata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	Cadena string
	Error  error
	Valido bool
}{
	{"caaaaaaabaaaabbbbba", nil, true},
	{"caabbbaaaabaaaabba", nil, true},
	{"caaaccccaaaabbbbba", nil, true},
}

func TestValidaAutomata(t *testing.T) {
	for _, v := range cases {
		valido, err := validaAFD(v.Cadena)
		assert.Equal(t, v.Error, err)
		assert.Equal(t, v.Valido, valido)
	}
}
