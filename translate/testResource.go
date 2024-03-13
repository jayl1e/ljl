package translate

import (
	_ "embed"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"jayjaylee.com/ljl/parser"
	"jayjaylee.com/ljl/sym"
)

//go:embed test.ljl
var testFile string

func getTestFileCtx() (parser.IFileContext, *sym.SymbolTree) {

	ist := antlr.NewInputStream(testFile)
	lexer := parser.NewLjlLexer(ist)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	myParser := parser.NewLjlParser(tokens)
	iFileContext := myParser.File()
	symTree := &sym.SymbolTree{}
	symTree.InitSymTree()
	return iFileContext, symTree
}
