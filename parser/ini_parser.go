// Code generated from parser/Ini.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Ini

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 96, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 7, 2, 36,
	10, 2, 12, 2, 14, 2, 39, 11, 2, 5, 2, 41, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 52, 10, 5, 12, 5, 14, 5, 55, 11, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 61, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8,
	67, 10, 8, 3, 9, 3, 9, 3, 9, 6, 9, 72, 10, 9, 13, 9, 14, 9, 73, 3, 10,
	3, 10, 3, 10, 5, 10, 79, 10, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 5, 14, 90, 10, 14, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 2, 2, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	2, 2, 2, 90, 2, 40, 3, 2, 2, 2, 4, 44, 3, 2, 2, 2, 6, 47, 3, 2, 2, 2, 8,
	53, 3, 2, 2, 2, 10, 56, 3, 2, 2, 2, 12, 62, 3, 2, 2, 2, 14, 66, 3, 2, 2,
	2, 16, 68, 3, 2, 2, 2, 18, 78, 3, 2, 2, 2, 20, 80, 3, 2, 2, 2, 22, 83,
	3, 2, 2, 2, 24, 85, 3, 2, 2, 2, 26, 89, 3, 2, 2, 2, 28, 91, 3, 2, 2, 2,
	30, 93, 3, 2, 2, 2, 32, 37, 5, 4, 3, 2, 33, 34, 7, 14, 2, 2, 34, 36, 5,
	4, 3, 2, 35, 33, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37,
	38, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 32, 3, 2, 2,
	2, 40, 41, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 7, 2, 2, 3, 43, 3, 3,
	2, 2, 2, 44, 45, 5, 6, 4, 2, 45, 46, 5, 8, 5, 2, 46, 5, 3, 2, 2, 2, 47,
	48, 7, 5, 2, 2, 48, 7, 3, 2, 2, 2, 49, 52, 5, 10, 6, 2, 50, 52, 5, 20,
	11, 2, 51, 49, 3, 2, 2, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53,
	51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 9, 3, 2, 2, 2, 55, 53, 3, 2, 2,
	2, 56, 57, 7, 14, 2, 2, 57, 58, 5, 12, 7, 2, 58, 60, 7, 3, 2, 2, 59, 61,
	5, 14, 8, 2, 60, 59, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 11, 3, 2, 2, 2,
	62, 63, 7, 12, 2, 2, 63, 13, 3, 2, 2, 2, 64, 67, 5, 18, 10, 2, 65, 67,
	5, 16, 9, 2, 66, 64, 3, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67, 15, 3, 2, 2, 2,
	68, 71, 5, 18, 10, 2, 69, 70, 7, 4, 2, 2, 70, 72, 5, 18, 10, 2, 71, 69,
	3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2,
	74, 17, 3, 2, 2, 2, 75, 79, 5, 22, 12, 2, 76, 79, 5, 24, 13, 2, 77, 79,
	5, 26, 14, 2, 78, 75, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2, 2,
	2, 79, 19, 3, 2, 2, 2, 80, 81, 7, 14, 2, 2, 81, 82, 7, 13, 2, 2, 82, 21,
	3, 2, 2, 2, 83, 84, 7, 8, 2, 2, 84, 23, 3, 2, 2, 2, 85, 86, 7, 9, 2, 2,
	86, 25, 3, 2, 2, 2, 87, 90, 5, 28, 15, 2, 88, 90, 5, 30, 16, 2, 89, 87,
	3, 2, 2, 2, 89, 88, 3, 2, 2, 2, 90, 27, 3, 2, 2, 2, 91, 92, 7, 10, 2, 2,
	92, 29, 3, 2, 2, 2, 93, 94, 7, 11, 2, 2, 94, 31, 3, 2, 2, 2, 11, 37, 40,
	51, 53, 60, 66, 73, 78, 89,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "','", "", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "SectionHeader", "SectionHeaderStart", "SectionHeaderEnd",
	"BooleanLiteral", "StringLiteral", "IntegerLiteral", "DecimalLiteral",
	"Identifier", "CommentLine", "MultiNewLine", "NewLine", "WS",
}

