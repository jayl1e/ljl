package translate

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	ire "github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	irval "github.com/llir/llvm/ir/value"
	"jayjaylee.com/ljl/parser"
	"jayjaylee.com/ljl/sym"
	"strconv"
)

type LlvmIrGen struct {
	parser.BaseLjlVisitor
	*sym.SymbolTree

	module  *ir.Module
	current *ir.Block
	seq     int
	valDict map[antlr.ParseTree]sym.ValSym
}

func (v LlvmIrGen) TargetIL() string {
	return v.module.String()
}

func (v *LlvmIrGen) Seq() int {
	v.seq += 1
	return v.seq
}

var _ parser.LjlVisitor = &LlvmIrGen{}

func NewLlvmIrGen(symbols *sym.SymbolTree) *LlvmIrGen {
	l := &LlvmIrGen{}
	l.SymbolTree = symbols
	l.Init(l)
	l.valDict = map[antlr.ParseTree]sym.ValSym{}
	return l
}

func (v *LlvmIrGen) VisitFile(ctx *parser.FileContext) interface{} {
	v.Push(ctx, v.SM[ctx])
	v.module = ir.NewModule()
	v.VisitChildren(ctx)
	v.Pop()
	return nil
}

func (v *LlvmIrGen) VisitFuncDefine(ctx *parser.FuncDefineContext) interface{} {
	_f := v.SM[ctx.FuncProto()]
	if _, ok := _f.(*sym.FuncVal); !ok {
		panic(sym.ESymbolTypeError)
	}
	f := _f.(*sym.FuncVal).LlVal().(*ir.Func)
	f.Parent = v.module
	v.module.Funcs = append(v.module.Funcs, f)
	v.Push(ctx, v.SM[ctx.FuncProto()])
	blk := f.NewBlock("entry")
	v.current = blk
	v.Visit(ctx.Block())
	if v.valDict[ctx.Block()].Typ() != _f.(*sym.FuncVal).RetVal.Typ() {
		panic(fmt.Errorf("return type not match, expect %v, actual %v: %s",
			_f.(*sym.FuncVal).RetVal.Typ(),
			v.valDict[ctx.Block()].Typ(),
			ctx.FuncProto().(*parser.FuncProtoContext).ID().GetText()))
	}
	v.current.NewRet(v.valDict[ctx.Block()].LlVal())
	v.Pop()
	return nil
}

func (v *LlvmIrGen) VisitBlock(ctx *parser.BlockContext) interface{} {
	var last parser.IExprContext
	for _, exp := range ctx.AllExpr() {
		exp.Accept(v)
		last = exp
	}
	if last == nil {
		v.valDict[ctx] = sym.UnitVal
	} else {
		v.valDict[ctx] = v.valDict[last]
	}
	return nil
}

func (v *LlvmIrGen) VisitIfExpr(ctx *parser.IfExprContext) interface{} {
	v.Visit(ctx.Expr(0))
	flag := v.valDict[ctx.Expr(0)]
	if flag.Typ() != sym.BoolTyp {
		panic(fmt.Sprintf("condition is not bool: @%s", ctx.Expr(0).GetText()))
	}
	f := v.current.Parent
	thenBB := f.NewBlock("")
	elseBB := ir.NewBlock("")
	mergeBB := ir.NewBlock("")
	v.current.NewCondBr(flag.LlVal(), thenBB, elseBB)
	v.current = thenBB
	v.Visit(ctx.Expr(1))
	v.current.NewBr(mergeBB)
	f.Blocks = append(f.Blocks, elseBB)
	v.current = elseBB
	v.Visit(ctx.Expr(2))
	v.current.NewBr(mergeBB)
	f.Blocks = append(f.Blocks, mergeBB)
	v.current = mergeBB
	thenVal := v.valDict[ctx.Expr(1)]
	elseVal := v.valDict[ctx.Expr(2)]
	if thenVal.Typ() != elseVal.Typ() {
		panic("if expr type not match")
	}
	phi := v.current.NewPhi(&ir.Incoming{X: thenVal.LlVal(), Pred: thenBB},
		&ir.Incoming{X: elseVal.LlVal(), Pred: elseBB})
	ifVal := thenVal.Typ().NewInstance()
	err := ifVal.SetLlVal(phi)
	if err != nil {
		panic(err)
	}
	v.valDict[ctx] = ifVal
	return nil
}

