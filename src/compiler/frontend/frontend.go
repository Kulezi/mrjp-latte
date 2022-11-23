package frontend

import (
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func Run(filename string) error {
	s := state{}
	if err := s.parse(filename); err != nil {
		return err
	}

	if err := s.semanticCheck(); err != nil {
		return err
	}

	return nil
}

func (s *state) parse(filename string) error {
	input, err := antlr.NewFileStream(filename)
	if err != nil {
		return err
	}

	lexErrors := &CustomErrorListener{}
	lexer := parser.NewLatteLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexErrors)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	parseErrors := &CustomErrorListener{}
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
	// Evaluate method/function signatures and inheritance tree.
	visitor := MakeGlobalDeclVisitor()
	if err := visitor.Run(s.tree); err != nil {
		return err
	}

	s.signatures.Globals = visitor.signatures.Globals
	s.signatures.Parent = visitor.signatures.Parent

	if err := s.signatures.Inherit(); err != nil {
		return err
	}

	// fmt.Fprintln(os.Stderr, (&TypeCheckVisitor{}).Visit(c.tree))
	return nil
}