var ruleNames = []string{
	"main", "section", "sectionHeader", "sectionBody", "valueLine", "variableName",
	"value", "listValue", "basicValue", "commentLine", "boolValue", "stringValue",
	"numberValue", "integerValue", "decimalValue",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type IniParser struct {
	*antlr.BaseParser
}

func NewIniParser(input antlr.TokenStream) *IniParser {
	this := new(IniParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Ini.g4"

	return this
}

// IniParser tokens.
const (
	IniParserEOF                = antlr.TokenEOF
	IniParserT__0               = 1
	IniParserT__1               = 2
	IniParserSectionHeader      = 3
	IniParserSectionHeaderStart = 4
	IniParserSectionHeaderEnd   = 5
	IniParserBooleanLiteral     = 6
	IniParserStringLiteral      = 7
	IniParserIntegerLiteral     = 8
	IniParserDecimalLiteral     = 9
	IniParserIdentifier         = 10
	IniParserCommentLine        = 11
	IniParserMultiNewLine       = 12
	IniParserNewLine            = 13
	IniParserWS                 = 14
)

// IniParser rules.
const (
	IniParserRULE_main          = 0
	IniParserRULE_section       = 1
	IniParserRULE_sectionHeader = 2
	IniParserRULE_sectionBody   = 3
	IniParserRULE_valueLine     = 4
	IniParserRULE_variableName  = 5
	IniParserRULE_value         = 6
	IniParserRULE_listValue     = 7
	IniParserRULE_basicValue    = 8
	IniParserRULE_commentLine   = 9
	IniParserRULE_boolValue     = 10
	IniParserRULE_stringValue   = 11
	IniParserRULE_numberValue   = 12
	IniParserRULE_integerValue  = 13
	IniParserRULE_decimalValue  = 14
)

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) EOF() antlr.TerminalNode {
	return s.GetToken(IniParserEOF, 0)
}

func (s *MainContext) AllSection() []ISectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISectionContext)(nil)).Elem())
	var tst = make([]ISectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISectionContext)
		}
	}

	return tst
}

func (s *MainContext) Section(i int) ISectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *MainContext) AllMultiNewLine() []antlr.TerminalNode {
	return s.GetTokens(IniParserMultiNewLine)
}

func (s *MainContext) MultiNewLine(i int) antlr.TerminalNode {
	return s.GetToken(IniParserMultiNewLine, i)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitMain(s)
	}
}

func (p *IniParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IniParserRULE_main)
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == IniParserSectionHeader {
		{
			p.SetState(30)
			p.Section()
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == IniParserMultiNewLine {
			{
				p.SetState(31)
				p.Match(IniParserMultiNewLine)
			}
			{
				p.SetState(32)
				p.Section()
			}

			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(40)
		p.Match(IniParserEOF)
	}

	return localctx
}

// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) SectionHeader() ISectionHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionHeaderContext)
}

func (s *SectionContext) SectionBody() ISectionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionBodyContext)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterSection(s)
	}
}

func (s *SectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitSection(s)
	}
}

func (p *IniParser) Section() (localctx ISectionContext) {
	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IniParserRULE_section)

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
		p.SetState(42)
		p.SectionHeader()
	}
	{
		p.SetState(43)
		p.SectionBody()
	}

	return localctx
}

// ISectionHeaderContext is an interface to support dynamic dispatch.
type ISectionHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionHeaderContext differentiates from other interfaces.
	IsSectionHeaderContext()
}

type SectionHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionHeaderContext() *SectionHeaderContext {
	var p = new(SectionHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_sectionHeader
	return p
}

func (*SectionHeaderContext) IsSectionHeaderContext() {}

func NewSectionHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionHeaderContext {
	var p = new(SectionHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_sectionHeader

	return p
}

func (s *SectionHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionHeaderContext) SectionHeader() antlr.TerminalNode {
	return s.GetToken(IniParserSectionHeader, 0)
}

func (s *SectionHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterSectionHeader(s)
	}
}

func (s *SectionHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitSectionHeader(s)
	}
}

func (p *IniParser) SectionHeader() (localctx ISectionHeaderContext) {
	localctx = NewSectionHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, IniParserRULE_sectionHeader)

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
		p.SetState(45)
		p.Match(IniParserSectionHeader)
	}

	return localctx
}

// ISectionBodyContext is an interface to support dynamic dispatch.
type ISectionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionBodyContext differentiates from other interfaces.
	IsSectionBodyContext()
}

type SectionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionBodyContext() *SectionBodyContext {
	var p = new(SectionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_sectionBody
	return p
}

func (*SectionBodyContext) IsSectionBodyContext() {}

func NewSectionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionBodyContext {
	var p = new(SectionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_sectionBody

	return p
}

func (s *SectionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionBodyContext) AllValueLine() []IValueLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueLineContext)(nil)).Elem())
	var tst = make([]IValueLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueLineContext)
		}
	}

	return tst
}

func (s *SectionBodyContext) ValueLine(i int) IValueLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueLineContext)
}

func (s *SectionBodyContext) AllCommentLine() []ICommentLineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentLineContext)(nil)).Elem())
	var tst = make([]ICommentLineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentLineContext)
		}
	}

	return tst
}

func (s *SectionBodyContext) CommentLine(i int) ICommentLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentLineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentLineContext)
}

func (s *SectionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterSectionBody(s)
	}
}

func (s *SectionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitSectionBody(s)
	}
}

func (p *IniParser) SectionBody() (localctx ISectionBodyContext) {
	localctx = NewSectionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, IniParserRULE_sectionBody)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(47)
					p.ValueLine()
				}

			case 2:
				{
					p.SetState(48)
					p.CommentLine()
				}

			}

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IValueLineContext is an interface to support dynamic dispatch.
type IValueLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueLineContext differentiates from other interfaces.
	IsValueLineContext()
}

type ValueLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueLineContext() *ValueLineContext {
	var p = new(ValueLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_valueLine
	return p
}

func (*ValueLineContext) IsValueLineContext() {}

func NewValueLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueLineContext {
	var p = new(ValueLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_valueLine

	return p
}

func (s *ValueLineContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueLineContext) MultiNewLine() antlr.TerminalNode {
	return s.GetToken(IniParserMultiNewLine, 0)
}

func (s *ValueLineContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *ValueLineContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterValueLine(s)
	}
}

func (s *ValueLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitValueLine(s)
	}
}

func (p *IniParser) ValueLine() (localctx IValueLineContext) {
	localctx = NewValueLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, IniParserRULE_valueLine)
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
		p.SetState(54)
		p.Match(IniParserMultiNewLine)
	}
	{
		p.SetState(55)
		p.VariableName()
	}
	{
		p.SetState(56)
		p.Match(IniParserT__0)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<IniParserBooleanLiteral)|(1<<IniParserStringLiteral)|(1<<IniParserIntegerLiteral)|(1<<IniParserDecimalLiteral))) != 0 {
		{
			p.SetState(57)
			p.Value()
		}

	}

	return localctx
}

// IVariableNameContext is an interface to support dynamic dispatch.
type IVariableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableNameContext differentiates from other interfaces.
	IsVariableNameContext()
}

type VariableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableNameContext() *VariableNameContext {
	var p = new(VariableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_variableName
	return p
}

func (*VariableNameContext) IsVariableNameContext() {}

func NewVariableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableNameContext {
	var p = new(VariableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_variableName

	return p
}

func (s *VariableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(IniParserIdentifier, 0)
}

func (s *VariableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterVariableName(s)
	}
}

func (s *VariableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitVariableName(s)
	}
}

