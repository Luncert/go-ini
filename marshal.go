package go_ini

import (
	"io/ioutil"
	"sort"
	"strings"
)

func WriteConfigFile(filename string, cfg *Config) error {
	s := Marshal(cfg)
	return ioutil.WriteFile(filename, []byte(s), 0644)
}

func Marshal(cfg *Config) string {
	buf := strings.Builder{}

	i := 0
	secNames := make([]string, len(cfg.Sections))
	for name := range cfg.Sections {
		secNames[i] = name
		i++
	}
	sort.Strings(secNames)
	for _, name := range secNames {
		sec := cfg.Section(name)

		buf.WriteRune('[')
		buf.WriteString(name)
		buf.WriteString("]\n")

		i = 0
		varNames := make([]string, len(sec.Variables))
		for name := range sec.Variables {
			varNames[i] = name
			i++
		}
		sort.Strings(varNames)
		for _, name := range varNames {
			buf.WriteString(name)
			buf.WriteString(" = ")
			buf.WriteString(sec.Variable(name).Value.ToString())
			buf.WriteRune('\n')
		}
	}

	return buf.String()
}
