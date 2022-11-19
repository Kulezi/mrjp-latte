// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Latte

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseLatteListener is a complete listener for a parse tree produced by LatteParser.
type BaseLatteListener struct{}

var _ LatteListener = &BaseLatteListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLatteListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLatteListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLatteListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLatteListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseLatteListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseLatteListener) ExitProgram(ctx *ProgramContext) {}

// EnterTopDef is called when production topDef is entered.
func (s *BaseLatteListener) EnterTopDef(ctx *TopDefContext) {}

// ExitTopDef is called when production topDef is exited.
func (s *BaseLatteListener) ExitTopDef(ctx *TopDefContext) {}

// EnterFunDef is called when production FunDef is entered.
func (s *BaseLatteListener) EnterFunDef(ctx *FunDefContext) {}

// ExitFunDef is called when production FunDef is exited.
func (s *BaseLatteListener) ExitFunDef(ctx *FunDefContext) {}

// EnterBaseClassDef is called when production BaseClassDef is entered.
func (s *BaseLatteListener) EnterBaseClassDef(ctx *BaseClassDefContext) {}

// ExitBaseClassDef is called when production BaseClassDef is exited.
func (s *BaseLatteListener) ExitBaseClassDef(ctx *BaseClassDefContext) {}

// EnterDerivedClassDef is called when production DerivedClassDef is entered.
func (s *BaseLatteListener) EnterDerivedClassDef(ctx *DerivedClassDefContext) {}

// ExitDerivedClassDef is called when production DerivedClassDef is exited.
func (s *BaseLatteListener) ExitDerivedClassDef(ctx *DerivedClassDefContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseLatteListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseLatteListener) ExitArg(ctx *ArgContext) {}

// EnterClassFieldDef is called when production ClassFieldDef is entered.
func (s *BaseLatteListener) EnterClassFieldDef(ctx *ClassFieldDefContext) {}

// ExitClassFieldDef is called when production ClassFieldDef is exited.
func (s *BaseLatteListener) ExitClassFieldDef(ctx *ClassFieldDefContext) {}

// EnterClassMethodDef is called when production ClassMethodDef is entered.
func (s *BaseLatteListener) EnterClassMethodDef(ctx *ClassMethodDefContext) {}

// ExitClassMethodDef is called when production ClassMethodDef is exited.
func (s *BaseLatteListener) ExitClassMethodDef(ctx *ClassMethodDefContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLatteListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLatteListener) ExitBlock(ctx *BlockContext) {}

// EnterLVArrayRef is called when production LVArrayRef is entered.
func (s *BaseLatteListener) EnterLVArrayRef(ctx *LVArrayRefContext) {}

// ExitLVArrayRef is called when production LVArrayRef is exited.
func (s *BaseLatteListener) ExitLVArrayRef(ctx *LVArrayRefContext) {}

// EnterLVId is called when production LVId is entered.
func (s *BaseLatteListener) EnterLVId(ctx *LVIdContext) {}

// ExitLVId is called when production LVId is exited.
func (s *BaseLatteListener) ExitLVId(ctx *LVIdContext) {}

// EnterLVField is called when production LVField is entered.
func (s *BaseLatteListener) EnterLVField(ctx *LVFieldContext) {}

// ExitLVField is called when production LVField is exited.
func (s *BaseLatteListener) ExitLVField(ctx *LVFieldContext) {}

// EnterSEmpty is called when production SEmpty is entered.
func (s *BaseLatteListener) EnterSEmpty(ctx *SEmptyContext) {}

// ExitSEmpty is called when production SEmpty is exited.
func (s *BaseLatteListener) ExitSEmpty(ctx *SEmptyContext) {}

// EnterSBlockStmt is called when production SBlockStmt is entered.
func (s *BaseLatteListener) EnterSBlockStmt(ctx *SBlockStmtContext) {}

// ExitSBlockStmt is called when production SBlockStmt is exited.
func (s *BaseLatteListener) ExitSBlockStmt(ctx *SBlockStmtContext) {}

