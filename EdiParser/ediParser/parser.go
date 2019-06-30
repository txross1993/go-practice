package ediParser

import (
	"bytes"
	"fmt"
	"io"

	log "github.com/txross1993/go-practice/EdiParser/logwrapper"
)

var li *log.StandardLogger

type EdiFile struct {
	Segments []EdiSegment
}

type EdiField struct {
	Position int
	Literal  string
}

type EdiSegment struct {
	Keyword string
	Fields  []EdiField
}

func (s *EdiStatement) String() string {
	return fmt.Sprintf("Keyword: %v,\nFields: %v", s.Keyword, s.Fields)
}

type Parser struct {
	s   *Scanner
	buf struct {
		tok EdiToken // last read token
		lit string   // last read literal
		n   int      // buffer size (max=1)
	}
	elemDelim string
	segDelim  string
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok EdiToken, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise, read next token
	tok, lit = p.s.Scan()

	// Save it to buffer in case we unscan it later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

func (p *Parser) scanIgnoreElemDelims() (tok EdiToken, lit string) {
	tok, lit = p.scan()
	if lit == p.elemDelim {
		tok, lit = p.scan()
	}
	return
}

func isISA(tok EdiToken) bool {
	return tok == ISA
}

func isDelim(tok EdiToken) bool {
	return tok == DELIM
}

func isIdent(tok EdiToken) bool {
	return tok == IDENT
}

func isKeyword(lit string) bool {
	return LookupToken(lit) != ILLEGAL
}

func (p *Parser) findHeaderSetDelims(tok EdiToken) *EdiSegment {
	// Construct header segment
	headerSegment := &EdiSegment{
		Keyword: tok.String(),
		Fields:  []EdiField{},
	}

	tok, lit := p.scan()

	assert(tok==DELIM)
	p.elemDelim = lit

	for {
		var elementIndex = 1
		// Now scan for segment delimiter
		tok, lit := p.scanIgnoreElemDelims()

		// Since we're ignoring elemental delimiters, any token that comes through as DELIM may be a segment delimiter
		if isDelim(tok) {
			var buf bytes.Buffer
			buf.WriteString(lit)

			// Is the next token GS?
			tok, lit := p.scanIgnoreElemDelims()
			if tok == GS {
				p.segDelim, _ = buf.ReadString()
				p.unscan()
				break
			} else if isDelim(tok) {
				buf.WriteString(lit)
			}
		} else if isIdent(tok) {
			field := &EdiField{Position: elementIndex, Literal: lit}
			headerSegment.Fields = append(headerSegment.Fields, field)
			elementIndex += 1
		}
	}

	return headerSegment
}

func (p *Parser) Parse() (EdiFile, error) {
	tok, lit := p.scan()

	// Beginning of EDI should be ISA keyword
	if isISA(tok) {
		var file EdiFile
		headerSegment := findHeaderSetDelims(tok)
		EdiFile = append(file, headerSegment)
	}


	}

	p.unscan()

	for {
		// The beginning of the scan should be a keyword
		stmt := &EdiStatement{}

		tok, lit := p.scanIgnoreAsterisk()

		if isKeyword(lit) {
			// Token is a keyword
			stmt.Keyword = tok.String()
		} else {
			return nil, &KeywordError{lit}
		}

		for {
			// Read a field.
			// If the next token is not an indentifier, unscan and break
			if tok, lit := p.scanIgnoreAsterisk(); tok == ILLEGAL {
				continue
			} else if tok == IDENT {
				stmt.Fields = append(stmt.Fields, lit)
			} else { //tok is EOF or NL
				p.unscan()
				break
			}
		}

		if tok, _ := p.scanIgnoreAsterisk(); tok == NL {
			// New Line marks the end of the statement
			// Don't unscan New Line, just consume it.
			file = append(file, stmt)
		}

		if tok, _ := p.scanIgnoreAsterisk(); tok != EOF {
			p.unscan()
		} else {
			break
		}
	}

	return file, nil
}

func init() {
	li = log.NewLogger()
}
