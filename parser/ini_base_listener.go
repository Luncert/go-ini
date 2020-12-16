// Code generated from parser/Ini.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Ini

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIniListener is a complete listener for a parse tree produced by IniParser.
type BaseIniListener struct{}

var _ IniListener = &BaseIniListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIniListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIniListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIniListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIniListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMain is called when production main is entered.
func (s *BaseIniListener) EnterMain(ctx *MainContext) {}

// ExitMain is called when production main is exited.
func (s *BaseIniListener) ExitMain(ctx *MainContext) {}

// EnterSection is called when production section is entered.
func (s *BaseIniListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BaseIniListener) ExitSection(ctx *SectionContext) {}

// EnterSectionHeader is called when production sectionHeader is entered.
func (s *BaseIniListener) EnterSectionHeader(ctx *SectionHeaderContext) {}

// ExitSectionHeader is called when production sectionHeader is exited.
func (s *BaseIniListener) ExitSectionHeader(ctx *SectionHeaderContext) {}

// EnterSectionBody is called when production sectionBody is entered.
func (s *BaseIniListener) EnterSectionBody(ctx *SectionBodyContext) {}

// ExitSectionBody is called when production sectionBody is exited.
func (s *BaseIniListener) ExitSectionBody(ctx *SectionBodyContext) {}

// EnterValueLine is called when production valueLine is entered.
func (s *BaseIniListener) EnterValueLine(ctx *ValueLineContext) {}

// ExitValueLine is called when production valueLine is exited.
func (s *BaseIniListener) ExitValueLine(ctx *ValueLineContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BaseIniListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BaseIniListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterValue is called when production value is entered.
func (s *BaseIniListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseIniListener) ExitValue(ctx *ValueContext) {}

// EnterListValue is called when production listValue is entered.
func (s *BaseIniListener) EnterListValue(ctx *ListValueContext) {}

// ExitListValue is called when production listValue is exited.
func (s *BaseIniListener) ExitListValue(ctx *ListValueContext) {}

// EnterBasicValue is called when production basicValue is entered.
func (s *BaseIniListener) EnterBasicValue(ctx *BasicValueContext) {}

// ExitBasicValue is called when production basicValue is exited.
func (s *BaseIniListener) ExitBasicValue(ctx *BasicValueContext) {}

// EnterCommentLine is called when production commentLine is entered.
func (s *BaseIniListener) EnterCommentLine(ctx *CommentLineContext) {}

// ExitCommentLine is called when production commentLine is exited.
func (s *BaseIniListener) ExitCommentLine(ctx *CommentLineContext) {}

// EnterBoolValue is called when production boolValue is entered.
func (s *BaseIniListener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production boolValue is exited.
func (s *BaseIniListener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseIniListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseIniListener) ExitStringValue(ctx *StringValueContext) {}

// EnterNumberValue is called when production numberValue is entered.
func (s *BaseIniListener) EnterNumberValue(ctx *NumberValueContext) {}

// ExitNumberValue is called when production numberValue is exited.
func (s *BaseIniListener) ExitNumberValue(ctx *NumberValueContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseIniListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseIniListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterDecimalValue is called when production decimalValue is entered.
func (s *BaseIniListener) EnterDecimalValue(ctx *DecimalValueContext) {}

// ExitDecimalValue is called when production decimalValue is exited.
func (s *BaseIniListener) ExitDecimalValue(ctx *DecimalValueContext) {}
