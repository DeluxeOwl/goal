package goal

import (
	"fmt"
	"io"
	"strings"
)

// Token represents a token information.
type Token struct {
	Type TokenType // token type
	Rune rune      // context text when only one rune is enough
	Pos  int       // token's offset in the source
	Text string    // content text (identifier, string, number)
}

func (t Token) String() string {
	switch t.Type {
	case ERROR:
		return t.Text
	case ADVERB, IDENT, VERB, NUMBER:
		return t.Text
	case LEFTBRACE:
		return "{"
	case LEFTBRACKET:
		return "["
	case LEFTPAREN:
		return "("
	case RIGHTBRACE:
		return "}"
	case RIGHTBRACKET:
		return "]"
	case RIGHTPAREN:
		return ")"
	case SEMICOLON:
		return ";"
	case STRING:
		return "\"" + t.Text + "\""
	default:
		return t.Type.String()
	}
}

// TokenType represents the different kinds of tokens.
type TokenType int32

// These constants describe the possible kinds of tokens.
const (
	NONE TokenType = iota
	EOF
	ERROR
	ADVERB
	IDENT
	LEFTBRACE
	LEFTBRACKET
	LEFTPAREN
	NEWLINE
	NUMBER
	RIGHTBRACE
	RIGHTBRACKET
	RIGHTPAREN
	SEMICOLON
	STRING
	VERB
)

// Scanner represents the state of the scanner.
type Scanner struct {
	reader  *strings.Reader // rune reader
	err     error           // scanning error (if any)
	peeked  bool            // peeked next
	npos    int             // position of next rune in the input
	tpos    int             // current token start position
	pr      rune            // peeked rune
	psize   int             // size of last peeked rune
	start   bool            // at line start
	exprEnd bool            // at expression start
	token   Token           // last token
	source  string          // source string
}

type stateFn func(*Scanner) stateFn

// Init initializes the scanner with a given reader. It can be reused again
// with a new reader, but position information will be reset.
func (s *Scanner) Init(source string) {
	if s.reader != nil {
		*s = Scanner{}
	}
	s.source = source
	s.reader = strings.NewReader(source)
	s.start = true
	s.token = Token{Type: EOF, Pos: 0}
}

// Next produces the next token from the input reader.
func (s *Scanner) Next() Token {
	state := scanAny
	for {
		state = state(s)
		if state == nil {
			return s.token
		}
	}
}

const eof = -1

func (s *Scanner) peek() rune {
	if s.peeked {
		return s.pr
	}
	r, size, err := s.reader.ReadRune()
	if err != nil {
		if err != io.EOF {
			s.err = err
		}
		r = eof
	}
	s.peeked = true
	s.pr = r
	s.psize = size
	return s.pr
}

func (s *Scanner) next() rune {
	if s.peeked {
		s.peeked = false
		s.npos += s.psize
		return s.pr
	}
	r, sz, err := s.reader.ReadRune()
	s.npos += sz
	if err != nil {
		if err != io.EOF {
			s.err = err
		}
		return eof
	}
	//fmt.Printf("[%c]", r)
	return r
}

func (s *Scanner) emit(t TokenType) stateFn {
	s.token = Token{Type: t, Pos: s.tpos}
	s.start = t == NEWLINE
	switch t {
	case LEFTBRACE, LEFTBRACKET, LEFTPAREN, NEWLINE, SEMICOLON, EOF:
		// all of these don't have additional content, so we don't do
		// this test in the other emits.
		s.exprEnd = false
	default:
		s.exprEnd = true
	}
	return nil
}

func (s *Scanner) emitError(err string) stateFn {
	s.token = Token{Type: ERROR, Pos: s.tpos, Text: err}
	return nil
}

func (s *Scanner) emitString(t TokenType) stateFn {
	s.start = false
	s.token = Token{Type: t, Pos: s.tpos, Text: s.source[s.tpos:s.npos]}
	s.exprEnd = true
	return nil
}

func (s *Scanner) emitOp(t TokenType, r rune) stateFn {
	s.start = false
	s.token = Token{Type: t, Pos: s.tpos, Rune: r}
	s.exprEnd = false
	return nil
}

func (s *Scanner) emitEOF() stateFn {
	if s.err != nil {
		return s.emitError(s.err.Error())
	}
	return s.emit(EOF)
}

