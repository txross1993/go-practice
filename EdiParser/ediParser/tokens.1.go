package ediParser

import (
	"reflect"
	"strconv"
	"strings"
)

type EdiToken int

const (
	ILLEGAL EdiToken = iota
	DELIM
	IDENT // literals, batch number, names, etc.
	EOF   // end of file
	WS    // whitespace
	NL    // New Line

	keyword_beg
	//Keywords
	ISA
	GS
	GE
	IEA
	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	DELIM:   "",
	IDENT:   "IDENT",
	EOF:     "EOF",
	WS:      " ",
	NL:      "\\n",
	//Keywords
	ISA: "ISA",
	GS:  "GS",
	GE:  "GE", // Record count
	IEA: "IEA",
}

func (tok EdiToken) String() string {
	s := ""
	if 0 <= tok && tok < EdiToken(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

var keywords map[string]EdiToken
var kwStrings []reflect.Value

// var separators map[string]EdiToken

func init() {
	keywords = make(map[string]EdiToken)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}

	kwStrings = reflect.ValueOf(keywords).MapKeys()
}

func LookupToken(ident string) EdiToken {
	if tok, is_keyword := keywords[strings.ToUpper(ident)]; is_keyword {
		return tok
	}
	return ILLEGAL
}
