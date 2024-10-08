// Code generated from E:/02_Resource/01_Code/jc_go_manager/go-rules/parser/Rules.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Rules

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RulesParser struct {
	*antlr.BaseParser
}

var RulesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rulesParserInit() {
	staticData := &RulesParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'['", "']'", "'&&'", "'||'", "'=='", "'!='", "'<'", "'<='",
		"'>'", "'>='", "'in'", "'+'", "'-'", "'*'", "'/'", "'%'", "'contains'",
		"'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "AND", "OR", "EQ", "NEQ", "LT", "LTE", "GT", "GTE",
		"IN", "ADD", "SUB", "MUL", "DIV", "MOD", "CONTAINS", "LPARENT", "RPARENT",
		"BooleanLiteral", "NumberLiteral", "DurationLiteral", "StringLiteral",
		"Identifier", "WS",
	}
	staticData.RuleNames = []string{
		"rules", "argumentExpressionList", "numberList", "stringList", "valueList",
		"primaryExpression", "booleanExpression", "expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 120, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 24, 8, 1, 10, 1, 12, 1, 27, 9, 1, 3, 1, 29, 8, 1, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 2, 5, 2, 36, 8, 2, 10, 2, 12, 2, 39, 9, 2, 1, 3, 1,
		3, 1, 3, 5, 3, 44, 8, 3, 10, 3, 12, 3, 47, 9, 3, 1, 4, 1, 4, 4, 4, 51,
		8, 4, 11, 4, 12, 4, 52, 1, 4, 4, 4, 56, 8, 4, 11, 4, 12, 4, 57, 3, 4, 60,
		8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 82, 8, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 90, 8, 5, 10, 5, 12, 5, 93, 9, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 104, 8, 6,
		10, 6, 12, 6, 107, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 115,
		8, 7, 10, 7, 12, 7, 118, 9, 7, 1, 7, 0, 3, 10, 12, 14, 8, 0, 2, 4, 6, 8,
		10, 12, 14, 0, 6, 2, 0, 22, 22, 24, 24, 1, 0, 15, 17, 1, 0, 13, 14, 1,
		0, 6, 7, 1, 0, 8, 11, 1, 0, 4, 5, 131, 0, 16, 1, 0, 0, 0, 2, 19, 1, 0,
		0, 0, 4, 32, 1, 0, 0, 0, 6, 40, 1, 0, 0, 0, 8, 48, 1, 0, 0, 0, 10, 81,
		1, 0, 0, 0, 12, 94, 1, 0, 0, 0, 14, 108, 1, 0, 0, 0, 16, 17, 3, 14, 7,
		0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0, 0, 19, 28, 5, 19, 0, 0, 20, 25,
		3, 14, 7, 0, 21, 22, 5, 1, 0, 0, 22, 24, 3, 14, 7, 0, 23, 21, 1, 0, 0,
		0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 29,
		1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 20, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0,
		29, 30, 1, 0, 0, 0, 30, 31, 5, 20, 0, 0, 31, 3, 1, 0, 0, 0, 32, 37, 5,
		22, 0, 0, 33, 34, 5, 1, 0, 0, 34, 36, 5, 22, 0, 0, 35, 33, 1, 0, 0, 0,
		36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 5, 1, 0,
		0, 0, 39, 37, 1, 0, 0, 0, 40, 45, 5, 24, 0, 0, 41, 42, 5, 1, 0, 0, 42,
		44, 5, 24, 0, 0, 43, 41, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0,
		0, 0, 45, 46, 1, 0, 0, 0, 46, 7, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 59,
		5, 2, 0, 0, 49, 51, 3, 4, 2, 0, 50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0,
		52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 60, 1, 0, 0, 0, 54, 56, 3,
		6, 3, 0, 55, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57,
		58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 50, 1, 0, 0, 0, 59, 55, 1, 0, 0,
		0, 60, 61, 1, 0, 0, 0, 61, 62, 5, 3, 0, 0, 62, 9, 1, 0, 0, 0, 63, 64, 6,
		5, -1, 0, 64, 82, 5, 25, 0, 0, 65, 82, 5, 21, 0, 0, 66, 82, 5, 22, 0, 0,
		67, 82, 5, 24, 0, 0, 68, 82, 5, 23, 0, 0, 69, 70, 5, 25, 0, 0, 70, 71,
		5, 18, 0, 0, 71, 82, 7, 0, 0, 0, 72, 73, 5, 25, 0, 0, 73, 74, 5, 12, 0,
		0, 74, 82, 3, 8, 4, 0, 75, 76, 5, 25, 0, 0, 76, 82, 3, 2, 1, 0, 77, 78,
		5, 19, 0, 0, 78, 79, 3, 14, 7, 0, 79, 80, 5, 20, 0, 0, 80, 82, 1, 0, 0,
		0, 81, 63, 1, 0, 0, 0, 81, 65, 1, 0, 0, 0, 81, 66, 1, 0, 0, 0, 81, 67,
		1, 0, 0, 0, 81, 68, 1, 0, 0, 0, 81, 69, 1, 0, 0, 0, 81, 72, 1, 0, 0, 0,
		81, 75, 1, 0, 0, 0, 81, 77, 1, 0, 0, 0, 82, 91, 1, 0, 0, 0, 83, 84, 10,
		2, 0, 0, 84, 85, 7, 1, 0, 0, 85, 90, 3, 10, 5, 3, 86, 87, 10, 1, 0, 0,
		87, 88, 7, 2, 0, 0, 88, 90, 3, 10, 5, 2, 89, 83, 1, 0, 0, 0, 89, 86, 1,
		0, 0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		11, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 95, 6, 6, -1, 0, 95, 96, 3, 10,
		5, 0, 96, 105, 1, 0, 0, 0, 97, 98, 10, 2, 0, 0, 98, 99, 7, 3, 0, 0, 99,
		104, 3, 10, 5, 0, 100, 101, 10, 1, 0, 0, 101, 102, 7, 4, 0, 0, 102, 104,
		3, 10, 5, 0, 103, 97, 1, 0, 0, 0, 103, 100, 1, 0, 0, 0, 104, 107, 1, 0,
		0, 0, 105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 13, 1, 0, 0, 0,
		107, 105, 1, 0, 0, 0, 108, 109, 6, 7, -1, 0, 109, 110, 3, 12, 6, 0, 110,
		116, 1, 0, 0, 0, 111, 112, 10, 1, 0, 0, 112, 113, 7, 5, 0, 0, 113, 115,
		3, 12, 6, 0, 114, 111, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0,
		0, 0, 116, 117, 1, 0, 0, 0, 117, 15, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0,
		13, 25, 28, 37, 45, 52, 57, 59, 81, 89, 91, 103, 105, 116,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RulesParserInit initializes any static state used to implement RulesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRulesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RulesParserInit() {
	staticData := &RulesParserStaticData
	staticData.once.Do(rulesParserInit)
}

// NewRulesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRulesParser(input antlr.TokenStream) *RulesParser {
	RulesParserInit()
	this := new(RulesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RulesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Rules.g4"

	return this
}

// RulesParser tokens.
const (
	RulesParserEOF             = antlr.TokenEOF
	RulesParserT__0            = 1
	RulesParserT__1            = 2
	RulesParserT__2            = 3
	RulesParserAND             = 4
	RulesParserOR              = 5
	RulesParserEQ              = 6
	RulesParserNEQ             = 7
	RulesParserLT              = 8
	RulesParserLTE             = 9
	RulesParserGT              = 10
	RulesParserGTE             = 11
	RulesParserIN              = 12
	RulesParserADD             = 13
	RulesParserSUB             = 14
	RulesParserMUL             = 15
	RulesParserDIV             = 16
	RulesParserMOD             = 17
	RulesParserCONTAINS        = 18
	RulesParserLPARENT         = 19
	RulesParserRPARENT         = 20
	RulesParserBooleanLiteral  = 21
	RulesParserNumberLiteral   = 22
	RulesParserDurationLiteral = 23
	RulesParserStringLiteral   = 24
	RulesParserIdentifier      = 25
	RulesParserWS              = 26
)

// RulesParser rules.
const (
	RulesParserRULE_rules                  = 0
	RulesParserRULE_argumentExpressionList = 1
	RulesParserRULE_numberList             = 2
	RulesParserRULE_stringList             = 3
	RulesParserRULE_valueList              = 4
	RulesParserRULE_primaryExpression      = 5
	RulesParserRULE_booleanExpression      = 6
	RulesParserRULE_expression             = 7
)

// IRulesContext is an interface to support dynamic dispatch.
type IRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode

	// IsRulesContext differentiates from other interfaces.
	IsRulesContext()
}

type RulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesContext() *RulesContext {
	var p = new(RulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_rules
	return p
}

func InitEmptyRulesContext(p *RulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_rules
}

func (*RulesContext) IsRulesContext() {}

func NewRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesContext {
	var p = new(RulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_rules

	return p
}

func (s *RulesContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RulesContext) EOF() antlr.TerminalNode {
	return s.GetToken(RulesParserEOF, 0)
}

func (s *RulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterRules(s)
	}
}

func (s *RulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitRules(s)
	}
}

func (s *RulesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitRules(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RulesParser) Rules() (localctx IRulesContext) {
	localctx = NewRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RulesParserRULE_rules)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.expression(0)
	}
	{
		p.SetState(17)
		p.Match(RulesParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentExpressionListContext is an interface to support dynamic dispatch.
type IArgumentExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPARENT() antlr.TerminalNode
	RPARENT() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsArgumentExpressionListContext differentiates from other interfaces.
	IsArgumentExpressionListContext()
}

type ArgumentExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentExpressionListContext() *ArgumentExpressionListContext {
	var p = new(ArgumentExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_argumentExpressionList
	return p
}

func InitEmptyArgumentExpressionListContext(p *ArgumentExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_argumentExpressionList
}

func (*ArgumentExpressionListContext) IsArgumentExpressionListContext() {}

func NewArgumentExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentExpressionListContext {
	var p = new(ArgumentExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_argumentExpressionList

	return p
}

func (s *ArgumentExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentExpressionListContext) LPARENT() antlr.TerminalNode {
	return s.GetToken(RulesParserLPARENT, 0)
}

func (s *ArgumentExpressionListContext) RPARENT() antlr.TerminalNode {
	return s.GetToken(RulesParserRPARENT, 0)
}

func (s *ArgumentExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterArgumentExpressionList(s)
	}
}

func (s *ArgumentExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitArgumentExpressionList(s)
	}
}

func (s *ArgumentExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitArgumentExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RulesParser) ArgumentExpressionList() (localctx IArgumentExpressionListContext) {
	localctx = NewArgumentExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RulesParserRULE_argumentExpressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(RulesParserLPARENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&65536000) != 0 {
		{
			p.SetState(20)
			p.expression(0)
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == RulesParserT__0 {
			{
				p.SetState(21)
				p.Match(RulesParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(22)
				p.expression(0)
			}

			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(30)
		p.Match(RulesParserRPARENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberListContext is an interface to support dynamic dispatch.
type INumberListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumberLiteral() []antlr.TerminalNode
	NumberLiteral(i int) antlr.TerminalNode

	// IsNumberListContext differentiates from other interfaces.
	IsNumberListContext()
}

type NumberListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberListContext() *NumberListContext {
	var p = new(NumberListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_numberList
	return p
}

func InitEmptyNumberListContext(p *NumberListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_numberList
}

func (*NumberListContext) IsNumberListContext() {}

func NewNumberListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberListContext {
	var p = new(NumberListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_numberList

	return p
}

func (s *NumberListContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberListContext) AllNumberLiteral() []antlr.TerminalNode {
	return s.GetTokens(RulesParserNumberLiteral)
}

func (s *NumberListContext) NumberLiteral(i int) antlr.TerminalNode {
	return s.GetToken(RulesParserNumberLiteral, i)
}

func (s *NumberListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterNumberList(s)
	}
}

func (s *NumberListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitNumberList(s)
	}
}

func (s *NumberListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitNumberList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RulesParser) NumberList() (localctx INumberListContext) {
	localctx = NewNumberListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RulesParserRULE_numberList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(RulesParserNumberLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RulesParserT__0 {
		{
			p.SetState(33)
			p.Match(RulesParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Match(RulesParserNumberLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringListContext is an interface to support dynamic dispatch.
type IStringListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStringLiteral() []antlr.TerminalNode
	StringLiteral(i int) antlr.TerminalNode

	// IsStringListContext differentiates from other interfaces.
	IsStringListContext()
}

type StringListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringListContext() *StringListContext {
	var p = new(StringListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_stringList
	return p
}

func InitEmptyStringListContext(p *StringListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_stringList
}

func (*StringListContext) IsStringListContext() {}

func NewStringListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringListContext {
	var p = new(StringListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_stringList

	return p
}

func (s *StringListContext) GetParser() antlr.Parser { return s.parser }

func (s *StringListContext) AllStringLiteral() []antlr.TerminalNode {
	return s.GetTokens(RulesParserStringLiteral)
}

func (s *StringListContext) StringLiteral(i int) antlr.TerminalNode {
	return s.GetToken(RulesParserStringLiteral, i)
}

func (s *StringListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterStringList(s)
	}
}

func (s *StringListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitStringList(s)
	}
}

func (s *StringListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitStringList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RulesParser) StringList() (localctx IStringListContext) {
	localctx = NewStringListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RulesParserRULE_stringList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(RulesParserStringLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == RulesParserT__0 {
		{
			p.SetState(41)
			p.Match(RulesParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(42)
			p.Match(RulesParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueListContext is an interface to support dynamic dispatch.
type IValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumberList() []INumberListContext
	NumberList(i int) INumberListContext
	AllStringList() []IStringListContext
	StringList(i int) IStringListContext

	// IsValueListContext differentiates from other interfaces.
	IsValueListContext()
}

type ValueListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueListContext() *ValueListContext {
	var p = new(ValueListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_valueList
	return p
}

func InitEmptyValueListContext(p *ValueListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_valueList
}

func (*ValueListContext) IsValueListContext() {}

func NewValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueListContext {
	var p = new(ValueListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_valueList

	return p
}

func (s *ValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueListContext) AllNumberList() []INumberListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberListContext); ok {
			len++
		}
	}

	tst := make([]INumberListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberListContext); ok {
			tst[i] = t.(INumberListContext)
			i++
		}
	}

	return tst
}

func (s *ValueListContext) NumberList(i int) INumberListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberListContext)
}

func (s *ValueListContext) AllStringList() []IStringListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStringListContext); ok {
			len++
		}
	}

	tst := make([]IStringListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStringListContext); ok {
			tst[i] = t.(IStringListContext)
			i++
		}
	}

	return tst
}

func (s *ValueListContext) StringList(i int) IStringListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringListContext)
}

func (s *ValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterValueList(s)
	}
}

func (s *ValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitValueList(s)
	}
}

func (s *ValueListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitValueList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RulesParser) ValueList() (localctx IValueListContext) {
	localctx = NewValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RulesParserRULE_valueList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(RulesParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RulesParserNumberLiteral:
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == RulesParserNumberLiteral {
			{
				p.SetState(49)
				p.NumberList()
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case RulesParserStringLiteral:
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == RulesParserStringLiteral {
			{
				p.SetState(54)
				p.StringList()
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(61)
		p.Match(RulesParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CopyAll(ctx *PrimaryExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiplicativeExprContext struct {
	PrimaryExpressionContext
	op antlr.Token
}

func NewMultiplicativeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *MultiplicativeExprContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicativeExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExprContext) AllPrimaryExpression() []IPrimaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryExpressionContext); ok {
			tst[i] = t.(IPrimaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExprContext) PrimaryExpression(i int) IPrimaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *MultiplicativeExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(RulesParserMUL, 0)
}

func (s *MultiplicativeExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(RulesParserDIV, 0)
}

func (s *MultiplicativeExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(RulesParserMOD, 0)
}

func (s *MultiplicativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitMultiplicativeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentContext struct {
	PrimaryExpressionContext
}

func NewIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentContext {
	var p = new(IdentContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RulesParserIdentifier, 0)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContainsCondContext struct {
	PrimaryExpressionContext
}

func NewContainsCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContainsCondContext {
	var p = new(ContainsCondContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *ContainsCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainsCondContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RulesParserIdentifier, 0)
}

func (s *ContainsCondContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(RulesParserCONTAINS, 0)
}

func (s *ContainsCondContext) NumberLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserNumberLiteral, 0)
}

func (s *ContainsCondContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserStringLiteral, 0)
}

func (s *ContainsCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterContainsCond(s)
	}
}

func (s *ContainsCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitContainsCond(s)
	}
}

func (s *ContainsCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitContainsCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type AdditiveExprContext struct {
	PrimaryExpressionContext
	op antlr.Token
}

func NewAdditiveExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *AdditiveExprContext) GetOp() antlr.Token { return s.op }

func (s *AdditiveExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) AllPrimaryExpression() []IPrimaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IPrimaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimaryExpressionContext); ok {
			tst[i] = t.(IPrimaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExprContext) PrimaryExpression(i int) IPrimaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *AdditiveExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(RulesParserADD, 0)
}

func (s *AdditiveExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(RulesParserSUB, 0)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitAdditiveExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberLitContext struct {
	PrimaryExpressionContext
}

func NewNumberLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLitContext {
	var p = new(NumberLitContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *NumberLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLitContext) NumberLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserNumberLiteral, 0)
}

func (s *NumberLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterNumberLit(s)
	}
}

func (s *NumberLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitNumberLit(s)
	}
}

func (s *NumberLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitNumberLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type InCondContext struct {
	PrimaryExpressionContext
}

func NewInCondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InCondContext {
	var p = new(InCondContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *InCondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InCondContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RulesParserIdentifier, 0)
}

func (s *InCondContext) IN() antlr.TerminalNode {
	return s.GetToken(RulesParserIN, 0)
}

func (s *InCondContext) ValueList() IValueListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueListContext)
}

func (s *InCondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterInCond(s)
	}
}

func (s *InCondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitInCond(s)
	}
}

func (s *InCondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitInCond(s)

	default:
		return t.VisitChildren(s)
	}
}

