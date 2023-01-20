// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Latte

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by LatteParser.
type LatteVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LatteParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by LatteParser#topDef.
	VisitTopDef(ctx *TopDefContext) interface{}

	// Visit a parse tree produced by LatteParser#FunDef.
	VisitFunDef(ctx *FunDefContext) interface{}

	// Visit a parse tree produced by LatteParser#BaseClassDef.
	VisitBaseClassDef(ctx *BaseClassDefContext) interface{}

	// Visit a parse tree produced by LatteParser#DerivedClassDef.
	VisitDerivedClassDef(ctx *DerivedClassDefContext) interface{}

	// Visit a parse tree produced by LatteParser#arg.
	VisitArg(ctx *ArgContext) interface{}

	// Visit a parse tree produced by LatteParser#ClassFieldDef.
	VisitClassFieldDef(ctx *ClassFieldDefContext) interface{}

	// Visit a parse tree produced by LatteParser#ClassMethodDef.
	VisitClassMethodDef(ctx *ClassMethodDefContext) interface{}

	// Visit a parse tree produced by LatteParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LatteParser#LVArrayRef.
	VisitLVArrayRef(ctx *LVArrayRefContext) interface{}

	// Visit a parse tree produced by LatteParser#LVFieldArrayRef.
	VisitLVFieldArrayRef(ctx *LVFieldArrayRefContext) interface{}

	// Visit a parse tree produced by LatteParser#LVFieldMethodCall.
	VisitLVFieldMethodCall(ctx *LVFieldMethodCallContext) interface{}

	// Visit a parse tree produced by LatteParser#LVField.
	VisitLVField(ctx *LVFieldContext) interface{}

	// Visit a parse tree produced by LatteParser#LVId.
	VisitLVId(ctx *LVIdContext) interface{}

	// Visit a parse tree produced by LatteParser#SEmpty.
	VisitSEmpty(ctx *SEmptyContext) interface{}

	// Visit a parse tree produced by LatteParser#SBlockStmt.
	VisitSBlockStmt(ctx *SBlockStmtContext) interface{}

	// Visit a parse tree produced by LatteParser#SDecl.
	VisitSDecl(ctx *SDeclContext) interface{}

	// Visit a parse tree produced by LatteParser#SAss.
	VisitSAss(ctx *SAssContext) interface{}

	// Visit a parse tree produced by LatteParser#SIncr.
	VisitSIncr(ctx *SIncrContext) interface{}

	// Visit a parse tree produced by LatteParser#SDecr.
	VisitSDecr(ctx *SDecrContext) interface{}

	// Visit a parse tree produced by LatteParser#SRet.
	VisitSRet(ctx *SRetContext) interface{}

	// Visit a parse tree produced by LatteParser#SVRet.
	VisitSVRet(ctx *SVRetContext) interface{}

	// Visit a parse tree produced by LatteParser#SCond.
	VisitSCond(ctx *SCondContext) interface{}

	// Visit a parse tree produced by LatteParser#SCondElse.
	VisitSCondElse(ctx *SCondElseContext) interface{}

	// Visit a parse tree produced by LatteParser#SWhile.
	VisitSWhile(ctx *SWhileContext) interface{}

	// Visit a parse tree produced by LatteParser#SFor.
	VisitSFor(ctx *SForContext) interface{}

	// Visit a parse tree produced by LatteParser#SExp.
	VisitSExp(ctx *SExpContext) interface{}

	// Visit a parse tree produced by LatteParser#TNonVoid.
	VisitTNonVoid(ctx *TNonVoidContext) interface{}

	// Visit a parse tree produced by LatteParser#TVoid.
	VisitTVoid(ctx *TVoidContext) interface{}

	// Visit a parse tree produced by LatteParser#TArray.
	VisitTArray(ctx *TArrayContext) interface{}

	// Visit a parse tree produced by LatteParser#TSingular.
	VisitTSingular(ctx *TSingularContext) interface{}

	// Visit a parse tree produced by LatteParser#TClass.
	VisitTClass(ctx *TClassContext) interface{}

	// Visit a parse tree produced by LatteParser#TInt.
	VisitTInt(ctx *TIntContext) interface{}

	// Visit a parse tree produced by LatteParser#TStr.
	VisitTStr(ctx *TStrContext) interface{}

	// Visit a parse tree produced by LatteParser#TBool.
	VisitTBool(ctx *TBoolContext) interface{}

	// Visit a parse tree produced by LatteParser#item.
	VisitItem(ctx *ItemContext) interface{}

	// Visit a parse tree produced by LatteParser#EId.
	VisitEId(ctx *EIdContext) interface{}

	// Visit a parse tree produced by LatteParser#ESelf.
	VisitESelf(ctx *ESelfContext) interface{}

	// Visit a parse tree produced by LatteParser#EFunCall.
	VisitEFunCall(ctx *EFunCallContext) interface{}

	// Visit a parse tree produced by LatteParser#ENewArray.
	VisitENewArray(ctx *ENewArrayContext) interface{}

	// Visit a parse tree produced by LatteParser#EArrayRef.
	VisitEArrayRef(ctx *EArrayRefContext) interface{}

	// Visit a parse tree produced by LatteParser#ERelOp.
	VisitERelOp(ctx *ERelOpContext) interface{}

	// Visit a parse tree produced by LatteParser#ETrue.
	VisitETrue(ctx *ETrueContext) interface{}

	// Visit a parse tree produced by LatteParser#EOr.
	VisitEOr(ctx *EOrContext) interface{}

	// Visit a parse tree produced by LatteParser#EInt.
	VisitEInt(ctx *EIntContext) interface{}

	// Visit a parse tree produced by LatteParser#EStr.
	VisitEStr(ctx *EStrContext) interface{}

	// Visit a parse tree produced by LatteParser#EFieldArrayAccess.
	VisitEFieldArrayAccess(ctx *EFieldArrayAccessContext) interface{}

	// Visit a parse tree produced by LatteParser#ENotOp.
	VisitENotOp(ctx *ENotOpContext) interface{}

	// Visit a parse tree produced by LatteParser#EMulOp.
	VisitEMulOp(ctx *EMulOpContext) interface{}

	// Visit a parse tree produced by LatteParser#EAnd.
	VisitEAnd(ctx *EAndContext) interface{}

	// Visit a parse tree produced by LatteParser#EMethodCall.
	VisitEMethodCall(ctx *EMethodCallContext) interface{}

	// Visit a parse tree produced by LatteParser#ENegOp.
	VisitENegOp(ctx *ENegOpContext) interface{}

	// Visit a parse tree produced by LatteParser#EParen.
	VisitEParen(ctx *EParenContext) interface{}

	// Visit a parse tree produced by LatteParser#EFalse.
	VisitEFalse(ctx *EFalseContext) interface{}

	// Visit a parse tree produced by LatteParser#ENew.
	VisitENew(ctx *ENewContext) interface{}

	// Visit a parse tree produced by LatteParser#EAddOp.
	VisitEAddOp(ctx *EAddOpContext) interface{}

	// Visit a parse tree produced by LatteParser#ENull.
	VisitENull(ctx *ENullContext) interface{}

	// Visit a parse tree produced by LatteParser#EFieldAccess.
	VisitEFieldAccess(ctx *EFieldAccessContext) interface{}

	// Visit a parse tree produced by LatteParser#addOp.
	VisitAddOp(ctx *AddOpContext) interface{}

	// Visit a parse tree produced by LatteParser#mulOp.
	VisitMulOp(ctx *MulOpContext) interface{}

	// Visit a parse tree produced by LatteParser#relOp.
	VisitRelOp(ctx *RelOpContext) interface{}
}
