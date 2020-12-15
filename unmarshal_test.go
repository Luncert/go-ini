package go_ini

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestParserIni(t *testing.T) {
	testData, err := ioutil.ReadFile("./data/test.ini")
	assert.Nil(t, err)
	Unmarshal(string(testData))
}
