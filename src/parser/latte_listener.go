// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Latte

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// LatteListener is a complete listener for a parse tree produced by LatteParser.
type LatteListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTopDef is called when entering the topDef production.
	EnterTopDef(c *TopDefContext)

	// EnterFunDef is called when entering the FunDef production.
	EnterFunDef(c *FunDefContext)

	// EnterBaseClassDef is called when entering the BaseClassDef production.
	EnterBaseClassDef(c *BaseClassDefContext)

	// EnterDerivedClassDef is called when entering the DerivedClassDef production.
	EnterDerivedClassDef(c *DerivedClassDefContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterClassFieldDef is called when entering the ClassFieldDef production.
	EnterClassFieldDef(c *ClassFieldDefContext)

	// EnterClassMethodDef is called when entering the ClassMethodDef production.
	EnterClassMethodDef(c *ClassMethodDefContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterLVArrayRef is called when entering the LVArrayRef production.
	EnterLVArrayRef(c *LVArrayRefContext)

	// EnterLVId is called when entering the LVId production.
	EnterLVId(c *LVIdContext)

	// EnterLVField is called when entering the LVField production.
	EnterLVField(c *LVFieldContext)

	// EnterSEmpty is called when entering the SEmpty production.
	EnterSEmpty(c *SEmptyContext)

	// EnterSBlockStmt is called when entering the SBlockStmt production.
	EnterSBlockStmt(c *SBlockStmtContext)

	// EnterSDecl is called when entering the SDecl production.
	EnterSDecl(c *SDeclContext)

	// EnterSAss is called when entering the SAss production.
	EnterSAss(c *SAssContext)

	// EnterSIncr is called when entering the SIncr production.
	EnterSIncr(c *SIncrContext)

	// EnterSDecr is called when entering the SDecr production.
	EnterSDecr(c *SDecrContext)

	// EnterSRet is called when entering the SRet production.
	EnterSRet(c *SRetContext)

	// EnterSVRet is called when entering the SVRet production.
	EnterSVRet(c *SVRetContext)

	// EnterSCond is called when entering the SCond production.
	EnterSCond(c *SCondContext)

	// EnterSCondElse is called when entering the SCondElse production.
	EnterSCondElse(c *SCondElseContext)

	// EnterSWhile is called when entering the SWhile production.
	EnterSWhile(c *SWhileContext)

	// EnterSFor is called when entering the SFor production.
	EnterSFor(c *SForContext)

	// EnterSExp is called when entering the SExp production.
	EnterSExp(c *SExpContext)

	// EnterTSingular is called when entering the TSingular production.
	EnterTSingular(c *TSingularContext)

	// EnterTArray is called when entering the TArray production.
	EnterTArray(c *TArrayContext)

	// EnterTClass is called when entering the TClass production.
	EnterTClass(c *TClassContext)

	// EnterTInt is called when entering the TInt production.
	EnterTInt(c *TIntContext)

	// EnterTStr is called when entering the TStr production.
	EnterTStr(c *TStrContext)

	// EnterTBool is called when entering the TBool production.
	EnterTBool(c *TBoolContext)

	// EnterTVoid is called when entering the TVoid production.
	EnterTVoid(c *TVoidContext)

	// EnterItem is called when entering the item production.
	EnterItem(c *ItemContext)

	// EnterEId is called when entering the EId production.
	EnterEId(c *EIdContext)

	// EnterESelf is called when entering the ESelf production.
	EnterESelf(c *ESelfContext)

	// EnterEFunCall is called when entering the EFunCall production.
	EnterEFunCall(c *EFunCallContext)

	// EnterEArrayRef is called when entering the EArrayRef production.
	EnterEArrayRef(c *EArrayRefContext)

	// EnterERelOp is called when entering the ERelOp production.
	EnterERelOp(c *ERelOpContext)

	// EnterETrue is called when entering the ETrue production.
	EnterETrue(c *ETrueContext)

	// EnterEOr is called when entering the EOr production.
	EnterEOr(c *EOrContext)

	// EnterEInt is called when entering the EInt production.
	EnterEInt(c *EIntContext)

	// EnterEUnOp is called when entering the EUnOp production.
	EnterEUnOp(c *EUnOpContext)

	// EnterEStr is called when entering the EStr production.
	EnterEStr(c *EStrContext)

	// EnterEMulOp is called when entering the EMulOp production.
	EnterEMulOp(c *EMulOpContext)

	// EnterEAnd is called when entering the EAnd production.
	EnterEAnd(c *EAndContext)

	// EnterEParen is called when entering the EParen production.
	EnterEParen(c *EParenContext)

	// EnterEFalse is called when entering the EFalse production.
	EnterEFalse(c *EFalseContext)

	// EnterENew is called when entering the ENew production.
	EnterENew(c *ENewContext)

	// EnterEAddOp is called when entering the EAddOp production.
	EnterEAddOp(c *EAddOpContext)

	// EnterENull is called when entering the ENull production.
	EnterENull(c *ENullContext)

	// EnterEFieldAccess is called when entering the EFieldAccess production.
	EnterEFieldAccess(c *EFieldAccessContext)

	// EnterAddOp is called when entering the addOp production.
	EnterAddOp(c *AddOpContext)

	// EnterMulOp is called when entering the mulOp production.
	EnterMulOp(c *MulOpContext)

	// EnterRelOp is called when entering the relOp production.
	EnterRelOp(c *RelOpContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTopDef is called when exiting the topDef production.
	ExitTopDef(c *TopDefContext)

	// ExitFunDef is called when exiting the FunDef production.
	ExitFunDef(c *FunDefContext)

	// ExitBaseClassDef is called when exiting the BaseClassDef production.
	ExitBaseClassDef(c *BaseClassDefContext)

	// ExitDerivedClassDef is called when exiting the DerivedClassDef production.
	ExitDerivedClassDef(c *DerivedClassDefContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitClassFieldDef is called when exiting the ClassFieldDef production.
	ExitClassFieldDef(c *ClassFieldDefContext)

	// ExitClassMethodDef is called when exiting the ClassMethodDef production.
	ExitClassMethodDef(c *ClassMethodDefContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitLVArrayRef is called when exiting the LVArrayRef production.
	ExitLVArrayRef(c *LVArrayRefContext)

	// ExitLVId is called when exiting the LVId production.
	ExitLVId(c *LVIdContext)

	// ExitLVField is called when exiting the LVField production.
	ExitLVField(c *LVFieldContext)

	// ExitSEmpty is called when exiting the SEmpty production.
	ExitSEmpty(c *SEmptyContext)

	// ExitSBlockStmt is called when exiting the SBlockStmt production.
	ExitSBlockStmt(c *SBlockStmtContext)

	// ExitSDecl is called when exiting the SDecl production.
	ExitSDecl(c *SDeclContext)

	// ExitSAss is called when exiting the SAss production.
	ExitSAss(c *SAssContext)

	// ExitSIncr is called when exiting the SIncr production.
	ExitSIncr(c *SIncrContext)

	// ExitSDecr is called when exiting the SDecr production.
	ExitSDecr(c *SDecrContext)

	// ExitSRet is called when exiting the SRet production.
	ExitSRet(c *SRetContext)

	// ExitSVRet is called when exiting the SVRet production.
	ExitSVRet(c *SVRetContext)

	// ExitSCond is called when exiting the SCond production.
	ExitSCond(c *SCondContext)

	// ExitSCondElse is called when exiting the SCondElse production.
	ExitSCondElse(c *SCondElseContext)

	// ExitSWhile is called when exiting the SWhile production.
	ExitSWhile(c *SWhileContext)

	// ExitSFor is called when exiting the SFor production.
	ExitSFor(c *SForContext)

	// ExitSExp is called when exiting the SExp production.
	ExitSExp(c *SExpContext)

	// ExitTSingular is called when exiting the TSingular production.
	ExitTSingular(c *TSingularContext)

	// ExitTArray is called when exiting the TArray production.
	ExitTArray(c *TArrayContext)

	// ExitTClass is called when exiting the TClass production.
	ExitTClass(c *TClassContext)

	// ExitTInt is called when exiting the TInt production.
	ExitTInt(c *TIntContext)

	// ExitTStr is called when exiting the TStr production.
	ExitTStr(c *TStrContext)

	// ExitTBool is called when exiting the TBool production.
	ExitTBool(c *TBoolContext)

	// ExitTVoid is called when exiting the TVoid production.
	ExitTVoid(c *TVoidContext)

	// ExitItem is called when exiting the item production.
	ExitItem(c *ItemContext)

	// ExitEId is called when exiting the EId production.
	ExitEId(c *EIdContext)

	// ExitESelf is called when exiting the ESelf production.
	ExitESelf(c *ESelfContext)

	// ExitEFunCall is called when exiting the EFunCall production.
	ExitEFunCall(c *EFunCallContext)

	// ExitEArrayRef is called when exiting the EArrayRef production.
	ExitEArrayRef(c *EArrayRefContext)

	// ExitERelOp is called when exiting the ERelOp production.
	ExitERelOp(c *ERelOpContext)

	// ExitETrue is called when exiting the ETrue production.
	ExitETrue(c *ETrueContext)

	// ExitEOr is called when exiting the EOr production.
	ExitEOr(c *EOrContext)

	// ExitEInt is called when exiting the EInt production.
	ExitEInt(c *EIntContext)

	// ExitEUnOp is called when exiting the EUnOp production.
	ExitEUnOp(c *EUnOpContext)

	// ExitEStr is called when exiting the EStr production.
	ExitEStr(c *EStrContext)

	// ExitEMulOp is called when exiting the EMulOp production.
	ExitEMulOp(c *EMulOpContext)

	// ExitEAnd is called when exiting the EAnd production.
	ExitEAnd(c *EAndContext)

	// ExitEParen is called when exiting the EParen production.
	ExitEParen(c *EParenContext)

	// ExitEFalse is called when exiting the EFalse production.
	ExitEFalse(c *EFalseContext)

	// ExitENew is called when exiting the ENew production.
	ExitENew(c *ENewContext)

	// ExitEAddOp is called when exiting the EAddOp production.
	ExitEAddOp(c *EAddOpContext)

	// ExitENull is called when exiting the ENull production.
	ExitENull(c *ENullContext)

	// ExitEFieldAccess is called when exiting the EFieldAccess production.
	ExitEFieldAccess(c *EFieldAccessContext)

	// ExitAddOp is called when exiting the addOp production.
	ExitAddOp(c *AddOpContext)

	// ExitMulOp is called when exiting the mulOp production.
	ExitMulOp(c *MulOpContext)

	// ExitRelOp is called when exiting the relOp production.
	ExitRelOp(c *RelOpContext)
}
