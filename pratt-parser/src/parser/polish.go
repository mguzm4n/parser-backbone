package parser

import (
	"fmt"
	"strings"
)

type Node interface {
	isNode()
	String() string
}

type Atom struct {
	Char rune
}

func (Atom) isNode() {}
func (a Atom) String() string {
	return fmt.Sprintf("%q", a.Char)
}

type Cons struct {
	Head rune
	Tail []Node
}

func (Cons) isNode() {}

func (c Cons) String() string {
	if len(c.Tail) == 0 {
		return fmt.Sprintf("%q", c.Head)
	}

	var s strings.Builder
	s.WriteRune('(')
	s.WriteRune(c.Head)
	for _, it := range c.Tail {
		s.WriteRune(' ')
		s.WriteString(it.String())
	}
	s.WriteRune(')')
	return s.String()
}
