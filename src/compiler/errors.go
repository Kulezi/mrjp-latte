package compiler

import "fmt"

type DuplicateIdentifierError struct {
	Ident string
	Pos1  string
	Pos2  string
}

func (e DuplicateIdentifierError) Error() string {
	return fmt.Sprintf(
		"found two identical identifiers\n\t%s\nin the same scope\n\tat %v\n\tand %v",
		e.Ident,
		e.Pos1,
		e.Pos2,
	)
}
