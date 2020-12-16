package go_ini

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	testData := "[1]\nname = \"asd\""
	cfg := Unmarshal(testData)
	assert.Equal(t, "asd", cfg.Section("1").Variable("name").String())
}

func TestUnmarshalFile(t *testing.T) {
	cfg, _ := ReadConfigFile("./data/test.ini")
	assert.Equal(t, 12309809123, cfg.Section("1").Variable("x_123A").Int())
	assert.Equal(t, true, cfg.Section("1").Variable("f").Bool())
	assert.Equal(t, "d09ais0j001090i", cfg.Section("1").Variable("j").String())
}