// EnterSDecl is called when production SDecl is entered.
func (s *BaseLatteListener) EnterSDecl(ctx *SDeclContext) {}

// ExitSDecl is called when production SDecl is exited.
func (s *BaseLatteListener) ExitSDecl(ctx *SDeclContext) {}

// EnterSAss is called when production SAss is entered.
func (s *BaseLatteListener) EnterSAss(ctx *SAssContext) {}

// ExitSAss is called when production SAss is exited.
func (s *BaseLatteListener) ExitSAss(ctx *SAssContext) {}

// EnterSAssArray is called when production SAssArray is entered.
func (s *BaseLatteListener) EnterSAssArray(ctx *SAssArrayContext) {}

// ExitSAssArray is called when production SAssArray is exited.
func (s *BaseLatteListener) ExitSAssArray(ctx *SAssArrayContext) {}

// EnterSIncr is called when production SIncr is entered.
func (s *BaseLatteListener) EnterSIncr(ctx *SIncrContext) {}

// ExitSIncr is called when production SIncr is exited.
func (s *BaseLatteListener) ExitSIncr(ctx *SIncrContext) {}

// EnterSDecr is called when production SDecr is entered.
func (s *BaseLatteListener) EnterSDecr(ctx *SDecrContext) {}

// ExitSDecr is called when production SDecr is exited.
func (s *BaseLatteListener) ExitSDecr(ctx *SDecrContext) {}

// EnterSRet is called when production SRet is entered.
func (s *BaseLatteListener) EnterSRet(ctx *SRetContext) {}

// ExitSRet is called when production SRet is exited.
func (s *BaseLatteListener) ExitSRet(ctx *SRetContext) {}

// EnterSVRet is called when production SVRet is entered.
func (s *BaseLatteListener) EnterSVRet(ctx *SVRetContext) {}

// ExitSVRet is called when production SVRet is exited.
func (s *BaseLatteListener) ExitSVRet(ctx *SVRetContext) {}

// EnterSCond is called when production SCond is entered.
func (s *BaseLatteListener) EnterSCond(ctx *SCondContext) {}

// ExitSCond is called when production SCond is exited.
func (s *BaseLatteListener) ExitSCond(ctx *SCondContext) {}

// EnterSCondElse is called when production SCondElse is entered.
func (s *BaseLatteListener) EnterSCondElse(ctx *SCondElseContext) {}

// ExitSCondElse is called when production SCondElse is exited.
func (s *BaseLatteListener) ExitSCondElse(ctx *SCondElseContext) {}

// EnterSWhile is called when production SWhile is entered.
func (s *BaseLatteListener) EnterSWhile(ctx *SWhileContext) {}

// ExitSWhile is called when production SWhile is exited.
func (s *BaseLatteListener) ExitSWhile(ctx *SWhileContext) {}

// EnterSFor is called when production SFor is entered.
func (s *BaseLatteListener) EnterSFor(ctx *SForContext) {}

// ExitSFor is called when production SFor is exited.
func (s *BaseLatteListener) ExitSFor(ctx *SForContext) {}

// EnterSExp is called when production SExp is entered.
func (s *BaseLatteListener) EnterSExp(ctx *SExpContext) {}

// ExitSExp is called when production SExp is exited.
func (s *BaseLatteListener) ExitSExp(ctx *SExpContext) {}

// EnterTArray is called when production TArray is entered.
func (s *BaseLatteListener) EnterTArray(ctx *TArrayContext) {}

// ExitTArray is called when production TArray is exited.
func (s *BaseLatteListener) ExitTArray(ctx *TArrayContext) {}

// EnterTSingular is called when production TSingular is entered.
func (s *BaseLatteListener) EnterTSingular(ctx *TSingularContext) {}

// ExitTSingular is called when production TSingular is exited.
func (s *BaseLatteListener) ExitTSingular(ctx *TSingularContext) {}

