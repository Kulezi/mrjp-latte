package typecheck

import (
	. "latte/compiler/frontend/types"
	"latte/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func (v *visitor) evalType(ctx antlr.ParserRuleContext) (Type, error) {
	t := v.Visit(ctx)
	if err, ok := t.(error); ok {
		return nil, err
	}

	return t.(Type), nil
}

func (v *visitor) VisitTNonVoid(ctx *parser.TNonVoidContext) interface{} {
	return v.Visit(ctx.Nvtype_())
}

func (v *visitor) VisitTSingular(ctx *parser.TSingularContext) interface{} {
	return v.Visit(ctx.Singular_type_())
}

func (v *visitor) VisitTInt(ctx *parser.TIntContext) interface{} {
	return TInt{StartToken: ctx.GetStart()}
}

func (v *visitor) VisitTStr(ctx *parser.TStrContext) interface{} {
	return TString{StartToken: ctx.GetStart()}
}

func (v *visitor) VisitTBool(ctx *parser.TBoolContext) interface{} {
	return TBool{StartToken: ctx.GetStart()}
}

func (v *visitor) VisitTVoid(ctx *parser.TVoidContext) interface{} {
	return TVoid{StartToken: ctx.GetStart()}
}

func (v *visitor) VisitTClass(ctx *parser.TClassContext) interface{} {
	ident := ctx.ID().GetText()
	t, ok := v.TypeOfGlobal(ident)
	if !ok {
		return UnknownClassError{
			Type: TClassRef{
				ID: ctx.ID(),
			},
		}
	}

	class, ok := t.Type.(TClass)
	if !ok {
		return UnknownClassError{
			Type: t,
		}
	}

	return class
}

func (v *visitor) VisitTArray(ctx *parser.TArrayContext) interface{} {
	t, err := v.evalType(ctx.Singular_type_())
	if err != nil {
		return err
	}

	return TArray{
		StartToken: ctx.GetStart(),
		Elem:       t,
	}
}
