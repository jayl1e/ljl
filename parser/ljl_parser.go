// Code generated from /Volumes/Code/Workspace/lijie0123/ljl/Ljl.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Ljl
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 181, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 7, 2, 32, 10, 2, 12, 2, 14, 2, 
	35, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 44, 10, 3, 3, 
	4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 51, 10, 4, 12, 4, 14, 4, 54, 11, 4, 3, 
	4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 61, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 85, 10, 9, 12, 9, 14, 9, 88, 11, 
	9, 5, 9, 90, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 
	11, 3, 11, 3, 11, 3, 11, 7, 11, 103, 10, 11, 12, 11, 14, 11, 106, 11, 11, 
	3, 11, 3, 11, 3, 11, 5, 11, 111, 10, 11, 3, 12, 3, 12, 7, 12, 115, 10, 
	12, 12, 12, 14, 12, 118, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 
	3, 13, 3, 13, 5, 13, 138, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 161, 10, 13, 12, 13, 14, 13, 
	164, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 170, 10, 14, 12, 14, 14, 
	14, 173, 11, 14, 5, 14, 175, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 
	2, 3, 24, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2, 7, 
	3, 2, 16, 17, 3, 2, 18, 19, 3, 2, 20, 23, 3, 2, 24, 25, 3, 2, 33, 35, 2, 
	192, 2, 33, 3, 2, 2, 2, 4, 43, 3, 2, 2, 2, 6, 60, 3, 2, 2, 2, 8, 62, 3, 
	2, 2, 2, 10, 66, 3, 2, 2, 2, 12, 69, 3, 2, 2, 2, 14, 72, 3, 2, 2, 2, 16, 
	89, 3, 2, 2, 2, 18, 91, 3, 2, 2, 2, 20, 110, 3, 2, 2, 2, 22, 112, 3, 2, 
	2, 2, 24, 137, 3, 2, 2, 2, 26, 165, 3, 2, 2, 2, 28, 178, 3, 2, 2, 2, 30, 
	32, 5, 4, 3, 2, 31, 30, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31, 3, 2, 2, 
	2, 33, 34, 3, 2, 2, 2, 34, 36, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 37, 
	7, 2, 2, 3, 37, 3, 3, 2, 2, 2, 38, 44, 5, 12, 7, 2, 39, 40, 5, 10, 6, 2, 
	40, 41, 7, 3, 2, 2, 41, 44, 3, 2, 2, 2, 42, 44, 5, 6, 4, 2, 43, 38, 3, 
	2, 2, 2, 43, 39, 3, 2, 2, 2, 43, 42, 3, 2, 2, 2, 44, 5, 3, 2, 2, 2, 45, 
	46, 7, 4, 2, 2, 46, 47, 7, 36, 2, 2, 47, 48, 7, 5, 2, 2, 48, 52, 7, 6, 
	2, 2, 49, 51, 5, 8, 5, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 
	3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 
	55, 61, 7, 7, 2, 2, 56, 57, 7, 4, 2, 2, 57, 58, 7, 36, 2, 2, 58, 59, 7, 
	8, 2, 2, 59, 61, 5, 20, 11, 2, 60, 45, 3, 2, 2, 2, 60, 56, 3, 2, 2, 2, 
	61, 7, 3, 2, 2, 2, 62, 63, 7, 36, 2, 2, 63, 64, 7, 9, 2, 2, 64, 65, 5, 
	20, 11, 2, 65, 9, 3, 2, 2, 2, 66, 67, 7, 10, 2, 2, 67, 68, 5, 14, 8, 2, 
	68, 11, 3, 2, 2, 2, 69, 70, 5, 14, 8, 2, 70, 71, 5, 22, 12, 2, 71, 13, 
	3, 2, 2, 2, 72, 73, 7, 11, 2, 2, 73, 74, 7, 36, 2, 2, 74, 75, 7, 12, 2, 
	2, 75, 76, 5, 16, 9, 2, 76, 77, 7, 13, 2, 2, 77, 78, 7, 14, 2, 2, 78, 79, 
	5, 20, 11, 2, 79, 15, 3, 2, 2, 2, 80, 90, 3, 2, 2, 2, 81, 86, 5, 18, 10, 
	2, 82, 83, 7, 15, 2, 2, 83, 85, 5, 18, 10, 2, 84, 82, 3, 2, 2, 2, 85, 88, 
	3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 
	88, 86, 3, 2, 2, 2, 89, 80, 3, 2, 2, 2, 89, 81, 3, 2, 2, 2, 90, 17, 3, 
	2, 2, 2, 91, 92, 7, 36, 2, 2, 92, 93, 7, 9, 2, 2, 93, 94, 5, 20, 11, 2, 
	94, 19, 3, 2, 2, 2, 95, 111, 7, 36, 2, 2, 96, 97, 7, 12, 2, 2, 97, 98, 
	7, 36, 2, 2, 98, 99, 7, 15, 2, 2, 99, 104, 7, 36, 2, 2, 100, 101, 7, 15, 
	2, 2, 101, 103, 7, 36, 2, 2, 102, 100, 3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 
	104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 107, 3, 2, 2, 2, 106, 
	104, 3, 2, 2, 2, 107, 111, 7, 13, 2, 2, 108, 109, 7, 12, 2, 2, 109, 111, 
	7, 13, 2, 2, 110, 95, 3, 2, 2, 2, 110, 96, 3, 2, 2, 2, 110, 108, 3, 2, 
	2, 2, 111, 21, 3, 2, 2, 2, 112, 116, 7, 6, 2, 2, 113, 115, 5, 24, 13, 2, 
	114, 113, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2, 116, 
	117, 3, 2, 2, 2, 117, 119, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 119, 120, 
	7, 7, 2, 2, 120, 23, 3, 2, 2, 2, 121, 122, 8, 13, 1, 2, 122, 138, 5, 22, 
	12, 2, 123, 124, 7, 26, 2, 2, 124, 138, 5, 24, 13, 11, 125, 126, 7, 36, 
	2, 2, 126, 138, 5, 26, 14, 2, 127, 138, 7, 36, 2, 2, 128, 138, 5, 28, 15, 
	2, 129, 130, 7, 29, 2, 2, 130, 131, 5, 24, 13, 2, 131, 132, 7, 30, 2, 2, 
	132, 133, 5, 24, 13, 2, 133, 134, 7, 31, 2, 2, 134, 135, 5, 24, 13, 5, 
	135, 138, 3, 2, 2, 2, 136, 138, 5, 26, 14, 2, 137, 121, 3, 2, 2, 2, 137, 
	123, 3, 2, 2, 2, 137, 125, 3, 2, 2, 2, 137, 127, 3, 2, 2, 2, 137, 128, 
	3, 2, 2, 2, 137, 129, 3, 2, 2, 2, 137, 136, 3, 2, 2, 2, 138, 162, 3, 2, 
	2, 2, 139, 140, 12, 15, 2, 2, 140, 141, 9, 2, 2, 2, 141, 161, 5, 24, 13, 
	16, 142, 143, 12, 14, 2, 2, 143, 144, 9, 3, 2, 2, 144, 161, 5, 24, 13, 
	15, 145, 146, 12, 13, 2, 2, 146, 147, 9, 4, 2, 2, 147, 161, 5, 24, 13, 
	14, 148, 149, 12, 12, 2, 2, 149, 150, 9, 5, 2, 2, 150, 161, 5, 24, 13, 
	13, 151, 152, 12, 10, 2, 2, 152, 153, 7, 27, 2, 2, 153, 161, 5, 24, 13, 
	11, 154, 155, 12, 9, 2, 2, 155, 156, 7, 28, 2, 2, 156, 161, 5, 24, 13, 
	10, 157, 158, 12, 3, 2, 2, 158, 159, 7, 32, 2, 2, 159, 161, 7, 36, 2, 2, 
	160, 139, 3, 2, 2, 2, 160, 142, 3, 2, 2, 2, 160, 145, 3, 2, 2, 2, 160, 
	148, 3, 2, 2, 2, 160, 151, 3, 2, 2, 2, 160, 154, 3, 2, 2, 2, 160, 157, 
	3, 2, 2, 2, 161, 164, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2, 
	2, 2, 163, 25, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 165, 174, 7, 12, 2, 2, 
	166, 171, 5, 24, 13, 2, 167, 168, 7, 15, 2, 2, 168, 170, 5, 24, 13, 2, 
	169, 167, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 
	172, 3, 2, 2, 2, 172, 175, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 174, 166, 
	3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 177, 7, 13, 
	2, 2, 177, 27, 3, 2, 2, 2, 178, 179, 9, 6, 2, 2, 179, 29, 3, 2, 2, 2, 16, 
	33, 43, 52, 60, 86, 89, 104, 110, 116, 137, 160, 162, 171, 174,
}
var literalNames = []string{
	"", "';'", "'type'", "'struct'", "'{'", "'}'", "'='", "':'", "'extern'", 
	"'fn'", "'('", "')'", "'->'", "','", "'*'", "'/'", "'+'", "'-'", "'>'", 
	"'<'", "'>='", "'<='", "'=='", "'!='", "'!'", "'&&'", "'||'", "'if'", "'then'", 
	"'else'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "INT", "FLOAT", "STRING", 
	"ID", "COMMENT", "WS",
}