// EnterTClass is called when production TClass is entered.
func (s *BaseLatteListener) EnterTClass(ctx *TClassContext) {}

// ExitTClass is called when production TClass is exited.
func (s *BaseLatteListener) ExitTClass(ctx *TClassContext) {}

// EnterTInt is called when production TInt is entered.
func (s *BaseLatteListener) EnterTInt(ctx *TIntContext) {}

// ExitTInt is called when production TInt is exited.
func (s *BaseLatteListener) ExitTInt(ctx *TIntContext) {}

// EnterTStr is called when production TStr is entered.
func (s *BaseLatteListener) EnterTStr(ctx *TStrContext) {}

// ExitTStr is called when production TStr is exited.
func (s *BaseLatteListener) ExitTStr(ctx *TStrContext) {}

// EnterTBool is called when production TBool is entered.
func (s *BaseLatteListener) EnterTBool(ctx *TBoolContext) {}

// ExitTBool is called when production TBool is exited.
func (s *BaseLatteListener) ExitTBool(ctx *TBoolContext) {}

// EnterTVoid is called when production TVoid is entered.
func (s *BaseLatteListener) EnterTVoid(ctx *TVoidContext) {}

// ExitTVoid is called when production TVoid is exited.
func (s *BaseLatteListener) ExitTVoid(ctx *TVoidContext) {}

// EnterItem is called when production item is entered.
func (s *BaseLatteListener) EnterItem(ctx *ItemContext) {}

// ExitItem is called when production item is exited.
func (s *BaseLatteListener) ExitItem(ctx *ItemContext) {}

// EnterEId is called when production EId is entered.
func (s *BaseLatteListener) EnterEId(ctx *EIdContext) {}

// ExitEId is called when production EId is exited.
func (s *BaseLatteListener) ExitEId(ctx *EIdContext) {}

// EnterESelf is called when production ESelf is entered.
func (s *BaseLatteListener) EnterESelf(ctx *ESelfContext) {}

// ExitESelf is called when production ESelf is exited.
func (s *BaseLatteListener) ExitESelf(ctx *ESelfContext) {}

// EnterEFunCall is called when production EFunCall is entered.
func (s *BaseLatteListener) EnterEFunCall(ctx *EFunCallContext) {}

// ExitEFunCall is called when production EFunCall is exited.
func (s *BaseLatteListener) ExitEFunCall(ctx *EFunCallContext) {}

// EnterENewArray is called when production ENewArray is entered.
func (s *BaseLatteListener) EnterENewArray(ctx *ENewArrayContext) {}

// ExitENewArray is called when production ENewArray is exited.
func (s *BaseLatteListener) ExitENewArray(ctx *ENewArrayContext) {}

// EnterEArrayRef is called when production EArrayRef is entered.
func (s *BaseLatteListener) EnterEArrayRef(ctx *EArrayRefContext) {}

// ExitEArrayRef is called when production EArrayRef is exited.
func (s *BaseLatteListener) ExitEArrayRef(ctx *EArrayRefContext) {}

// EnterERelOp is called when production ERelOp is entered.
func (s *BaseLatteListener) EnterERelOp(ctx *ERelOpContext) {}

// ExitERelOp is called when production ERelOp is exited.
func (s *BaseLatteListener) ExitERelOp(ctx *ERelOpContext) {}

// EnterETrue is called when production ETrue is entered.
func (s *BaseLatteListener) EnterETrue(ctx *ETrueContext) {}

// ExitETrue is called when production ETrue is exited.
func (s *BaseLatteListener) ExitETrue(ctx *ETrueContext) {}

// EnterEOr is called when production EOr is entered.
func (s *BaseLatteListener) EnterEOr(ctx *EOrContext) {}

// ExitEOr is called when production EOr is exited.
func (s *BaseLatteListener) ExitEOr(ctx *EOrContext) {}

