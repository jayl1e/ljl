package sym

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"jayjaylee.com/ljl/parser"
)

type SymbolTree struct {
	SM    map[antlr.ParseTree]SymTab
	ST    SymTab
	stack []SymTab
}

func (t *SymbolTree) Push(at antlr.ParseTree, st SymTab) {
	if st == nil {
		panic(fmt.Errorf("push nil sym table"))
	}
	t.SM[at] = st
	t.ST = st
	t.stack = append(t.stack, st)
}

func (t *SymbolTree) Pop() SymTab {
	tmp := t.ST
	t.stack = t.stack[:len(t.stack)-1]
	t.ST = t.stack[len(t.stack)-1]
	return tmp
}

func (t *SymbolTree) Add(name string, sym Sym) error {
	return t.ST.Add(name, sym)
}

func (t SymbolTree) Resolve(ctx antlr.ParseTree, name string) Sym {
	for ctx != nil {
		if st, ok := t.SM[ctx]; ok {
			return st.Resolve(name)
		}
		if ctx.GetParent() == nil {
			return nil
		}
		ctx = ctx.GetParent().(antlr.ParseTree)
	}
	return nil
}

func (t *SymbolTree) InitSymTree() {
	t.SM = map[antlr.ParseTree]SymTab{}
	t.Push(nil, NewBuiltinTable())
}

func (t *SymbolTree) GetTypSym(iCtx parser.ITypContext) TypSym {
	ctx := iCtx.(*parser.TypContext)
	var tp TypSym
	tps := []TypSym{}
	for _, name := range ctx.AllID() {
		_tpsym := t.Resolve(ctx, name.GetText())
		if _tpsym == nil {
			panic(fmt.Errorf("symbol %s not found at %s", name.GetText(), ctx.GetText()))
		}
		if _tpsym.Type() != SType {
			panic(fmt.Errorf("symbol %s should be type", name.GetText()))
		}
		tps = append(tps, _tpsym.(TypSym))
	}
	tp = GetTupleTyp(tps)
	return tp
}

func (t *SymbolTree) Recover(ctx antlr.ParseTree) SymTab {
	if st, ok := t.SM[ctx]; ok {
		t.ST = st
		return st
	}
	return nil
}
