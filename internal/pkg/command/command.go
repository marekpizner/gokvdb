package command

import (
	"errors"

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

type Command interface {
	Name() string

	Help() string

	Execute(args ...string) Result
}

type dataStore interface {
	//Put puts a new value at the given Key.
	Put(key storage.Key, sttr storage.ValueSetter) error
	//Get gets a value by the given key.
	Get(key storage.Key) (*storage.Value, error)
	//Del deletes a value by the given key.
	Del(key storage.Key) error
	//Keys returns all stored keys.
	Keys() ([]storage.Key, error)
}

type commandParser interface {
	Parse(str string) (cmd Command, args []string, err error)
}
