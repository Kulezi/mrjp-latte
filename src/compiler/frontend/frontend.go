package frontend

import (
	"fmt"
	"latte/compiler/frontend/typecheck"
	. "latte/compiler/frontend/types"
	"latte/parser"
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type State struct {
	Tree       parser.IProgramContext
	Signatures Signatures
}

func printClass(class TClass) {
	s := fmt.Sprintf("class %s {\n", class.ID)
	s += fmt.Sprintf("\tvtable of size %d:\n", class.TotalMethods)
	vtable := make([]Type, class.TotalMethods)
	layout := make([]Type, class.TotalNonMethods+1)
	for _, v := range class.Fields {
		if _, ok := v.Type.(TFun); ok {
			vtable[v.Offset] = v.Type
		} else {
			layout[v.Offset] = v.Type
		}
	}

	for i, v := range vtable {
		s += fmt.Sprintf("\t\ti = %d: %s\n", i, v)
	}

	s += fmt.Sprintf("\tstruct layout of size %d:\n", class.TotalNonMethods+1)
	for i, v := range layout {
		if i == 0 {
			s += "\t\ti = 0: vtable ptr\n"
		} else {
			s += fmt.Sprintf("i = %d: %s\n", i, v)
		}
	}

	s += "}\n"
	log.Println(s)
}

func Run(filename string) (State, error) {
	s := State{}
	if err := s.parse(filename); err != nil {
		return State{}, err
	}

	if err := s.semanticCheck(); err != nil {
		return State{}, err
	}

	for _, v := range s.Signatures.Globals {
		if class, ok := v.Type.(TClass); ok {
			printClass(class)
		}
	}

	return s, nil
}

func (s *State) parse(filename string) error {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return err
	}

	lexErrors := &customErrorListener{}
	lexer := parser.NewLatteLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErrors)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	parseErrors := &customErrorListener{}
	p := parser.NewLatteParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(parseErrors)
	p.BuildParseTrees = true
	s.Tree = p.Program()

	if err := lexErrors.Check("lexer error:"); err != nil {
		return err
	}

	if err := parseErrors.Check("parse error:"); err != nil {
		return err
	}

	return nil
}

func (s *State) semanticCheck() error {
	var err error
	// Evaluate method/function signatures and inheritance tree.

	if err = s.evalGlobalSignatures(); err != nil {
		return err
	}

	if main, ok := s.Signatures.Globals["main"]; !ok {
		return fmt.Errorf("missing main function")
	} else if !SameType(main.Type, TFun{Ident: "main", Result: TInt{}}) {
		return MainInvalidSignatureError{Type: main.Type}
	}

	if err = s.Signatures.InheritClasses(); err != nil {
		return err
	}

	// for _, v := range s.signatures.Globals {
	// 	if class, ok := v.Type.(TClass); ok {
	// 		class.Print()
	// 		continue
	// 	}
	// 	fmt.Println(v.Type)
	// }

	if err = s.typeCheck(); err != nil {
		return err
	}

	return nil
}

func (s *State) evalGlobalSignatures() error {
	var err error
	s.Signatures, err = makeGlobalDeclVisitor().Run(s.Tree)
	return err
}

func (s *State) typeCheck() error {
	if err, ok := typecheck.MakeVisitor(s.Signatures).Visit(s.Tree).(error); ok {
		return err
	}

	return nil
}