func scanAny(s *Scanner) stateFn {
	s.tpos = s.npos
	r := s.next()
	switch r {
	case eof:
		return s.emitEOF()
	case '\n':
		return s.emit(NEWLINE)
	case ' ', '\t':
		return scanSpace
	case '/':
		if s.start {
			return scanCommentLine
		}
		return s.emitOp(ADVERB, r)
	case '\'', '\\':
		return s.emitOp(ADVERB, r)
	case '{':
		return s.emit(LEFTBRACE)
	case '[':
		return s.emit(LEFTBRACKET)
	case '(':
		return s.emit(LEFTPAREN)
	case '}':
		return s.emit(RIGHTBRACE)
	case ']':
		return s.emit(RIGHTBRACKET)
	case ')':
		return s.emit(RIGHTPAREN)
	case ';':
		return s.emit(SEMICOLON)
	case '-':
		if !s.exprEnd {
			return scanMinus
		}
		return s.emitOp(VERB, r)
	case ':', '+', '*', '%', '!', '&', '|', '<', '>',
		'=', '~', ',', '^', '#', '_', '$', '?', '@', '.':
		return s.emitOp(VERB, r)
	case '"':
		return scanString
	case '`':
		return scanSymbolString
	}
	switch {
	case isDigit(r):
		return scanNumber
	case isAlpha(r):
		return scanIdent
	default:
		return s.emitError(fmt.Sprintf("unexpected character: %c", r))
	}
	return nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isAlpha(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

func isAlphaNum(r rune) bool {
	return isAlpha(r) || isDigit(r)
}

func scanSpace(s *Scanner) stateFn {
	for {
		r := s.peek()
		switch r {
		case eof:
			return scanAny
		case '/':
			s.next()
			return scanComment
		case ' ', '\t':
			s.next()
		case '-':
			s.tpos = s.npos
			s.next()
			return scanMinus
		default:
			return scanAny
		}
	}
}

func scanComment(s *Scanner) stateFn {
	for {
		r := s.next()
		switch r {
		case eof:
			return s.emitEOF()
		case '\n':
			return s.emit(NEWLINE)
		}
	}
}

func scanCommentLine(s *Scanner) stateFn {
	r := s.peek()
	if r == '\n' {
		return scanMultiLineComment
	}
	return scanComment
}

func scanMultiLineComment(s *Scanner) stateFn {
	for {
		r := s.next()
		switch {
		case r == eof:
			return s.emitEOF()
		case r == '\\' && s.start:
			r := s.next()
			if r == '\n' {
				return scanAny
			}
		}
	}
}

func scanString(s *Scanner) stateFn {
	for {
		// XXX: catch invalid newline here?
		r := s.next()
		switch r {
		case eof:
			return s.emitError("non terminated string: unexpected EOF")
		case '\\':
			nr := s.peek()
			if nr == '"' {
				s.next()
			}
		case '"':
			return s.emitString(STRING)
		}
	}
}

func scanSymbolString(s *Scanner) stateFn {
	for {
		r := s.next()
		switch r {
		case eof:
			return s.emitError("non terminated string: unexpected EOF")
		case '`':
			return s.emitString(STRING)
		}
	}
}

func scanNumber(s *Scanner) stateFn {
	for {
		r := s.peek()
		switch {
		case r == eof:
			return s.emitString(NUMBER)
		case r == '.':
			s.next()
		case r == 'e':
			s.next()
			r = s.peek()
			if r == '+' || r == '-' {
				s.next()
				return scanExponent
			}
		case !isAlphaNum(r):
			return s.emitString(NUMBER)
		default:
			s.next()
		}
	}
}

func scanExponent(s *Scanner) stateFn {
	for {
		r := s.peek()
		switch {
		case r == eof:
			return s.emitString(NUMBER)
		case !isDigit(r):
			return s.emitString(NUMBER)
		default:
			s.next()
		}
	}
}

func scanMinus(s *Scanner) stateFn {
	r := s.peek()
	if isDigit(r) {
		return scanNumber
	}
	return s.emitOp(VERB, '-')
}

func scanIdent(s *Scanner) stateFn {
	for {
		r := s.peek()
		switch {
		case r == eof:
			return s.emitString(IDENT)
		case r == '.':
			r = s.peek()
			if !isAlpha(r) {
				return s.emitString(IDENT)
			}
			s.next()
		case isAlphaNum(r):
			s.next()
		default:
			return s.emitString(IDENT)
		}
	}
}
