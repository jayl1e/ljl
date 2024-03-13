package runtime

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseParseTreeVisitor struct{
	self antlr.ParseTreeVisitor
}

var _ antlr.ParseTreeVisitor = &BaseParseTreeVisitor{}

func (v *BaseParseTreeVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v.self)
}
func (v *BaseParseTreeVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for _, tree := range node.GetChildren() {
		v.Visit(tree.(antlr.ParseTree))
	}
	return nil
}
func (v *BaseParseTreeVisitor) VisitTerminal(node antlr.TerminalNode) interface{} { return nil }
func (v *BaseParseTreeVisitor) VisitErrorNode(node antlr.ErrorNode) interface{}   { return nil }
func (v *BaseParseTreeVisitor) Init(self antlr.ParseTreeVisitor){ v.self=self }