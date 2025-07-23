package lexer

import (
	"fmt"
)

type SNot struct {
	Head Token
	Tail []SNot
}

func (s *SNot) String() {
	switch {
	case s.Head.Type == Atom:
		fmt.Print(s.Head.Value)
	case (s.Head.Type != Atom) && len(s.Tail) > 0:
		fmt.Print("(", s.Head.Value)
		for _, item := range s.Tail {
			fmt.Println(item.Head.Value)
		}
		fmt.Print(")\n")
	}
}
