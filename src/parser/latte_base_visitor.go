package parser // Latte

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type BaseLatteVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLatteVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTopDef(ctx *TopDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitFunDef(ctx *FunDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitBaseClassDef(ctx *BaseClassDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitDerivedClassDef(ctx *DerivedClassDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitArg(ctx *ArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitClassFieldDef(ctx *ClassFieldDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitClassMethodDef(ctx *ClassMethodDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitLVArrayRef(ctx *LVArrayRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitLVId(ctx *LVIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitLVField(ctx *LVFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSEmpty(ctx *SEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSBlockStmt(ctx *SBlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSDecl(ctx *SDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSAss(ctx *SAssContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSIncr(ctx *SIncrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSDecr(ctx *SDecrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSRet(ctx *SRetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSVRet(ctx *SVRetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSCond(ctx *SCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSCondElse(ctx *SCondElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSWhile(ctx *SWhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSFor(ctx *SForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitSExp(ctx *SExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTSingular(ctx *TSingularContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTArray(ctx *TArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTClass(ctx *TClassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTInt(ctx *TIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTStr(ctx *TStrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTBool(ctx *TBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitTVoid(ctx *TVoidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitItem(ctx *ItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEId(ctx *EIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitESelf(ctx *ESelfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEFunCall(ctx *EFunCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEArrayRef(ctx *EArrayRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitERelOp(ctx *ERelOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitETrue(ctx *ETrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEOr(ctx *EOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEInt(ctx *EIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEUnOp(ctx *EUnOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEStr(ctx *EStrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEMulOp(ctx *EMulOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEAnd(ctx *EAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEParen(ctx *EParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEFalse(ctx *EFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitENew(ctx *ENewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEAddOp(ctx *EAddOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitENull(ctx *ENullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitEFieldAccess(ctx *EFieldAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitAddOp(ctx *AddOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitMulOp(ctx *MulOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLatteVisitor) VisitRelOp(ctx *RelOpContext) interface{} {
	return v.VisitChildren(ctx)
}

var _ LatteVisitor = &BaseLatteVisitor{}
