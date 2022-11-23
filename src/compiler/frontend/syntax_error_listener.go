package frontend

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type SyntaxError struct {
	line, column int
	msg          string
}

func (s *SyntaxError) Error() string {
	return fmt.Sprintf("%s, at line: %d, column: %d", s.msg, s.line, s.column)
}

type customErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func (c *customErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	c.Errors = append(c.Errors, &SyntaxError{
		line:   line,
		column: column,

		msg: msg,
	})
}

func (c *customErrorListener) Check(errMsg string) error {
	if len(c.Errors) > 0 {
		return fmt.Errorf("%s\n%v", errMsg, c.Errors)
	}

	return nil
}