type DurationLitContext struct {
	PrimaryExpressionContext
}

func NewDurationLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DurationLitContext {
	var p = new(DurationLitContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *DurationLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationLitContext) DurationLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserDurationLiteral, 0)
}

func (s *DurationLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterDurationLit(s)
	}
}

func (s *DurationLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitDurationLit(s)
	}
}

func (s *DurationLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitDurationLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParentExprContext struct {
	PrimaryExpressionContext
}

func NewParentExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentExprContext {
	var p = new(ParentExprContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *ParentExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentExprContext) LPARENT() antlr.TerminalNode {
	return s.GetToken(RulesParserLPARENT, 0)
}

func (s *ParentExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParentExprContext) RPARENT() antlr.TerminalNode {
	return s.GetToken(RulesParserRPARENT, 0)
}

func (s *ParentExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterParentExpr(s)
	}
}

func (s *ParentExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitParentExpr(s)
	}
}

func (s *ParentExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitParentExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncExprContext struct {
	PrimaryExpressionContext
}

func NewFuncExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncExprContext {
	var p = new(FuncExprContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *FuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(RulesParserIdentifier, 0)
}

func (s *FuncExprContext) ArgumentExpressionList() IArgumentExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentExpressionListContext)
}

