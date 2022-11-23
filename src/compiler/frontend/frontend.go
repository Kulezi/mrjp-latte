package frontend

import (
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func Run(filename string) (Signatures, error) {
	s := state{}
	if err := s.parse(filename); err != nil {
		return Signatures{}, err
	}

	if err := s.semanticCheck(); err != nil {
		return Signatures{}, err
	}

	return s.signatures, nil
}

func (s *state) parse(filename string) error {
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
	s.tree = p.Program()

	if err := lexErrors.Check("lexer error:"); err != nil {
		return err
	}

	if err := parseErrors.Check("parse error:"); err != nil {
		return err
	}

	return nil
}

func (s *state) semanticCheck() error {
	var err error
	// Evaluate method/function signatures and inheritance tree.

	if err = s.evalGlobalSignatures(); err != nil {
		return err
	}

	if err = s.signatures.inheritClasses(); err != nil {
		return err
	}

	if err = s.typeCheck(); err != nil {
		return err
	}

	return nil
}

func (s *state) evalGlobalSignatures() error {
	var err error
	s.signatures, err = makeGlobalDeclVisitor().Run(s.tree)
	return err
}

func (s *state) typeCheck() error {
	if err, ok := makeTypeCheckVisitor(s).Visit(s.tree).(error); ok {
		return err
	}

	return nil
}