func (v *LlvmIrGen) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	if ctx.INT() != nil {
		n, err := strconv.Atoi(ctx.INT().GetText())
		if err != nil {
			panic(err)
		}
		typ := sym.IntTyp
		val := typ.NewInstance()
		err = val.SetLlVal(constant.NewInt(types.I64, int64(n)))
		if err != nil {
			panic(err)
		}
		v.valDict[ctx] = val
	} else if ctx.FLOAT() != nil {
		n, err := strconv.ParseFloat(ctx.FLOAT().GetText(), 64)
		if err != nil {
			panic(err)
		}
		typ := sym.FloatTyp
		val := typ.NewInstance()
		err = val.SetLlVal(constant.NewFloat(types.Double, n))
		if err != nil {
			panic(err)
		}
		v.valDict[ctx] = val
	} else if ctx.STRING() != nil {
		unquote, err := strconv.Unquote(ctx.STRING().GetText())
		if err != nil {
			panic(err)
		}
		lit := unquote + "\x00"
		arr := constant.NewCharArrayFromString(lit)
		strCst := v.module.NewGlobalDef(fmt.Sprintf(".str.%d", v.Seq()), arr)
		charPtr := v.current.NewGetElementPtr(strCst.Typ.ElemType, strCst, constant.NewInt(types.I64, 0), constant.NewInt(types.I64, 0))
		typ := sym.StringTyp
		val := typ.NewInstance()
		err = val.SetLlVal(charPtr)
		if err != nil {
			return nil
		}
		v.valDict[ctx] = val
	} else {
		panic(fmt.Sprintf("unsupported literal: %s", ctx.GetText()))
	}
	return v.VisitChildren(ctx)
}

func (v *LlvmIrGen) VisitValRefExpr(ctx *parser.ValRefExprContext) interface{} {
	name := ctx.ID().GetText()
	s := v.Resolve(ctx, name)
	if s == nil {
		panic(fmt.Sprintf("symbol not found:%s", name))
	}
	if vs, ok := s.(sym.ValSym); ok {
		v.valDict[ctx] = vs
	} else {
		panic(fmt.Sprintf("should be val ref: %s", name))
	}
	return v.VisitChildren(ctx)
}

func (v *LlvmIrGen) VisitLitExpr(ctx *parser.LitExprContext) interface{} {
	v.VisitChildren(ctx)
	v.valDict[ctx] = v.valDict[ctx.Literal()]
	return nil
}

func (v *LlvmIrGen) VisitCmpExpr(ctx *parser.CmpExprContext) interface{} {
	v.VisitChildren(ctx)
	left := v.valDict[ctx.Expr(0)]
	right := v.valDict[ctx.Expr(1)]
	if left.Typ() != right.Typ() {
		panic(fmt.Sprintf("can not compare different type"))
	}
	cmpType := left.Typ()
	var val sym.ValSym

	switch cmpType {
	case sym.IntTyp:
		var op ire.IPred
		switch ctx.GetOp().GetText() {
		case ">":
			op = ire.IPredSGT
		case "<":
			op = ire.IPredSLT
		case ">=":
			op = ire.IPredSGE
		case "<=":
			op = ire.IPredSLE
		default:
			panic(fmt.Sprintf("unknown op: %s", ctx.GetOp().GetText()))
		}
		llv := v.current.NewICmp(op, left.LlVal(), right.LlVal())
		val = sym.BoolTyp.NewInstance()
		err := val.SetLlVal(llv)
		if err != nil {
			panic(err)
		}
	case sym.FloatTyp:
		var op ire.FPred
		switch ctx.GetOp().GetText() {
		case ">":
			op = ire.FPredOGT
		case "<":
			op = ire.FPredOLT
		case ">=":
			op = ire.FPredOLE
		case "<=":
			op = ire.FPredOLE
		default:
			panic(fmt.Sprintf("unknown op: %s", ctx.GetOp().GetText()))
		}
		llv := v.current.NewFCmp(op, left.LlVal(), right.LlVal())
		val = sym.BoolTyp.NewInstance()
		err := val.SetLlVal(llv)
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintf("unsuported type cmp %s", ctx.GetText()))
	}
	v.valDict[ctx] = val
	return nil
}

