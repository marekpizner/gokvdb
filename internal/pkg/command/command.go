package command

import (
	"bytes"
	"errors"
	"strings"
	"unicode"

	"github.com/khan745/gokvdb/internal/pkg/storage"
)

var (
	//ErrCommandNotFound means that command could not be parsed. Returns by Parse
	ErrCommandNotFound = errors.New("command: not found")
	//ErrWrongArgsNumber means that given arguments not acceptable by Command. Returns by Parse
	ErrWrongArgsNumber = errors.New("command: wrong args number")
	//ErrWrongTypeOp means that operation is not acceptable for the given key
	ErrWrongTypeOp = errors.New("command: wrong type operation")
)

var commands = make(map[string]Command)

type Command interface {
	Name() string

	Help() string

	ValidateArgs(args ...string) error

	Execute(strg storage.Storage, args ...string) Result
}

func extractArgs(val string) []string {
	args := make([]string, 0)
	var inQuote bool
	var buf bytes.Buffer
	for _, r := range []rune(val) {
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

func Parse(value string) (Command, []string, error) {
	args := extractArgs(value)

	cmd, ok := commands[strings.ToUpper(args[0])]
	if !ok {
		return nil, nil, ErrCommandNotFound
	}

	args = args[1:]

	if err := cmd.ValidateArgs(args...); err != nil {
		return nil, nil, err
	}

	return cmd, args, nil
}
