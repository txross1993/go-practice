package ediParser

import (
	"fmt"
	"io"
)

type EdiFile []EdiStatement

type EdiStatement struct {
	Keyword string
	Fields  []string
}

type Parser struct {
	s   *Scanner
	buf struct {
		tok EdiToken // last read token
		lit string   // last read literal
		n   int      // buffer size (max=1)
	}
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

func (p *Parser) scanIgnoreAsterisk() (tok EdiToken, lit string) {
	tok, lit = p.scan()
	if tok == ASTERISK {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) Parse() (*EdiStatement, error) {
	var EdiFile []EdiStatement

	//TODO
	stmt := &EdiStatement{}

	tok, lit := p.scanIgnoreAsterisk()
	if LookupToken(lit) != IEA {
		return nil, fmt.Errorf("Found %v, expected IEA", lit)
	}

	for {
		// Read a field.
		tok, lit := p.scanIgnoreAsterisk()

		// If the next token is a ~ then break the loop.
		if tok, _ := p.scanIgnoreAsterisk(); tok == TILDE {
			break
		}
	}
}