func (v *LlvmIrGen) VisitEqExpr(ctx *parser.EqExprContext) interface{} {
	v.VisitChildren(ctx)
	left := v.valDict[ctx.Expr(0)]
	right := v.valDict[ctx.Expr(1)]
	if left.Typ() != right.Typ() {
		panic(fmt.Sprintf("can not compare different type"))
	}
	cmpType := left.Typ()
	var val sym.ValSym

	switch cmpType {
	case sym.IntTyp:
		var op ire.IPred
		switch ctx.GetOp().GetText() {
		case "==":
			op = ire.IPredEQ
		case "!=":
			op = ire.IPredNE
		default:
			panic(fmt.Sprintf("unknown op: %s", ctx.GetOp().GetText()))
		}
		llv := v.current.NewICmp(op, left.LlVal(), right.LlVal())
		val = cmpType.NewInstance()
		err := val.SetLlVal(llv)
		if err != nil {
			panic(err)
		}

	case sym.FloatTyp:
		var op ire.FPred
		switch ctx.GetOp().GetText() {
		case "==":
			op = ire.FPredOEQ
		case "!=":
			op = ire.FPredONE
		default:
			panic(fmt.Sprintf("unknown op: %s", ctx.GetOp().GetText()))
		}
		llv := v.current.NewFCmp(op, left.LlVal(), right.LlVal())
		val = cmpType.NewInstance()
		err := val.SetLlVal(llv)
		if err != nil {
			panic(err)
		}
	case sym.BoolTyp:
		var op ire.IPred
		switch ctx.GetOp().GetText() {
		case "==":
			op = ire.IPredEQ
		case "!=":
			op = ire.IPredNE
		default:
			panic(fmt.Sprintf("unknown op: %s", ctx.GetOp().GetText()))
		}
		llv := v.current.NewICmp(op, left.LlVal(), right.LlVal())
		val = cmpType.NewInstance()
		err := val.SetLlVal(llv)
		if err != nil {
			panic(err)
		}
	default:
		panic(fmt.Sprintf("unsuported type cmp %s", ctx.GetText()))
	}
	v.valDict[ctx] = val
	return nil
}

func (v *LlvmIrGen) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	v.VisitChildren(ctx)
	right := v.valDict[ctx.Expr()]
	if right.Typ() != sym.BoolTyp {
		panic(fmt.Sprintf("not op on non-bool: %s", ctx.GetText()))
	}
	llv := v.current.NewXor(right.LlVal(), constant.True)
	val := sym.BoolTyp.NewInstance()
	err := val.SetLlVal(llv)
	if err != nil {
		panic(err)
	}
	v.valDict[ctx] = val
	return nil
}

func (v *LlvmIrGen) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	v.VisitChildren(ctx)
	left := v.valDict[ctx.Expr(0)]
	right := v.valDict[ctx.Expr(1)]
	if left.Typ() != sym.BoolTyp || right.Typ() != sym.BoolTyp {
		panic(fmt.Sprintf("not op on non-bool: %s", ctx.GetText()))
	}
	llv := v.current.NewAnd(left.LlVal(), right.LlVal())
	val := sym.BoolTyp.NewInstance()
	err := val.SetLlVal(llv)
	if err != nil {
		panic(err)
	}
	v.valDict[ctx] = val
	return nil
}

func (v *LlvmIrGen) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	v.VisitChildren(ctx)
	left := v.valDict[ctx.Expr(0)]
	right := v.valDict[ctx.Expr(1)]
	if left.Typ() != sym.BoolTyp || right.Typ() != sym.BoolTyp {
		panic(fmt.Sprintf("not op on non-bool: %s", ctx.GetText()))
	}
	llv := v.current.NewOr(left.LlVal(), right.LlVal())
	val := sym.BoolTyp.NewInstance()
	err := val.SetLlVal(llv)
	if err != nil {
		panic(err)
	}
	v.valDict[ctx] = val
	return nil
}

func (v *LlvmIrGen) VisitMulDivExpr(ctx *parser.MulDivExprContext) interface{} {
	v.VisitChildren(ctx)
	left := v.valDict[ctx.Expr(0)]
	right := v.valDict[ctx.Expr(1)]
	if left.Typ() != right.Typ() {
		panic(fmt.Sprintf("can not compare different type"))
	}
	sideType := left.Typ()
	var valSym sym.ValSym = sideType.NewInstance()
	var val irval.Value

	switch sideType {
	case sym.IntTyp:
		switch ctx.GetOp().GetText() {
		case "*":
			val = v.current.NewMul(left.LlVal(), right.LlVal())
		case "/":
			val = v.current.NewSDiv(left.LlVal(), right.LlVal())
		default:
			panic(fmt.Sprintf("unknown op: %s", ctx.GetOp().GetText()))
		}
	case sym.FloatTyp:
		switch ctx.GetOp().GetText() {
		case "*":
			val = v.current.NewFMul(left.LlVal(), right.LlVal())
		case "/":
			val = v.current.NewFDiv(left.LlVal(), right.LlVal())
		default:
			panic(fmt.Sprintf("unknown op: %s", ctx.GetOp().GetText()))
		}
	default:
		panic(fmt.Sprintf("unsuported type cmp %s", ctx.GetText()))
	}
	err := valSym.SetLlVal(val)
	if err != nil {
		panic(err)
	}

	v.valDict[ctx] = valSym
	return nil
}

