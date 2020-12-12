package go_ini

import (
	"fmt"
	"github.com/Luncert/go-ini/util"
	"strconv"
	"strings"
)

const (
	BoolType = iota
	StringType
	IntegerType
	DecimalType
	ListType
)
const indentSpacing = "  "

type Config struct {
	Sections map[string]*Section
}

func (c *Config) Section(name string) *Section {
	return c.Sections[name]
}

func (c *Config) ToString() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("Config {")
	builder.IncIndent()

	builder.WriteLine("Sections = [")
	builder.IncIndent()

	for _, section := range c.Sections {
		builder.WriteMultiLine(section.ToString(), true)
	}

	builder.DecIndent()
	builder.WriteLine("]")

	builder.DecIndent()
	builder.WriteLine("}")
	return builder.String()
}

type Section struct {
	Name      string
	Variables map[string]*Variable
}

func (s *Section) Variable(name string) *Variable {
	return s.Variables[name]
}

func (s *Section) ToString() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("Section(", s.Name, ") {")
	builder.IncIndent()

	builder.WriteLine("Variables = {")
	builder.IncIndent()

	var maxNameLen int
	for name := range s.Variables {
		if maxNameLen < len(name) {
			maxNameLen = len(name)
		}
	}

	for name, variable := range s.Variables {
		builder.WriteWithIndent(name, strings.Repeat(" ", maxNameLen-len(name)), " = ")
		builder.WriteMultiLine(variable.Value.ToString(), false)
	}

	builder.DecIndent()
	builder.WriteLine("}")

	builder.DecIndent()
	builder.WriteWithIndent("}")
	return builder.String()
}

type Variable struct {
	Value VariableValue
}

func (v *Variable) Bool() bool {
	return v.Value.Value().(bool)
}

func (v *Variable) String() string {
	return v.Value.Value().(string)
}

func (v *Variable) Int() int {
	return v.Value.Value().(int)
}

func (v *Variable) Float32() float32 {
	return float32(v.Float64())
}

func (v *Variable) Float64() float64 {
	return v.Value.Value().(float64)
}

func (v *Variable) List() []VariableValue {
	return v.Value.Value().([]VariableValue)
}

type VariableValue interface {
	Type() int
	Value() interface{}
	ToString() string
}

type BoolValue struct {
	V bool
}

func (b *BoolValue) Type() int {
	return BoolType
}

func (b *BoolValue) Value() interface{} {
	return b.V
}

func (b *BoolValue) ToString() string {
	if b.V {
		return "yes"
	}
	return "no"
}

type StringValue struct {
	V string
}

func (s *StringValue) Type() int {
	return StringType
}

func (s *StringValue) Value() interface{} {
	return s.V
}

func (s *StringValue) ToString() string {
	return fmt.Sprintf("\"%s\"", s.V)
}

type IntegerValue struct {
	V int
}

func (i *IntegerValue) Type() int {
	return IntegerType
}

func (i *IntegerValue) Value() interface{} {
	return i.V
}

func (i *IntegerValue) ToString() string {
	return strconv.Itoa(i.V)
}

type DecimalValue struct {
	V float64
}

func (d *DecimalValue) Type() int {
	return DecimalType
}

func (d *DecimalValue) Value() interface{} {
	return d.V
}

func (d *DecimalValue) ToString() string {
	return fmt.Sprintf("%f", d.V)
}

type ListValue struct {
	V []VariableValue
}

func (l *ListValue) Type() int {
	return ListType
}

func (l *ListValue) Value() interface{} {
	return l.V
}

func (l *ListValue) ToString() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteWithIndent("[")

	if len(l.V) > 0 {
		builder.WriteString(l.V[0].ToString())
		for i := 1; i < len(l.V); i++ {
			builder.WriteString(", " + l.V[i].ToString())
		}
	}

	builder.WriteString("]")
	return builder.String()
}
