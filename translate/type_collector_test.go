package translate

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"ljl/parser"
	sym "ljl/sym"
	"testing"
)



func TestTypeCollector(t *testing.T) {
	ist:=antlr.NewInputStream(testFile)
	lexer:= parser.NewLjlLexer(ist)
	tokens := antlr.NewCommonTokenStream(lexer,0)
	myParser := parser.NewLjlParser(tokens)
	ifilectx := myParser.File()
	symTree :=&sym.SymbolTree{}
	symTree.InitSymTree()
	v1:=NewTypeCollector(symTree)
	v1.Visit(ifilectx)
	println(v1.SM)
}