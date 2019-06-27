package ediParser

import (
	"bufio"
	"bytes"
	"io"
)

var eof = rune(0)

// func isTilde(ch rune) bool {
// 	return ch == '~'
// }

func isAsterisk(ch rune) bool {
	return ch == '*'
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isNewLine(ch rune) bool {
	return ch == '\n' || ch == '\r'
}

func isIdent(ch rune) bool {
	//Asterisk value: 42
	// NL value: 10
	// CR value: 13
	// WS value: 32
	// Tab value: 9

	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || (ch == '-') || isWhitespace(ch) || (ch == '.') || (ch == '>')
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
	// Create a buffer and read the current character into it.
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
	li.Debug("Scanned asterisk string: ", buf.String())
	return ASTERISK, buf.String()
}

// func (s *Scanner) scanNewLine() (tok EdiToken, lit string) {
// 	// Create a buffer and read the current character into it.

// 	var buf bytes.Buffer
// 	buf.WriteRune(s.read())

// 	// Read every subsequent newline character into the buffer.
// 	// Non-newlines and EOF will cause the loop to exit.
// 	for {
// 		if ch := s.read(); ch == eof {
// 			break
// 		} else if !isNewLine(ch) {
// 			s.unread()
// 			break
// 		} else {
// 			buf.WriteRune(ch)
// 		}
// 	}
// 	li.Debug("Scanned new line string: ", buf.String())
// 	return NL, buf.String()
// }

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

	li.Debug("Scanned identity string: ", buf.String())

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
	} else if isIdent(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the character
	if isNewLine(ch) {
		return NL, "\\n"
	}

	switch ch {
	case eof:
		return EOF, ""
	default:
		return ILLEGAL, string(ch)
	}

}
