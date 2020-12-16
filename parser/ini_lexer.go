// Code generated from parser/Ini.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 161,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 6, 7, 61, 10, 7, 13, 7, 14, 7, 62, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 5, 8, 70, 10, 8, 3, 9, 3, 9, 7, 9, 74, 10, 9, 12, 9, 14, 9, 77, 11,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 6, 10, 86, 10, 10, 13,
	10, 14, 10, 87, 3, 10, 5, 10, 91, 10, 10, 3, 10, 3, 10, 5, 10, 95, 10,
	10, 5, 10, 97, 10, 10, 3, 11, 5, 11, 100, 10, 11, 3, 11, 3, 11, 3, 11,
	5, 11, 105, 10, 11, 3, 11, 3, 11, 5, 11, 109, 10, 11, 3, 12, 3, 12, 3,
	13, 6, 13, 114, 10, 13, 13, 13, 14, 13, 115, 3, 14, 3, 14, 5, 14, 120,
	10, 14, 3, 15, 3, 15, 3, 16, 3, 16, 7, 16, 126, 10, 16, 12, 16, 14, 16,
	129, 11, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	20, 7, 20, 140, 10, 20, 12, 20, 14, 20, 143, 11, 20, 3, 21, 6, 21, 146,
	10, 21, 13, 21, 14, 21, 147, 3, 22, 5, 22, 151, 10, 22, 3, 22, 3, 22, 3,
	23, 6, 23, 156, 10, 23, 13, 23, 14, 23, 157, 3, 23, 3, 23, 2, 2, 24, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 2, 15, 8, 17, 9, 19, 10, 21, 11, 23, 2,
	25, 2, 27, 2, 29, 2, 31, 12, 33, 2, 35, 2, 37, 13, 39, 2, 41, 14, 43, 15,
	45, 16, 3, 2, 13, 4, 2, 93, 93, 95, 95, 4, 2, 91, 91, 123, 123, 4, 2, 80,
	80, 112, 112, 3, 2, 36, 36, 5, 2, 51, 59, 67, 72, 99, 104, 3, 2, 51, 59,
	4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 12,
	12, 3, 2, 15, 15, 5, 2, 11, 11, 15, 15, 34, 34, 2, 170, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3,
	2, 2, 2, 2, 45, 3, 2, 2, 2, 3, 47, 3, 2, 2, 2, 5, 49, 3, 2, 2, 2, 7, 51,
	3, 2, 2, 2, 9, 55, 3, 2, 2, 2, 11, 57, 3, 2, 2, 2, 13, 60, 3, 2, 2, 2,
	15, 69, 3, 2, 2, 2, 17, 71, 3, 2, 2, 2, 19, 96, 3, 2, 2, 2, 21, 108, 3,
	2, 2, 2, 23, 110, 3, 2, 2, 2, 25, 113, 3, 2, 2, 2, 27, 119, 3, 2, 2, 2,
	29, 121, 3, 2, 2, 2, 31, 123, 3, 2, 2, 2, 33, 130, 3, 2, 2, 2, 35, 132,
	3, 2, 2, 2, 37, 134, 3, 2, 2, 2, 39, 141, 3, 2, 2, 2, 41, 145, 3, 2, 2,
	2, 43, 150, 3, 2, 2, 2, 45, 155, 3, 2, 2, 2, 47, 48, 7, 63, 2, 2, 48, 4,
	3, 2, 2, 2, 49, 50, 7, 46, 2, 2, 50, 6, 3, 2, 2, 2, 51, 52, 5, 9, 5, 2,
	52, 53, 5, 13, 7, 2, 53, 54, 5, 11, 6, 2, 54, 8, 3, 2, 2, 2, 55, 56, 7,
	93, 2, 2, 56, 10, 3, 2, 2, 2, 57, 58, 7, 95, 2, 2, 58, 12, 3, 2, 2, 2,
	59, 61, 10, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 60, 3,
	2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 14, 3, 2, 2, 2, 64, 65, 9, 3, 2, 2, 65,
	66, 7, 103, 2, 2, 66, 70, 7, 117, 2, 2, 67, 68, 9, 4, 2, 2, 68, 70, 7,
	113, 2, 2, 69, 64, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 16, 3, 2, 2, 2,
	71, 75, 7, 36, 2, 2, 72, 74, 10, 5, 2, 2, 73, 72, 3, 2, 2, 2, 74, 77, 3,
	2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77,
	75, 3, 2, 2, 2, 78, 79, 7, 36, 2, 2, 79, 18, 3, 2, 2, 2, 80, 97, 7, 50,
	2, 2, 81, 82, 7, 50, 2, 2, 82, 83, 7, 122, 2, 2, 83, 85, 3, 2, 2, 2, 84,
	86, 5, 23, 12, 2, 85, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 85, 3, 2,
	2, 2, 87, 88, 3, 2, 2, 2, 88, 97, 3, 2, 2, 2, 89, 91, 7, 47, 2, 2, 90,
	89, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 94, 5, 29,
	15, 2, 93, 95, 5, 25, 13, 2, 94, 93, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95,
	97, 3, 2, 2, 2, 96, 80, 3, 2, 2, 2, 96, 81, 3, 2, 2, 2, 96, 90, 3, 2, 2,
	2, 97, 20, 3, 2, 2, 2, 98, 100, 7, 47, 2, 2, 99, 98, 3, 2, 2, 2, 99, 100,
	3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 5, 25, 13, 2, 102, 104, 7,
	48, 2, 2, 103, 105, 5, 25, 13, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2,
	2, 2, 105, 109, 3, 2, 2, 2, 106, 107, 7, 48, 2, 2, 107, 109, 5, 25, 13,
	2, 108, 99, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 22, 3, 2, 2, 2, 110,
	111, 9, 6, 2, 2, 111, 24, 3, 2, 2, 2, 112, 114, 5, 27, 14, 2, 113, 112,
	3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 3, 2,
	2, 2, 116, 26, 3, 2, 2, 2, 117, 120, 7, 50, 2, 2, 118, 120, 5, 29, 15,
	2, 119, 117, 3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 28, 3, 2, 2, 2, 121,
	122, 9, 7, 2, 2, 122, 30, 3, 2, 2, 2, 123, 127, 5, 33, 17, 2, 124, 126,
	5, 35, 18, 2, 125, 124, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3,
	2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 32, 3, 2, 2, 2, 129, 127, 3, 2, 2,
	2, 130, 131, 9, 8, 2, 2, 131, 34, 3, 2, 2, 2, 132, 133, 9, 9, 2, 2, 133,
	36, 3, 2, 2, 2, 134, 135, 5, 43, 22, 2, 135, 136, 7, 61, 2, 2, 136, 137,
	5, 39, 20, 2, 137, 38, 3, 2, 2, 2, 138, 140, 10, 10, 2, 2, 139, 138, 3,
	2, 2, 2, 140, 143, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2,
	2, 142, 40, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 144, 146, 5, 43, 22, 2, 145,
	144, 3, 2, 2, 2, 146, 147, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148,
	3, 2, 2, 2, 148, 42, 3, 2, 2, 2, 149, 151, 9, 11, 2, 2, 150, 149, 3, 2,
	2, 2, 150, 151, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 9, 10, 2, 2,
	153, 44, 3, 2, 2, 2, 154, 156, 9, 12, 2, 2, 155, 154, 3, 2, 2, 2, 156,
	157, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159,
	3, 2, 2, 2, 159, 160, 8, 23, 2, 2, 160, 46, 3, 2, 2, 2, 20, 2, 62, 69,
	75, 87, 90, 94, 96, 99, 104, 108, 115, 119, 127, 141, 147, 150, 157, 3,
	8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "','", "", "'['", "']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "SectionHeader", "SectionHeaderStart", "SectionHeaderEnd",
	"BooleanLiteral", "StringLiteral", "IntegerLiteral", "DecimalLiteral",
	"Identifier", "CommentLine", "MultiNewLine", "NewLine", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "SectionHeader", "SectionHeaderStart", "SectionHeaderEnd",
	"SectionName", "BooleanLiteral", "StringLiteral", "IntegerLiteral", "DecimalLiteral",
	"HexDigit", "Digits", "Digit", "NonZeroDigit", "Identifier", "Letter",
	"LetterOrDigitOrUnderscore", "CommentLine", "CommentLiteral", "MultiNewLine",
	"NewLine", "WS",
}

type IniLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewIniLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *IniLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewIniLexer(input antlr.CharStream) *IniLexer {
	l := new(IniLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Ini.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// IniLexer tokens.
const (
	IniLexerT__0               = 1
	IniLexerT__1               = 2
	IniLexerSectionHeader      = 3
	IniLexerSectionHeaderStart = 4
	IniLexerSectionHeaderEnd   = 5
	IniLexerBooleanLiteral     = 6
	IniLexerStringLiteral      = 7
	IniLexerIntegerLiteral     = 8
	IniLexerDecimalLiteral     = 9
	IniLexerIdentifier         = 10
	IniLexerCommentLine        = 11
	IniLexerMultiNewLine       = 12
	IniLexerNewLine            = 13
	IniLexerWS                 = 14
)
