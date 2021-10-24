package command

import (
	"strings"

	"github.com/khan745/gokvdb/internal/pkg/storage"
)

type Set struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Set) Name() string {
	return "SET"
}

//Help implements Help of Command interface
func (c *Set) Help() string {
	return `Usage: SET key value
Set key to hold the string value.
If key already holds a value, it is overwritten.`
}

//Execute implements Execute of Command interface
func (c *Set) Execute(args ...string) Result {
	if len(args) != 2 {
		return ErrResult{ErrWrongArgsNumber}
	}

	value := strings.Join(args[1:], " ")

	setter := func(old *storage.Value) (*storage.Value, error) {
		return storage.NewString(value), nil
	}

	if err := c.strg.Put(storage.Key(args[0]), setter); err != nil {
		return ErrResult{err}
	}
	return OkResult{}
}
