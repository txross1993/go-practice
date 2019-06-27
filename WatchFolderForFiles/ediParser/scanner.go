package ediParser

import (
	"bufio"
	"bytes"
	"io"
)

var eof = rune(0)

func isTilde(ch rune) bool {
	return ch == '~'
}

func isAsterisk(ch rune) bool {
	return ch == '*'
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isIdent(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || (ch == '-') || isWhitespace(ch) || (ch == '.')
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	s.r.UnreadRune()
}

func (s *Scanner) scanAsterisks() (tok EdiToken, lit string) {
	// Create a buffer and read the current character into it.package ediParser

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent asterisk character into the buffer.
	// Non-asterisks and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isAsterisk(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return ASTERISK, buf.String()
}

func (s *Scanner) scanIdent() (tok EdiToken, lit string) {
	// Create a buffer and read the current character into it.package ediParser

	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent non-asterisk character into the buffer.
	// Non-idents and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isIdent(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword, return that keyword.
	kw := LookupToken(buf.String())

	if kw != ILLEGAL {
		return kw, buf.String()
	} else {
		return IDENT, buf.String()
	}
}

func (s *Scanner) Scan() (tok EdiToken, lit string) {
	// Read the next rune
	ch := s.read()

	// Consume contiguous asterisks
	// If we see a letter or number, consume it as a keyword or identifier
	if isAsterisk(ch) {
		s.unread()
		return s.scanAsterisks()
	} else if isTilde(ch) {
		return TILDE, "~"
	} else if isIdent(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the character
	switch ch {
	case eof:
		return EOF, ""
	}

	return ILLEGAL, string(ch)
}
