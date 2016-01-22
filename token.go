package easylex

import (
	"fmt"
)

type TokenType int

const (
	TokenEOF   TokenType = -1
	TokenError TokenType = -2
)

type Token struct {
	Typ TokenType
	Val string
}

func (t token) String() string {
	switch t.Typ {
	case TokenError:
		return t.Val
	case tokenEOF:
		return "EOF"
	}
	if len(t.Val) > 10 {
		return fmt.Sprintf("%.10q...", t.Val)
	}
	return fmt.Sprintf("%q", t.Val)
}