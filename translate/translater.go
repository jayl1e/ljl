package translate

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"jayjaylee.com/ljl/parser"
	"jayjaylee.com/ljl/sym"
)

func Translate(src string) string {
	ist := antlr.NewInputStream(src)
	lexer := parser.NewLjlLexer(ist)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	myParser := parser.NewLjlParser(tokens)
	file := myParser.File()
	symTree := &sym.SymbolTree{}
	symTree.InitSymTree()
	v1 := NewTypeCollector(symTree)
	v1.Visit(file)
	v2 := NewFuncCollector(symTree)
	v2.Visit(file)
	v3 := NewLlvmIrGen(symTree)
	v3.Visit(file)
	return v3.TargetIL()
}
