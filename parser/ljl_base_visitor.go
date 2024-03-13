// Code generated from /Volumes/Code/Workspace/lijie0123/ljl/Ljl.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Ljl
import (
	"jayjaylee.com/ljl/runtime"
)

type BaseLjlVisitor struct {
	runtime.BaseParseTreeVisitor
}

func (v *BaseLjlVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitTop(ctx *TopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitStructDefine(ctx *StructDefineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitTypeAlias(ctx *TypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitStructMem(ctx *StructMemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitFuncExtern(ctx *FuncExternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitFuncDefine(ctx *FuncDefineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitFuncProto(ctx *FuncProtoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitTyp(ctx *TypContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitMemberExpr(ctx *MemberExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitCmpExpr(ctx *CmpExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitEqExpr(ctx *EqExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitValRefExpr(ctx *ValRefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitIfExpr(ctx *IfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitBlockExpr(ctx *BlockExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitTupleExpr(ctx *TupleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitLitExpr(ctx *LitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitMulDivExpr(ctx *MulDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitTupleVal(ctx *TupleValContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLjlVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
