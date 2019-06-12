package ast

import (
	"github.com/bitwitch/ape/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.VAR, Literal: "var"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "cones"},
					Value: "cones",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "pepperoni"},
					Value: "pepperoni",
				},
			},
		},
	}

	if program.String() != "var cones = pepperoni;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