func (p *IniParser) VariableName() (localctx IVariableNameContext) {
	localctx = NewVariableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, IniParserRULE_variableName)

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
		p.Match(IniParserIdentifier)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) BasicValue() IBasicValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBasicValueContext)
}

func (s *ValueContext) ListValue() IListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListValueContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *IniParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, IniParserRULE_value)

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

	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.BasicValue()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.ListValue()
		}

	}

	return localctx
}

// IListValueContext is an interface to support dynamic dispatch.
type IListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListValueContext differentiates from other interfaces.
	IsListValueContext()
}

type ListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListValueContext() *ListValueContext {
	var p = new(ListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_listValue
	return p
}

func (*ListValueContext) IsListValueContext() {}

func NewListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListValueContext {
	var p = new(ListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_listValue

	return p
}

func (s *ListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ListValueContext) AllBasicValue() []IBasicValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBasicValueContext)(nil)).Elem())
	var tst = make([]IBasicValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBasicValueContext)
		}
	}

	return tst
}

func (s *ListValueContext) BasicValue(i int) IBasicValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBasicValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBasicValueContext)
}

func (s *ListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterListValue(s)
	}
}

func (s *ListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitListValue(s)
	}
}

func (p *IniParser) ListValue() (localctx IListValueContext) {
	localctx = NewListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, IniParserRULE_listValue)
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
		p.SetState(66)
		p.BasicValue()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == IniParserT__1 {
		{
			p.SetState(67)
			p.Match(IniParserT__1)
		}
		{
			p.SetState(68)
			p.BasicValue()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBasicValueContext is an interface to support dynamic dispatch.
type IBasicValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasicValueContext differentiates from other interfaces.
	IsBasicValueContext()
}

type BasicValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasicValueContext() *BasicValueContext {
	var p = new(BasicValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_basicValue
	return p
}

func (*BasicValueContext) IsBasicValueContext() {}

func NewBasicValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasicValueContext {
	var p = new(BasicValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_basicValue

	return p
}

func (s *BasicValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BasicValueContext) BoolValue() IBoolValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolValueContext)
}

func (s *BasicValueContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *BasicValueContext) NumberValue() INumberValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberValueContext)
}

func (s *BasicValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasicValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasicValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterBasicValue(s)
	}
}

func (s *BasicValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitBasicValue(s)
	}
}

func (p *IniParser) BasicValue() (localctx IBasicValueContext) {
	localctx = NewBasicValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, IniParserRULE_basicValue)

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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case IniParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.BoolValue()
		}

	case IniParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.StringValue()
		}

	case IniParserIntegerLiteral, IniParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.NumberValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommentLineContext is an interface to support dynamic dispatch.
type ICommentLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentLineContext differentiates from other interfaces.
	IsCommentLineContext()
}

type CommentLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentLineContext() *CommentLineContext {
	var p = new(CommentLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_commentLine
	return p
}

func (*CommentLineContext) IsCommentLineContext() {}

func NewCommentLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentLineContext {
	var p = new(CommentLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_commentLine

	return p
}

func (s *CommentLineContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentLineContext) MultiNewLine() antlr.TerminalNode {
	return s.GetToken(IniParserMultiNewLine, 0)
}

func (s *CommentLineContext) CommentLine() antlr.TerminalNode {
	return s.GetToken(IniParserCommentLine, 0)
}

func (s *CommentLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterCommentLine(s)
	}
}

func (s *CommentLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitCommentLine(s)
	}
}

func (p *IniParser) CommentLine() (localctx ICommentLineContext) {
	localctx = NewCommentLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, IniParserRULE_commentLine)

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
		p.SetState(78)
		p.Match(IniParserMultiNewLine)
	}
	{
		p.SetState(79)
		p.Match(IniParserCommentLine)
	}

	return localctx
}