var ruleNames = []string{
	"file", "top", "typeDefine", "structMem", "funcExtern", "funcDefine", "funcProto", 
	"paramList", "param", "typ", "block", "expr", "tupleVal", "literal",
}
type LjlParser struct {
	*antlr.BaseParser
}

// NewLjlParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *LjlParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLjlParser(input antlr.TokenStream) *LjlParser {
	this := new(LjlParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Ljl.g4"

	return this
}


// LjlParser tokens.
const (
	LjlParserEOF = antlr.TokenEOF
	LjlParserT__0 = 1
	LjlParserT__1 = 2
	LjlParserT__2 = 3
	LjlParserT__3 = 4
	LjlParserT__4 = 5
	LjlParserT__5 = 6
	LjlParserT__6 = 7
	LjlParserT__7 = 8
	LjlParserT__8 = 9
	LjlParserT__9 = 10
	LjlParserT__10 = 11
	LjlParserT__11 = 12
	LjlParserT__12 = 13
	LjlParserT__13 = 14
	LjlParserT__14 = 15
	LjlParserT__15 = 16
	LjlParserT__16 = 17
	LjlParserT__17 = 18
	LjlParserT__18 = 19
	LjlParserT__19 = 20
	LjlParserT__20 = 21
	LjlParserT__21 = 22
	LjlParserT__22 = 23
	LjlParserT__23 = 24
	LjlParserT__24 = 25
	LjlParserT__25 = 26
	LjlParserT__26 = 27
	LjlParserT__27 = 28
	LjlParserT__28 = 29
	LjlParserT__29 = 30
	LjlParserINT = 31
	LjlParserFLOAT = 32
	LjlParserSTRING = 33
	LjlParserID = 34
	LjlParserCOMMENT = 35
	LjlParserWS = 36
)

