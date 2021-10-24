package command

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

type Parser struct {
	strg dataStore
}

//NewParser creates a new parser
func NewParser(strg dataStore) *Parser {
	return &Parser{strg: strg}
}

func (p *Parser) Parse(str string) (Command, []string, error) {
	var cmd Command
	args := p.extractArgs(str)
	fmt.Println(args)
	switch strings.ToUpper(args[0]) {
	case "SET":
		cmd = &Set{strg: p.strg}
	case "GET":
		cmd = &Get{strg: p.strg}
	default:
		return nil, nil, ErrCommandNotFound
	}

	return cmd, args[1:], nil
}

func (p *Parser) extractArgs(val string) []string {
	args := make([]string, 0)
	var inQuote bool
	var buf bytes.Buffer
	for _, r := range val {
		switch {
		case r == '"':
			inQuote = !inQuote
		case unicode.IsSpace(r):
			if !inQuote && buf.Len() > 0 {
				args = append(args, buf.String())
				buf.Reset()
			} else {
				buf.WriteRune(r)
			}
		default:
			buf.WriteRune(r)
		}
	}
	if buf.Len() > 0 {
		args = append(args, buf.String())
	}
	return args
}
