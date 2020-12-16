package go_ini

import (
	"github.com/Luncert/go-ini/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strconv"
)

func ReadConfigFile(filename string) (*Config, error) {
	fs, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}
	return parse(fs.InputStream), nil
}

func Unmarshal(data string) *Config {
	input := antlr.NewInputStream(data)
	return parse(input)
}

func parse(input *antlr.InputStream) *Config {
	lexer := parser.NewIniLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewIniParser(stream)
	p.AddErrorListener(newErrorListener()) // default is console error listener
	p.BuildParseTrees = true
	tree := p.Main()
	listener := newIniListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.config
}

// ErrorListener
type errorListener struct {
	*antlr.DiagnosticErrorListener
}

// NewErrorListener
func newErrorListener() *errorListener {
	return &errorListener{antlr.NewDiagnosticErrorListener(true)}
}

// SyntaxError
func (el *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	el.DiagnosticErrorListener.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
	os.Exit(1)
}

type iniListener struct {
	*parser.BaseIniListener
	config       *Config
	lastSection  *Section
	lastVariable *Variable
}

func newIniListener() *iniListener {
	cfg := &Config{
		Sections: make(map[string]*Section),
	}
	return &iniListener{config: cfg}
}

func (i *iniListener) EnterSectionHeader(ctx *parser.SectionHeaderContext) {
	sectionName := ctx.GetText()
	sectionName = sectionName[1 : len(sectionName)-1]

	if _, ok := i.config.Sections[sectionName]; ok {
		panic(DuplicatedSectionNameError(sectionName))
	}

	section := &Section{Name: sectionName, Variables: make(map[string]*Variable)}
	i.config.Sections[sectionName] = section
	i.lastSection = section
}

func (i *iniListener) EnterVariableName(ctx *parser.VariableNameContext) {
	variableName := ctx.GetText()
	variable := &Variable{Name: variableName}

	if _, ok := i.lastSection.Variables[variableName]; ok {
		panic(DuplicatedVariableNameError(variableName, i.lastSection.Name))
	}

	i.lastSection.Variables[variableName] = variable
	i.lastVariable = variable
}

func (i *iniListener) setValue(v VariableValue) {
	if i.lastVariable.Value != nil {
		value := i.lastVariable.Value.(*ListValue)
		value.V = append(value.V, v)

	} else {
		i.lastVariable.Value = v
	}
}

func (i *iniListener) EnterListValue(ctx *parser.ListValueContext) {
	i.lastVariable.Value = &ListValue{V: make([]VariableValue, 0)}
}

func (i *iniListener) EnterBoolValue(ctx *parser.BoolValueContext) {
	text := ctx.GetText()
	v := text[0] == 'Y' || text[0] == 'y'
	i.setValue(NewBoolValue(v))
}

func (i *iniListener) EnterStringValue(ctx *parser.StringValueContext) {
	text := ctx.GetText()
	v := text[1 : len(text)-1]
	i.setValue(NewStringValue(v))
}

func (i *iniListener) EnterIntegerValue(ctx *parser.IntegerValueContext) {
	text := ctx.GetText()
	var number int
	if len(text) > 2 && text[1] == 'x' {
		if v, err := strconv.ParseInt(text, 0, 64); err != nil {
			panic(err)
		} else {
			number = int(v)
		}
	} else {
		var err error
		if number, err = strconv.Atoi(text); err != nil {
			panic(err)
		}
	}
	i.setValue(NewIntegerValue(number))
}

func (i *iniListener) EnterDecimalValue(ctx *parser.DecimalValueContext) {
	text := ctx.GetText()
	number, err := strconv.ParseFloat(text, 64)
	if err != nil {
		panic(err)
	}
	i.setValue(NewDecimalValue(number))
}