// LjlParser rules.
const (
	LjlParserRULE_file = 0
	LjlParserRULE_top = 1
	LjlParserRULE_typeDefine = 2
	LjlParserRULE_structMem = 3
	LjlParserRULE_funcExtern = 4
	LjlParserRULE_funcDefine = 5
	LjlParserRULE_funcProto = 6
	LjlParserRULE_paramList = 7
	LjlParserRULE_param = 8
	LjlParserRULE_typ = 9
	LjlParserRULE_block = 10
	LjlParserRULE_expr = 11
	LjlParserRULE_tupleVal = 12
	LjlParserRULE_literal = 13
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(LjlParserEOF, 0)
}

func (s *FileContext) AllTop() []ITopContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITopContext)(nil)).Elem())
	var tst = make([]ITopContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITopContext)
		}
	}

	return tst
}

func (s *FileContext) Top(i int) ITopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITopContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITopContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LjlParserRULE_file)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << LjlParserT__1) | (1 << LjlParserT__7) | (1 << LjlParserT__8))) != 0) {
		{
			p.SetState(28)
			p.Top()
		}


		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Match(LjlParserEOF)
	}



	return localctx
}


// ITopContext is an interface to support dynamic dispatch.
type ITopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTopContext differentiates from other interfaces.
	IsTopContext()
}

type TopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopContext() *TopContext {
	var p = new(TopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_top
	return p
}

func (*TopContext) IsTopContext() {}

func NewTopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopContext {
	var p = new(TopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_top

	return p
}

func (s *TopContext) GetParser() antlr.Parser { return s.parser }

func (s *TopContext) FuncDefine() IFuncDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDefineContext)
}