func (v *LlvmIrGen) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	v.VisitChildren(ctx)
	left := v.valDict[ctx.Expr(0)]
	right := v.valDict[ctx.Expr(1)]
	if left.Typ() != right.Typ() {
		panic(fmt.Sprintf("can not compare different type"))
	}
	sideType := left.Typ()
	var valSym sym.ValSym = sideType.NewInstance()
	var val irval.Value

	switch sideType {
	case sym.IntTyp:
		switch ctx.GetOp().GetText() {
		case "+":
			val = v.current.NewAdd(left.LlVal(), right.LlVal())
		case "-":
			val = v.current.NewSub(left.LlVal(), right.LlVal())
		default:
			panic(fmt.Sprintf("unknown op: %s", ctx.GetOp().GetText()))
		}
	case sym.FloatTyp:
		switch ctx.GetOp().GetText() {
		case "+":
			val = v.current.NewFAdd(left.LlVal(), right.LlVal())
		case "-":
			val = v.current.NewFSub(left.LlVal(), right.LlVal())
		default:
			panic(fmt.Sprintf("unknown op: %s", ctx.GetOp().GetText()))
		}
	default:
		panic(fmt.Sprintf("unsuported type cmp %s", ctx.GetText()))
	}
	err := valSym.SetLlVal(val)
	if err != nil {
		panic(err)
	}

	v.valDict[ctx] = valSym
	return nil
}

func (v *LlvmIrGen) VisitTupleVal(ctx *parser.TupleValContext) interface{} {
	v.VisitChildren(ctx)
	mem := make([]sym.ValSym, 0, len(ctx.AllExpr()))
	for _, s := range ctx.AllExpr() {
		mem = append(mem, v.valDict[s])
	}
	_val := &sym.TupleVal{Members: mem}
	val := _val.Reduce()
	if tv, ok := val.(*sym.TupleVal); ok {
		tv.SetLlValFunc(func() irval.Value {
			var crt irval.Value = constant.NewUndef(tv.Typ().LlTyp())
			for i, member := range tv.Members {
				crt = v.current.NewInsertValue(crt, member.LlVal(), uint64(i))
			}
			return crt
		})
	}
	v.valDict[ctx] = val
	return nil
}

func (v *LlvmIrGen) VisitTupleExpr(ctx *parser.TupleExprContext) interface{} {
	v.VisitChildren(ctx)
	v.valDict[ctx] = v.valDict[ctx.TupleVal()]
	return nil
}

func (v *LlvmIrGen) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	v.VisitChildren(ctx)
	fs := v.Resolve(ctx, ctx.ID().GetText()).(*sym.FuncVal)
	arg := v.valDict[ctx.TupleVal()]
	if fs.ArgVal.Typ() != arg.Typ() {
		panic(fmt.Errorf("func call arg type mismatch, expect %v, got %v: %s", fs.ArgVal.Typ(), arg.Typ(), ctx.GetText()))
	}
	args := make([]irval.Value, 0)
	if tv, ok := arg.(*sym.TupleVal); ok {
		for i, _ := range tv.Members {
			args = append(args, tv.GetField(fmt.Sprintf("_%d", i), v.current).LlVal())
		}
	} else if arg != sym.UnitVal {
		args = append(args, arg.LlVal())
	}
	llv := v.current.NewCall(fs.LlVal(), args...)
	val := fs.RetVal.Typ().NewInstance()
	err := val.SetLlVal(llv)
	if err != nil {
		panic(err)
	}
	v.valDict[ctx] = val
	return nil
}

func (v *LlvmIrGen) VisitFuncExtern(ctx *parser.FuncExternContext) interface{} {
	_f := v.SM[ctx.FuncProto()]
	if fs, ok := _f.(*sym.FuncVal); ok {
		f := fs.LlVal().(*ir.Func)
		f.Parent = v.module
		v.module.Funcs = append(v.module.Funcs, f)
	} else {
		panic(sym.ESymbolTypeError)
	}
	return nil
}

func (v *LlvmIrGen) VisitStructDefine(ctx *parser.StructDefineContext) interface{} {
	_s := v.Resolve(ctx, ctx.ID().GetText())
	if st, ok := _s.(*sym.StructTyp); ok {
		v.module.NewTypeDef(ctx.ID().GetText(), st.LlTyp())
	} else {
		panic(sym.ESymbolTypeError)
	}
	return nil
}

func (v *LlvmIrGen) VisitMemberExpr(ctx *parser.MemberExprContext) interface{} {
	ctx.Expr().Accept(v)
	lv := v.valDict[ctx.Expr()]
	switch tlv := lv.(type) {
	case *sym.StructVal:
		mv := tlv.GetField(ctx.ID().GetText(), v.current)
		v.valDict[ctx] = mv.(sym.ValSym)
	case *sym.TupleVal:
		mv := tlv.GetField(ctx.ID().GetText(), v.current)
		v.valDict[ctx] = mv.(sym.ValSym)
	default:
		panic(fmt.Errorf("%s is not an agregate type: %s", ctx.Expr().GetText(), ctx.GetText()))
	}
	return nil
}
