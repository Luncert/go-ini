// Code generated from parser/Ini.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Ini

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IniListener is a complete listener for a parse tree produced by IniParser.
type IniListener interface {
	antlr.ParseTreeListener

	// EnterMain is called when entering the main production.
	EnterMain(c *MainContext)

	// EnterSection is called when entering the section production.
	EnterSection(c *SectionContext)

	// EnterSectionHeader is called when entering the sectionHeader production.
	EnterSectionHeader(c *SectionHeaderContext)

	// EnterSectionBody is called when entering the sectionBody production.
	EnterSectionBody(c *SectionBodyContext)

	// EnterValueLine is called when entering the valueLine production.
	EnterValueLine(c *ValueLineContext)

	// EnterVariableName is called when entering the variableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterListValue is called when entering the listValue production.
	EnterListValue(c *ListValueContext)

	// EnterBasicValue is called when entering the basicValue production.
	EnterBasicValue(c *BasicValueContext)

	// EnterCommentLine is called when entering the commentLine production.
	EnterCommentLine(c *CommentLineContext)

	// EnterBoolValue is called when entering the boolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterNumberValue is called when entering the numberValue production.
	EnterNumberValue(c *NumberValueContext)

	// EnterIntegerValue is called when entering the integerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterDecimalValue is called when entering the decimalValue production.
	EnterDecimalValue(c *DecimalValueContext)

	// ExitMain is called when exiting the main production.
	ExitMain(c *MainContext)

	// ExitSection is called when exiting the section production.
	ExitSection(c *SectionContext)

	// ExitSectionHeader is called when exiting the sectionHeader production.
	ExitSectionHeader(c *SectionHeaderContext)

	// ExitSectionBody is called when exiting the sectionBody production.
	ExitSectionBody(c *SectionBodyContext)

	// ExitValueLine is called when exiting the valueLine production.
	ExitValueLine(c *ValueLineContext)

	// ExitVariableName is called when exiting the variableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitListValue is called when exiting the listValue production.
	ExitListValue(c *ListValueContext)

	// ExitBasicValue is called when exiting the basicValue production.
	ExitBasicValue(c *BasicValueContext)

	// ExitCommentLine is called when exiting the commentLine production.
	ExitCommentLine(c *CommentLineContext)

	// ExitBoolValue is called when exiting the boolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitNumberValue is called when exiting the numberValue production.
	ExitNumberValue(c *NumberValueContext)

	// ExitIntegerValue is called when exiting the integerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitDecimalValue is called when exiting the decimalValue production.
	ExitDecimalValue(c *DecimalValueContext)
}
