package translate

import (
	"jayjaylee.com/ljl/parser"
	"jayjaylee.com/ljl/sym"
)

type TypeCollector struct {
	parser.BaseLjlVisitor
	*sym.SymbolTree
}

var _ parser.LjlVisitor = &TypeCollector{}

func NewTypeCollector(symTree *sym.SymbolTree) *TypeCollector {
	l := &TypeCollector{}
	l.Init(l)
	l.SymbolTree = symTree
	return l
}

func (v *TypeCollector) VisitFile(ctx *parser.FileContext) interface{} {
	v.Push(ctx, sym.NewSymTable(v.ST))
	v.VisitChildren(ctx)
	v.Pop()
	return nil
}

func (v *TypeCollector) VisitStructDefine(ctx *parser.StructDefineContext) interface{} {
	t := &sym.StructTyp{}
	t.SetParent(v.ST)
	err := v.Add(ctx.ID().GetText(), t)
	if err != nil {
		panic(err)
	}
	v.Push(ctx, t)
	v.VisitChildren(ctx)
	err = t.Fill()
	if err != nil {
		panic(err)
	}
	return nil
}

func (v *TypeCollector) VisitStructMem(ctx *parser.StructMemContext) interface{} {
	tpSym := ctx.Typ().Accept(v)
	err := v.Add(ctx.ID().GetText(), tpSym.(sym.TypSym))
	if err != nil {
		panic(err)
	}
	return nil
}

func (v *TypeCollector) VisitTypeAlias(ctx *parser.TypeAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *TypeCollector) VisitTyp(ctx *parser.TypContext) interface{} {
	return v.SymbolTree.GetTypSym(ctx)
}
