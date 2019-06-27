package ediParser

import (
	"fmt"
)

type KeywordError struct {
	lit string
}

func (e *KeywordError) Error() string {
	return fmt.Sprintf("Got %s, Expected one of %v", e.lit, kwStrings)
}