func (s *TopContext) FuncExtern() IFuncExternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExternContext)
}

func (s *TopContext) TypeDefine() ITypeDefineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefineContext)
}

func (s *TopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitTop(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) Top() (localctx ITopContext) {
	localctx = NewTopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LjlParserRULE_top)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LjlParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.FuncDefine()
		}


	case LjlParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.FuncExtern()
		}
		{
			p.SetState(38)
			p.Match(LjlParserT__0)
		}


	case LjlParserT__1:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.TypeDefine()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ITypeDefineContext is an interface to support dynamic dispatch.
type ITypeDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefineContext differentiates from other interfaces.
	IsTypeDefineContext()
}

type TypeDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefineContext() *TypeDefineContext {
	var p = new(TypeDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_typeDefine
	return p
}

func (*TypeDefineContext) IsTypeDefineContext() {}

func NewTypeDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefineContext {
	var p = new(TypeDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_typeDefine

	return p
}

func (s *TypeDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefineContext) CopyFrom(ctx *TypeDefineContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type StructDefineContext struct {
	*TypeDefineContext
}

func NewStructDefineContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDefineContext {
	var p = new(StructDefineContext)

	p.TypeDefineContext = NewEmptyTypeDefineContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeDefineContext))

	return p
}

func (s *StructDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefineContext) ID() antlr.TerminalNode {
	return s.GetToken(LjlParserID, 0)
}

func (s *StructDefineContext) AllStructMem() []IStructMemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructMemContext)(nil)).Elem())
	var tst = make([]IStructMemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructMemContext)
		}
	}

	return tst
}

func (s *StructDefineContext) StructMem(i int) IStructMemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructMemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructMemContext)
}


func (s *StructDefineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitStructDefine(s)

	default:
		return t.VisitChildren(s)
	}
}


type TypeAliasContext struct {
	*TypeDefineContext
}

func NewTypeAliasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeAliasContext {
	var p = new(TypeAliasContext)

	p.TypeDefineContext = NewEmptyTypeDefineContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeDefineContext))

	return p
}

func (s *TypeAliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAliasContext) ID() antlr.TerminalNode {
	return s.GetToken(LjlParserID, 0)
}

func (s *TypeAliasContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}


func (s *TypeAliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitTypeAlias(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *LjlParser) TypeDefine() (localctx ITypeDefineContext) {
	localctx = NewTypeDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LjlParserRULE_typeDefine)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructDefineContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.Match(LjlParserT__1)
		}
		{
			p.SetState(44)
			p.Match(LjlParserID)
		}
		{
			p.SetState(45)
			p.Match(LjlParserT__2)
		}
		{
			p.SetState(46)
			p.Match(LjlParserT__3)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == LjlParserID {
			{
				p.SetState(47)
				p.StructMem()
			}


			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(53)
			p.Match(LjlParserT__4)
		}


	case 2:
		localctx = NewTypeAliasContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(LjlParserT__1)
		}
		{
			p.SetState(55)
			p.Match(LjlParserID)
		}
		{
			p.SetState(56)
			p.Match(LjlParserT__5)
		}
		{
			p.SetState(57)
			p.Typ()
		}

	}


	return localctx
}


// IStructMemContext is an interface to support dynamic dispatch.
type IStructMemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructMemContext differentiates from other interfaces.
	IsStructMemContext()
}

type StructMemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructMemContext() *StructMemContext {
	var p = new(StructMemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_structMem
	return p
}

func (*StructMemContext) IsStructMemContext() {}

func NewStructMemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemContext {
	var p = new(StructMemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_structMem

	return p
}

func (s *StructMemContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemContext) ID() antlr.TerminalNode {
	return s.GetToken(LjlParserID, 0)
}

func (s *StructMemContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *StructMemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StructMemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitStructMem(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) StructMem() (localctx IStructMemContext) {
	localctx = NewStructMemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LjlParserRULE_structMem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(LjlParserID)
	}
	{
		p.SetState(61)
		p.Match(LjlParserT__6)
	}
	{
		p.SetState(62)
		p.Typ()
	}



	return localctx
}


// IFuncExternContext is an interface to support dynamic dispatch.
type IFuncExternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExternContext differentiates from other interfaces.
	IsFuncExternContext()
}

type FuncExternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExternContext() *FuncExternContext {
	var p = new(FuncExternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_funcExtern
	return p
}

func (*FuncExternContext) IsFuncExternContext() {}

func NewFuncExternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExternContext {
	var p = new(FuncExternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_funcExtern

	return p
}

func (s *FuncExternContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExternContext) FuncProto() IFuncProtoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncProtoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncProtoContext)
}

func (s *FuncExternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FuncExternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitFuncExtern(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) FuncExtern() (localctx IFuncExternContext) {
	localctx = NewFuncExternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LjlParserRULE_funcExtern)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(LjlParserT__7)
	}
	{
		p.SetState(65)
		p.FuncProto()
	}



	return localctx
}


