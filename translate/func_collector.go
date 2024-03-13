package translate

import (
	"github.com/llir/llvm/ir"
	"jayjaylee.com/ljl/parser"
	"jayjaylee.com/ljl/sym"
)

type FuncCollector struct {
	parser.BaseLjlVisitor
	*sym.SymbolTree
}

var _ parser.LjlVisitor = &FuncCollector{}

func NewFuncCollector(st *sym.SymbolTree) *FuncCollector {
	l := &FuncCollector{}
	l.SymbolTree = st
	l.Init(l)
	return l
}

func (v *FuncCollector) VisitFile(ctx *parser.FileContext) interface{} {
	v.Push(ctx, v.SM[ctx])
	return v.VisitChildren(ctx)
}

func (v *FuncCollector) VisitFuncExtern(ctx *parser.FuncExternContext) interface{} {
	v.VisitChildren(ctx)
	return nil
}

func (v *FuncCollector) VisitFuncProto(ctx *parser.FuncProtoContext) interface{} {
	f := &sym.FuncVal{ArgVal: &sym.TupleVal{Members: make([]sym.ValSym, 0)}}
	err := v.Add(ctx.ID().GetText(), f)
	if err != nil {
		panic(err)
	}
	f.SetParent(v.ST)
	v.Push(ctx, f)
	ctx.ParamList().Accept(v)
	ret := v.GetTypSym(ctx.Typ())
	f.RetVal = ret.NewInstance()
	f.AutoFill()
	params := []*ir.Param{}
	if tv, ok := f.ArgVal.(*sym.TupleVal); ok {
		for _, member := range tv.Members {
			p := &ir.Param{Typ: member.Typ().LlTyp()}
			err := member.SetLlVal(p)
			if err != nil {
				panic(err)
			}
			params = append(params, p)
		}
	} else if f.ArgVal != sym.UnitVal {
		p := &ir.Param{Typ: f.ArgVal.Typ().LlTyp()}
		err := f.ArgVal.SetLlVal(p)
		if err != nil {
			panic(err)
		}
		params = append(params, p)
	}

	llv := ir.NewFunc(ctx.ID().GetText(), f.RetVal.Typ().LlTyp(), params...)
	err = f.SetLlVal(llv)
	if err != nil {
		panic(err)
	}
	v.Pop()
	return nil
}

func (v *FuncCollector) VisitParam(ctx *parser.ParamContext) interface{} {
	err := v.Add(ctx.ID().GetText(), v.GetTypSym(ctx.Typ()))
	if err != nil {
		panic(err)
	}
	return nil
}

func (v *FuncCollector) VisitFuncDefine(ctx *parser.FuncDefineContext) interface{} {
	v.VisitChildren(ctx)
	return nil
}
