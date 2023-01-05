package frontend

import (
	"fmt"
	"latte/compiler/frontend/typecheck"
	. "latte/compiler/frontend/types"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type State struct {
	Tree       parser.IProgramContext
	Signatures Signatures
}

func Run(filename string) (State, error) {
	s := State{}
	if err := s.parse(filename); err != nil {
		return State{}, err
	}

	if err := s.semanticCheck(); err != nil {
		return State{}, err
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