// IFuncDefineContext is an interface to support dynamic dispatch.
type IFuncDefineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDefineContext differentiates from other interfaces.
	IsFuncDefineContext()
}

type FuncDefineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefineContext() *FuncDefineContext {
	var p = new(FuncDefineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_funcDefine
	return p
}

func (*FuncDefineContext) IsFuncDefineContext() {}

func NewFuncDefineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefineContext {
	var p = new(FuncDefineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_funcDefine

	return p
}

func (s *FuncDefineContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefineContext) FuncProto() IFuncProtoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncProtoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncProtoContext)
}

func (s *FuncDefineContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDefineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FuncDefineContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitFuncDefine(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) FuncDefine() (localctx IFuncDefineContext) {
	localctx = NewFuncDefineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LjlParserRULE_funcDefine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.FuncProto()
	}
	{
		p.SetState(68)
		p.Block()
	}



	return localctx
}


// IFuncProtoContext is an interface to support dynamic dispatch.
type IFuncProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncProtoContext differentiates from other interfaces.
	IsFuncProtoContext()
}

type FuncProtoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncProtoContext() *FuncProtoContext {
	var p = new(FuncProtoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_funcProto
	return p
}

func (*FuncProtoContext) IsFuncProtoContext() {}

func NewFuncProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncProtoContext {
	var p = new(FuncProtoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_funcProto

	return p
}

func (s *FuncProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncProtoContext) ID() antlr.TerminalNode {
	return s.GetToken(LjlParserID, 0)
}

func (s *FuncProtoContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FuncProtoContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *FuncProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FuncProtoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitFuncProto(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) FuncProto() (localctx IFuncProtoContext) {
	localctx = NewFuncProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LjlParserRULE_funcProto)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(LjlParserT__8)
	}
	{
		p.SetState(71)
		p.Match(LjlParserID)
	}
	{
		p.SetState(72)
		p.Match(LjlParserT__9)
	}
	{
		p.SetState(73)
		p.ParamList()
	}
	{
		p.SetState(74)
		p.Match(LjlParserT__10)
	}
	{
		p.SetState(75)
		p.Match(LjlParserT__11)
	}
	{
		p.SetState(76)
		p.Typ()
	}



	return localctx
}


// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *ParamListContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParamListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitParamList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) ParamList() (localctx IParamListContext) {
	localctx = NewParamListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LjlParserRULE_paramList)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LjlParserT__10:
		p.EnterOuterAlt(localctx, 1)


	case LjlParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Param()
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == LjlParserT__12 {
			{
				p.SetState(80)
				p.Match(LjlParserT__12)
			}
			{
				p.SetState(81)
				p.Param()
			}


			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(LjlParserID, 0)
}

func (s *ParamContext) Typ() ITypContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LjlParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(LjlParserID)
	}
	{
		p.SetState(90)
		p.Match(LjlParserT__6)
	}
	{
		p.SetState(91)
		p.Typ()
	}



	return localctx
}


// ITypContext is an interface to support dynamic dispatch.
type ITypContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypContext differentiates from other interfaces.
	IsTypContext()
}

type TypContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypContext() *TypContext {
	var p = new(TypContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_typ
	return p
}

func (*TypContext) IsTypContext() {}

func NewTypContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypContext {
	var p = new(TypContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_typ

	return p
}

func (s *TypContext) GetParser() antlr.Parser { return s.parser }

func (s *TypContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(LjlParserID)
}

func (s *TypContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(LjlParserID, i)
}

func (s *TypContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TypContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitTyp(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) Typ() (localctx ITypContext) {
	localctx = NewTypContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LjlParserRULE_typ)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(LjlParserID)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(LjlParserT__9)
		}
		{
			p.SetState(95)
			p.Match(LjlParserID)
		}
		{
			p.SetState(96)
			p.Match(LjlParserT__12)
		}
		{
			p.SetState(97)
			p.Match(LjlParserID)
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == LjlParserT__12 {
			{
				p.SetState(98)
				p.Match(LjlParserT__12)
			}
			{
				p.SetState(99)
				p.Match(LjlParserID)
			}


			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(105)
			p.Match(LjlParserT__10)
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
			p.Match(LjlParserT__9)
		}
		{
			p.SetState(107)
			p.Match(LjlParserT__10)
		}

	}


	return localctx
}


// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BlockContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LjlParserRULE_block)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(LjlParserT__3)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ((((_la - 4)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 4))) & ((1 << (LjlParserT__3 - 4)) | (1 << (LjlParserT__9 - 4)) | (1 << (LjlParserT__23 - 4)) | (1 << (LjlParserT__26 - 4)) | (1 << (LjlParserINT - 4)) | (1 << (LjlParserFLOAT - 4)) | (1 << (LjlParserSTRING - 4)) | (1 << (LjlParserID - 4)))) != 0) {
		{
			p.SetState(111)
			p.expr(0)
		}


		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(LjlParserT__4)
	}



	return localctx
}


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type AddSubExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *AddSubExprContext) GetOp() antlr.Token { return s.op }


func (s *AddSubExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *AddSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitAddSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type MemberExprContext struct {
	*ExprContext
}

func NewMemberExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberExprContext {
	var p = new(MemberExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MemberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MemberExprContext) ID() antlr.TerminalNode {
	return s.GetToken(LjlParserID, 0)
}


func (s *MemberExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitMemberExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type OrExprContext struct {
	*ExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type CmpExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewCmpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmpExprContext {
	var p = new(CmpExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *CmpExprContext) GetOp() antlr.Token { return s.op }


func (s *CmpExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *CmpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CmpExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *CmpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitCmpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type EqExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewEqExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqExprContext {
	var p = new(EqExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *EqExprContext) GetOp() antlr.Token { return s.op }


func (s *EqExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *EqExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitEqExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type NotExprContext struct {
	*ExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type ValRefExprContext struct {
	*ExprContext
}

func NewValRefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValRefExprContext {
	var p = new(ValRefExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ValRefExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValRefExprContext) ID() antlr.TerminalNode {
	return s.GetToken(LjlParserID, 0)
}


func (s *ValRefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitValRefExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type IfExprContext struct {
	*ExprContext
}

func NewIfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfExprContext {
	var p = new(IfExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IfExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitIfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type BlockExprContext struct {
	*ExprContext
}

func NewBlockExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockExprContext {
	var p = new(BlockExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BlockExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockExprContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}


func (s *BlockExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitBlockExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type CallExprContext struct {
	*ExprContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) ID() antlr.TerminalNode {
	return s.GetToken(LjlParserID, 0)
}

func (s *CallExprContext) TupleVal() ITupleValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITupleValContext)
}


func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type TupleExprContext struct {
	*ExprContext
}

func NewTupleExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleExprContext {
	var p = new(TupleExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TupleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleExprContext) TupleVal() ITupleValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITupleValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITupleValContext)
}


func (s *TupleExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitTupleExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type LitExprContext struct {
	*ExprContext
}

func NewLitExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitExprContext {
	var p = new(LitExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LitExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitExprContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}


func (s *LitExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitLitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type MulDivExprContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivExprContext {
	var p = new(MulDivExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}


func (s *MulDivExprContext) GetOp() antlr.Token { return s.op }


func (s *MulDivExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulDivExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *MulDivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitMulDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type AndExprContext struct {
	*ExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}


func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *LjlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *LjlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, LjlParserRULE_expr, _p)
	var _la int


	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBlockExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(120)
			p.Block()
		}


	case 2:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(121)
			p.Match(LjlParserT__23)
		}
		{
			p.SetState(122)
			p.expr(9)
		}


	case 3:
		localctx = NewCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(123)
			p.Match(LjlParserID)
		}
		{
			p.SetState(124)
			p.TupleVal()
		}


	case 4:
		localctx = NewValRefExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(125)
			p.Match(LjlParserID)
		}


	case 5:
		localctx = NewLitExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(126)
			p.Literal()
		}


	case 6:
		localctx = NewIfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(127)
			p.Match(LjlParserT__26)
		}
		{
			p.SetState(128)
			p.expr(0)
		}
		{
			p.SetState(129)
			p.Match(LjlParserT__27)
		}
		{
			p.SetState(130)
			p.expr(0)
		}
		{
			p.SetState(131)
			p.Match(LjlParserT__28)
		}
		{
			p.SetState(132)
			p.expr(3)
		}


	case 7:
		localctx = NewTupleExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(134)
			p.TupleVal()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LjlParserRULE_expr)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(138)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LjlParserT__13 || _la == LjlParserT__14) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(139)
					p.expr(14)
				}


			case 2:
				localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LjlParserRULE_expr)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(141)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LjlParserT__15 || _la == LjlParserT__16) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(142)
					p.expr(13)
				}


			case 3:
				localctx = NewCmpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LjlParserRULE_expr)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(144)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CmpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << LjlParserT__17) | (1 << LjlParserT__18) | (1 << LjlParserT__19) | (1 << LjlParserT__20))) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CmpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(145)
					p.expr(12)
				}


			case 4:
				localctx = NewEqExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LjlParserRULE_expr)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(147)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LjlParserT__21 || _la == LjlParserT__22) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(148)
					p.expr(11)
				}


			case 5:
				localctx = NewAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LjlParserRULE_expr)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(150)
					p.Match(LjlParserT__24)
				}
				{
					p.SetState(151)
					p.expr(9)
				}


			case 6:
				localctx = NewOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LjlParserRULE_expr)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(153)
					p.Match(LjlParserT__25)
				}
				{
					p.SetState(154)
					p.expr(8)
				}


			case 7:
				localctx = NewMemberExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LjlParserRULE_expr)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(156)
					p.Match(LjlParserT__29)
				}
				{
					p.SetState(157)
					p.Match(LjlParserID)
				}

			}

		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}



	return localctx
}


