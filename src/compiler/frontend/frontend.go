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
	// Evaluate method/function signatures and inheritance tree.
	visitor := makeGlobalDeclVisitor()
	if err := visitor.Run(s.tree); err != nil {
		return err
	}

	s.signatures.Globals = visitor.signatures.Globals
	s.signatures.Parent = visitor.signatures.Parent

	if err := s.signatures.inheritClasses(); err != nil {
		return err
	}

	// fmt.Fprintln(os.Stderr, (&TypeCheckVisitor{}).Visit(c.tree))
	return nil
}
