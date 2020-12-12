package go_ini

import (
	"errors"
	"fmt"
)

func DuplicatedSectionNameError(sectionName string) error {
	return errors.New("duplicated section name " + sectionName)
}

func DuplicatedVariableNameError(variableName, sectionName string) error {
	return fmt.Errorf("duplicated variable name %s in section %s", variableName, sectionName)
}
