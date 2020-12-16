package go_ini

import (
	"fmt"
	"github.com/Luncert/go-ini/util"
	"sort"
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

func NewConfig() *Config {
	return &Config{Sections: make(map[string]*Section)}
}

func (c *Config) Section(name string) *Section {
	return c.Sections[name]
}

func (c *Config) AddSection(section *Section) *Config {
	if len(section.Name) == 0 {
		panic("invalid section name")
	}
	if _, ok := c.Sections[section.Name]; ok {
		panic("duplicated section name " + section.Name)
	}

	c.Sections[section.Name] = section
	return c
}

func (c *Config) CreateIfAbsent(sectionName string) *Section {
	sec, ok := c.Sections[sectionName]
	if ok {
		return sec
	}
	sec = &Section{Name: sectionName}
	c.Sections[sectionName] = sec
	return sec
}

func (c *Config) SetIfAbsent(section *Section) *Section {
	sec, ok := c.Sections[section.Name]
	if ok {
		return sec
	}
	c.Sections[section.Name] = section
	return section
}

func (c *Config) ToString() string {
	builder := util.NewStringBuilder(indentSpacing)
	builder.WriteLine("Config {")
	builder.IncIndent()

	builder.WriteLine("Sections = [")
	builder.IncIndent()

	// sort by section names and print each section
	names := make([]string, 0)
	for name := range c.Sections {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		builder.WriteMultiLine(c.Sections[name].ToString(), true)
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

func NewSection(name string) *Section {
	return &Section{
		Name:      name,
		Variables: make(map[string]*Variable),
	}
}

func (s *Section) Variable(name string) *Variable {
	return s.Variables[name]
}

func (s *Section) AddVariable(v *Variable) *Section {
	if len(v.Name) == 0 {
		panic("invalid variable name")
	}
	if _, ok := s.Variables[v.Name]; ok {
		panic("duplicated variable name in section " + s.Name)
	}

	s.Variables[v.Name] = v
	return s
}

func (s *Section) CreateIfAbsent(name string) *Variable {
	v, ok := s.Variables[name]
	if ok {
		return v
	}
	v = &Variable{Name: name}
	s.Variables[name] = v
	return v
}

func (s *Section) SetIfAbsent(variable *Variable) *Variable {
	v, ok := s.Variables[variable.Name]
	if ok {
		return v
	}
	s.Variables[variable.Name] = variable
	return variable
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

	// sort by variable names and print each variable
	names := make([]string, 0)
	for name := range s.Variables {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		v := s.Variables[name]
		builder.WriteWithIndent(name, strings.Repeat(" ", maxNameLen-len(name)), " = ")
		builder.WriteMultiLine(v.Value.ToString(), false)
	}

	builder.DecIndent()
	builder.WriteLine("}")

	builder.DecIndent()
	builder.WriteWithIndent("}")
	return builder.String()
}

type Variable struct {
	Name  string
	Value VariableValue
}

// NewVariable creates a new variable, supported value type are: bool, string, int, float64 and slice of these primitive types
func NewVariable(name string, value interface{}) *Variable {
	var v VariableValue
	switch value.(type) {
	case bool:
		v = NewBoolValue(value.(bool))
	case string:
		v = NewStringValue(value.(string))
	case int:
		v = NewIntegerValue(value.(int))
	case float64:
		v = NewDecimalValue(value.(float64))
	default:
		list := &ListValue{}
		switch value.(type) {
		case []bool:
			for _, tmp := range value.([]bool) {
				list.V = append(list.V, NewBoolValue(tmp))
			}
		case []string:
			for _, tmp := range value.([]string) {
				list.V = append(list.V, NewStringValue(tmp))
			}
		case []int:
			for _, tmp := range value.([]int) {
				list.V = append(list.V, NewIntegerValue(tmp))
			}
		case []float64:
			for _, tmp := range value.([]float64) {
				list.V = append(list.V, NewDecimalValue(tmp))
			}
		default:
			panic("unacceptable variable value")
		}
		v = list
	}
	return &Variable{Name: name, Value: v}
}

// Type equals to VariableValue.Type()
func (v *Variable) Type() int {
	return v.Value.Type()
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

func NewBoolValue(v bool) *BoolValue {
	return &BoolValue{V: v}
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

func NewStringValue(v string) *StringValue {
	return &StringValue{V: v}
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

func NewIntegerValue(v int) *IntegerValue {
	return &IntegerValue{V: v}
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

func NewDecimalValue(v float64) *DecimalValue {
	return &DecimalValue{V: v}
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

func NewListValue(v ...VariableValue) *ListValue {
	return &ListValue{V: v}
}

func (l *ListValue) Type() int {
	return ListType
}

func (l *ListValue) Value() interface{} {
	return l.V
}

func (l *ListValue) ToString() string {
	builder := util.NewStringBuilder(indentSpacing)

	if len(l.V) > 0 {
		builder.WriteString(l.V[0].ToString())
		for i := 1; i < len(l.V); i++ {
			builder.WriteString(", " + l.V[i].ToString())
		}
	}

	return builder.String()
}