func (s *FuncExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterFuncExpr(s)
	}
}

func (s *FuncExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitFuncExpr(s)
	}
}

func (s *FuncExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitFuncExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolLitContext struct {
	PrimaryExpressionContext
}

func NewBoolLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLitContext {
	var p = new(BoolLitContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserBooleanLiteral, 0)
}

func (s *BoolLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterBoolLit(s)
	}
}

func (s *BoolLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitBoolLit(s)
	}
}

func (s *BoolLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitBoolLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringLitContext struct {
	PrimaryExpressionContext
}

func NewStringLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLitContext {
	var p = new(StringLitContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *StringLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLitContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(RulesParserStringLiteral, 0)
}

func (s *StringLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterStringLit(s)
	}
}

func (s *StringLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitStringLit(s)
	}
}

func (s *StringLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitStringLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RulesParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *RulesParser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, RulesParserRULE_primaryExpression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIdentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(64)
			p.Match(RulesParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewBoolLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Match(RulesParserBooleanLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewNumberLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(RulesParserNumberLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStringLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)
			p.Match(RulesParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewDurationLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(RulesParserDurationLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewContainsCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(RulesParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Match(RulesParserCONTAINS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RulesParserNumberLiteral || _la == RulesParserStringLiteral) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 7:
		localctx = NewInCondContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(72)
			p.Match(RulesParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(RulesParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(74)
			p.ValueList()
		}

	case 8:
		localctx = NewFuncExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(75)
			p.Match(RulesParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(76)
			p.ArgumentExpressionList()
		}

	case 9:
		localctx = NewParentExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(77)
			p.Match(RulesParserLPARENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.expression(0)
		}
		{
			p.SetState(79)
			p.Match(RulesParserRPARENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExprContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_primaryExpression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(84)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicativeExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&229376) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicativeExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(85)
					p.primaryExpression(3)
				}

			case 2:
				localctx = NewAdditiveExprContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_primaryExpression)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(87)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditiveExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RulesParserADD || _la == RulesParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditiveExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(88)
					p.primaryExpression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanExpressionContext is an interface to support dynamic dispatch.
type IBooleanExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBooleanExpressionContext differentiates from other interfaces.
	IsBooleanExpressionContext()
}

type BooleanExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanExpressionContext() *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_booleanExpression
	return p
}

func InitEmptyBooleanExpressionContext(p *BooleanExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_booleanExpression
}

func (*BooleanExpressionContext) IsBooleanExpressionContext() {}

func NewBooleanExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_booleanExpression

	return p
}

func (s *BooleanExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanExpressionContext) CopyAll(ctx *BooleanExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqualityExprContext struct {
	BooleanExpressionContext
	op antlr.Token
}

func NewEqualityExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExprContext {
	var p = new(EqualityExprContext)

	InitEmptyBooleanExpressionContext(&p.BooleanExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*BooleanExpressionContext))

	return p
}

func (s *EqualityExprContext) GetOp() antlr.Token { return s.op }

func (s *EqualityExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) BooleanExpression() IBooleanExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *EqualityExprContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *EqualityExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(RulesParserEQ, 0)
}

func (s *EqualityExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(RulesParserNEQ, 0)
}

func (s *EqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterEqualityExpr(s)
	}
}

func (s *EqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitEqualityExpr(s)
	}
}