// IBoolValueContext is an interface to support dynamic dispatch.
type IBoolValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolValueContext differentiates from other interfaces.
	IsBoolValueContext()
}

type BoolValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolValueContext() *BoolValueContext {
	var p = new(BoolValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_boolValue
	return p
}

func (*BoolValueContext) IsBoolValueContext() {}

func NewBoolValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolValueContext {
	var p = new(BoolValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_boolValue

	return p
}

func (s *BoolValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolValueContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(IniParserBooleanLiteral, 0)
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (p *IniParser) BoolValue() (localctx IBoolValueContext) {
	localctx = NewBoolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, IniParserRULE_boolValue)

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
		p.SetState(81)
		p.Match(IniParserBooleanLiteral)
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(IniParserStringLiteral, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *IniParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, IniParserRULE_stringValue)

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
		p.SetState(83)
		p.Match(IniParserStringLiteral)
	}

	return localctx
}

// INumberValueContext is an interface to support dynamic dispatch.
type INumberValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberValueContext differentiates from other interfaces.
	IsNumberValueContext()
}

type NumberValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberValueContext() *NumberValueContext {
	var p = new(NumberValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_numberValue
	return p
}

func (*NumberValueContext) IsNumberValueContext() {}

func NewNumberValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberValueContext {
	var p = new(NumberValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_numberValue

	return p
}

func (s *NumberValueContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberValueContext) IntegerValue() IIntegerValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
}

func (s *NumberValueContext) DecimalValue() IDecimalValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalValueContext)
}

func (s *NumberValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterNumberValue(s)
	}
}

func (s *NumberValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitNumberValue(s)
	}
}

func (p *IniParser) NumberValue() (localctx INumberValueContext) {
	localctx = NewNumberValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, IniParserRULE_numberValue)

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
	case IniParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.IntegerValue()
		}

	case IniParserDecimalLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.DecimalValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerValueContext is an interface to support dynamic dispatch.
type IIntegerValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerValueContext differentiates from other interfaces.
	IsIntegerValueContext()
}

type IntegerValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerValueContext() *IntegerValueContext {
	var p = new(IntegerValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_integerValue
	return p
}

func (*IntegerValueContext) IsIntegerValueContext() {}

func NewIntegerValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerValueContext {
	var p = new(IntegerValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_integerValue

	return p
}

func (s *IntegerValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerValueContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(IniParserIntegerLiteral, 0)
}

func (s *IntegerValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterIntegerValue(s)
	}
}

func (s *IntegerValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitIntegerValue(s)
	}
}

func (p *IniParser) IntegerValue() (localctx IIntegerValueContext) {
	localctx = NewIntegerValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, IniParserRULE_integerValue)

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
		p.Match(IniParserIntegerLiteral)
	}

	return localctx
}

// IDecimalValueContext is an interface to support dynamic dispatch.
type IDecimalValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalValueContext differentiates from other interfaces.
	IsDecimalValueContext()
}

type DecimalValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalValueContext() *DecimalValueContext {
	var p = new(DecimalValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IniParserRULE_decimalValue
	return p
}

func (*DecimalValueContext) IsDecimalValueContext() {}

func NewDecimalValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalValueContext {
	var p = new(DecimalValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IniParserRULE_decimalValue

	return p
}

func (s *DecimalValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalValueContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(IniParserDecimalLiteral, 0)
}

func (s *DecimalValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.EnterDecimalValue(s)
	}
}

func (s *DecimalValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IniListener); ok {
		listenerT.ExitDecimalValue(s)
	}
}

func (p *IniParser) DecimalValue() (localctx IDecimalValueContext) {
	localctx = NewDecimalValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, IniParserRULE_decimalValue)

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
		p.SetState(91)
		p.Match(IniParserDecimalLiteral)
	}

	return localctx
}
