// Code generated from /Volumes/Code/Workspace/lijie0123/ljl/Ljl.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Ljl
import "github.com/antlr/antlr4/runtime/Go/antlr"
// A complete Visitor for a parse tree produced by LjlParser.
type LjlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LjlParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by LjlParser#top.
	VisitTop(ctx *TopContext) interface{}

	// Visit a parse tree produced by LjlParser#structDefine.
	VisitStructDefine(ctx *StructDefineContext) interface{}

	// Visit a parse tree produced by LjlParser#typeAlias.
	VisitTypeAlias(ctx *TypeAliasContext) interface{}

	// Visit a parse tree produced by LjlParser#structMem.
	VisitStructMem(ctx *StructMemContext) interface{}

	// Visit a parse tree produced by LjlParser#funcExtern.
	VisitFuncExtern(ctx *FuncExternContext) interface{}

	// Visit a parse tree produced by LjlParser#funcDefine.
	VisitFuncDefine(ctx *FuncDefineContext) interface{}

	// Visit a parse tree produced by LjlParser#funcProto.
	VisitFuncProto(ctx *FuncProtoContext) interface{}

	// Visit a parse tree produced by LjlParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by LjlParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by LjlParser#typ.
	VisitTyp(ctx *TypContext) interface{}

	// Visit a parse tree produced by LjlParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by LjlParser#addSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by LjlParser#memberExpr.
	VisitMemberExpr(ctx *MemberExprContext) interface{}

	// Visit a parse tree produced by LjlParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by LjlParser#cmpExpr.
	VisitCmpExpr(ctx *CmpExprContext) interface{}

	// Visit a parse tree produced by LjlParser#eqExpr.
	VisitEqExpr(ctx *EqExprContext) interface{}

	// Visit a parse tree produced by LjlParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by LjlParser#valRefExpr.
	VisitValRefExpr(ctx *ValRefExprContext) interface{}

	// Visit a parse tree produced by LjlParser#ifExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by LjlParser#blockExpr.
	VisitBlockExpr(ctx *BlockExprContext) interface{}

	// Visit a parse tree produced by LjlParser#callExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by LjlParser#tupleExpr.
	VisitTupleExpr(ctx *TupleExprContext) interface{}

	// Visit a parse tree produced by LjlParser#litExpr.
	VisitLitExpr(ctx *LitExprContext) interface{}

	// Visit a parse tree produced by LjlParser#mulDivExpr.
	VisitMulDivExpr(ctx *MulDivExprContext) interface{}

	// Visit a parse tree produced by LjlParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by LjlParser#tupleVal.
	VisitTupleVal(ctx *TupleValContext) interface{}

	// Visit a parse tree produced by LjlParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

}