// EnterEInt is called when production EInt is entered.
func (s *BaseLatteListener) EnterEInt(ctx *EIntContext) {}

// ExitEInt is called when production EInt is exited.
func (s *BaseLatteListener) ExitEInt(ctx *EIntContext) {}

// EnterEUnOp is called when production EUnOp is entered.
func (s *BaseLatteListener) EnterEUnOp(ctx *EUnOpContext) {}

// ExitEUnOp is called when production EUnOp is exited.
func (s *BaseLatteListener) ExitEUnOp(ctx *EUnOpContext) {}

// EnterENewStruct is called when production ENewStruct is entered.
func (s *BaseLatteListener) EnterENewStruct(ctx *ENewStructContext) {}

// ExitENewStruct is called when production ENewStruct is exited.
func (s *BaseLatteListener) ExitENewStruct(ctx *ENewStructContext) {}

// EnterEStr is called when production EStr is entered.
func (s *BaseLatteListener) EnterEStr(ctx *EStrContext) {}

// ExitEStr is called when production EStr is exited.
func (s *BaseLatteListener) ExitEStr(ctx *EStrContext) {}

// EnterEMulOp is called when production EMulOp is entered.
func (s *BaseLatteListener) EnterEMulOp(ctx *EMulOpContext) {}

// ExitEMulOp is called when production EMulOp is exited.
func (s *BaseLatteListener) ExitEMulOp(ctx *EMulOpContext) {}

// EnterEAnd is called when production EAnd is entered.
func (s *BaseLatteListener) EnterEAnd(ctx *EAndContext) {}

// ExitEAnd is called when production EAnd is exited.
func (s *BaseLatteListener) ExitEAnd(ctx *EAndContext) {}

// EnterEParen is called when production EParen is entered.
func (s *BaseLatteListener) EnterEParen(ctx *EParenContext) {}

// ExitEParen is called when production EParen is exited.
func (s *BaseLatteListener) ExitEParen(ctx *EParenContext) {}

// EnterEFalse is called when production EFalse is entered.
func (s *BaseLatteListener) EnterEFalse(ctx *EFalseContext) {}

// ExitEFalse is called when production EFalse is exited.
func (s *BaseLatteListener) ExitEFalse(ctx *EFalseContext) {}

// EnterEAddOp is called when production EAddOp is entered.
func (s *BaseLatteListener) EnterEAddOp(ctx *EAddOpContext) {}

// ExitEAddOp is called when production EAddOp is exited.
func (s *BaseLatteListener) ExitEAddOp(ctx *EAddOpContext) {}

// EnterENull is called when production ENull is entered.
func (s *BaseLatteListener) EnterENull(ctx *ENullContext) {}

// ExitENull is called when production ENull is exited.
func (s *BaseLatteListener) ExitENull(ctx *ENullContext) {}

// EnterEFieldAccess is called when production EFieldAccess is entered.
func (s *BaseLatteListener) EnterEFieldAccess(ctx *EFieldAccessContext) {}

// ExitEFieldAccess is called when production EFieldAccess is exited.
func (s *BaseLatteListener) ExitEFieldAccess(ctx *EFieldAccessContext) {}

// EnterAddOp is called when production addOp is entered.
func (s *BaseLatteListener) EnterAddOp(ctx *AddOpContext) {}

// ExitAddOp is called when production addOp is exited.
func (s *BaseLatteListener) ExitAddOp(ctx *AddOpContext) {}

// EnterMulOp is called when production mulOp is entered.
func (s *BaseLatteListener) EnterMulOp(ctx *MulOpContext) {}

// ExitMulOp is called when production mulOp is exited.
func (s *BaseLatteListener) ExitMulOp(ctx *MulOpContext) {}

// EnterRelOp is called when production relOp is entered.
func (s *BaseLatteListener) EnterRelOp(ctx *RelOpContext) {}

// ExitRelOp is called when production relOp is exited.
func (s *BaseLatteListener) ExitRelOp(ctx *RelOpContext) {}