func (s *EqualityExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitEqualityExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExprContext struct {
	BooleanExpressionContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	InitEmptyBooleanExpressionContext(&p.BooleanExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*BooleanExpressionContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RelationalExprContext struct {
	BooleanExpressionContext
	op antlr.Token
}

func NewRelationalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExprContext {
	var p = new(RelationalExprContext)

	InitEmptyBooleanExpressionContext(&p.BooleanExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*BooleanExpressionContext))

	return p
}

func (s *RelationalExprContext) GetOp() antlr.Token { return s.op }

func (s *RelationalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) BooleanExpression() IBooleanExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *RelationalExprContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *RelationalExprContext) LT() antlr.TerminalNode {
	return s.GetToken(RulesParserLT, 0)
}

func (s *RelationalExprContext) LTE() antlr.TerminalNode {
	return s.GetToken(RulesParserLTE, 0)
}

func (s *RelationalExprContext) GT() antlr.TerminalNode {
	return s.GetToken(RulesParserGT, 0)
}

func (s *RelationalExprContext) GTE() antlr.TerminalNode {
	return s.GetToken(RulesParserGTE, 0)
}

func (s *RelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterRelationalExpr(s)
	}
}

func (s *RelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitRelationalExpr(s)
	}
}