// ITupleValContext is an interface to support dynamic dispatch.
type ITupleValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTupleValContext differentiates from other interfaces.
	IsTupleValContext()
}

type TupleValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTupleValContext() *TupleValContext {
	var p = new(TupleValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_tupleVal
	return p
}

func (*TupleValContext) IsTupleValContext() {}

func NewTupleValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleValContext {
	var p = new(TupleValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_tupleVal

	return p
}

func (s *TupleValContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleValContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TupleValContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TupleValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TupleValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitTupleVal(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) TupleVal() (localctx ITupleValContext) {
	localctx = NewTupleValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LjlParserRULE_tupleVal)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(LjlParserT__9)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 4)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 4))) & ((1 << (LjlParserT__3 - 4)) | (1 << (LjlParserT__9 - 4)) | (1 << (LjlParserT__23 - 4)) | (1 << (LjlParserT__26 - 4)) | (1 << (LjlParserINT - 4)) | (1 << (LjlParserFLOAT - 4)) | (1 << (LjlParserSTRING - 4)) | (1 << (LjlParserID - 4)))) != 0) {
		{
			p.SetState(164)
			p.expr(0)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == LjlParserT__12 {
			{
				p.SetState(165)
				p.Match(LjlParserT__12)
			}
			{
				p.SetState(166)
				p.expr(0)
			}


			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(174)
		p.Match(LjlParserT__10)
	}



	return localctx
}


// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LjlParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LjlParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INT() antlr.TerminalNode {
	return s.GetToken(LjlParserINT, 0)
}

func (s *LiteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(LjlParserFLOAT, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(LjlParserSTRING, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LjlVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *LjlParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LjlParserRULE_literal)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 31)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 31))) & ((1 << (LjlParserINT - 31)) | (1 << (LjlParserFLOAT - 31)) | (1 << (LjlParserSTRING - 31)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


func (p *LjlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LjlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

