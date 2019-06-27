package ediParser

import (
	"strconv"
	"strings"
)

type EdiToken int

const (
	ILLEGAL EdiToken = iota
	IDENT            // literals, batch number, names, etc.
	EOF              // end of file
	WS               // whitespace

	keyword_beg
	//Keywords
	ISA
	GS
	ST
	W12
	N9
	SE
	GE
	IEA
	keyword_end

	// separator_beg
	ASTERISK //*
	TILDE    //~
	// separator_end
)

var tokens = [...]string{
	ILLEGAL:  "ILLEGAL",
	IDENT:    "IDENT",
	EOF:      "EOF",
	WS:       " ",
	ISA:      "ISA",
	GS:       "GS",
	ST:       "ST",  // Transaction header
	W12:      "W12", // Date information
	N9:       "N9",  // Batch information
	SE:       "SE",  // Transaction footer
	GE:       "GE",  // Record count
	IEA:      "IEA",
	ASTERISK: "*",
	TILDE:    "~",
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

// var separators map[string]EdiToken

func init() {
	keywords = make(map[string]EdiToken)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}

	// separators = make(map[string]EdiToken)
	// for i := separator_beg + 1; i < separator_end; i++ {
	// 	separators[tokens[i]] = i
	// }
}

func LookupToken(ident string) EdiToken {
	if tok, is_keyword := keywords[strings.ToUpper(ident)]; is_keyword {
		return tok
	}
	return ILLEGAL
}