func (s *RelationalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitRelationalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RulesParser) BooleanExpression() (localctx IBooleanExpressionContext) {
	return p.booleanExpression(0)
}

func (p *RulesParser) booleanExpression(_p int) (localctx IBooleanExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewBooleanExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, RulesParserRULE_booleanExpression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPrimaryExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(95)
		p.primaryExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEqualityExprContext(p, NewBooleanExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_booleanExpression)
				p.SetState(97)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(98)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqualityExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == RulesParserEQ || _la == RulesParserNEQ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqualityExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(99)
					p.primaryExpression(0)
				}

			case 2:
				localctx = NewRelationalExprContext(p, NewBooleanExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_booleanExpression)
				p.SetState(100)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(101)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RelationalExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3840) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RelationalExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(102)
					p.primaryExpression(0)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RulesParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RulesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LogicalExprContext struct {
	ExpressionContext
	op antlr.Token
}

func NewLogicalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExprContext {
	var p = new(LogicalExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExprContext) GetOp() antlr.Token { return s.op }

func (s *LogicalExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalExprContext) BooleanExpression() IBooleanExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *LogicalExprContext) AND() antlr.TerminalNode {
	return s.GetToken(RulesParserAND, 0)
}

func (s *LogicalExprContext) OR() antlr.TerminalNode {
	return s.GetToken(RulesParserOR, 0)
}

func (s *LogicalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterLogicalExpr(s)
	}
}

func (s *LogicalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitLogicalExpr(s)
	}
}

func (s *LogicalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitLogicalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanExprContext struct {
	ExpressionContext
}

func NewBooleanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanExprContext {
	var p = new(BooleanExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExprContext) BooleanExpression() IBooleanExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *BooleanExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.EnterBooleanExpr(s)
	}
}

func (s *BooleanExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RulesListener); ok {
		listenerT.ExitBooleanExpr(s)
	}
}

func (s *BooleanExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case RulesVisitor:
		return t.VisitBooleanExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *RulesParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *RulesParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, RulesParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewBooleanExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(109)
		p.booleanExpression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, RulesParserRULE_expression)
			p.SetState(111)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(112)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*LogicalExprContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == RulesParserAND || _la == RulesParserOR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*LogicalExprContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(113)
				p.booleanExpression(0)
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *RulesParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	case 6:
		var t *BooleanExpressionContext = nil
		if localctx != nil {
			t = localctx.(*BooleanExpressionContext)
		}
		return p.BooleanExpression_Sempred(t, predIndex)

	case 7:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RulesParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RulesParser) BooleanExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *RulesParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
