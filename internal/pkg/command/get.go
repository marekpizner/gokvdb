package command

import (
	"github.com/khan745/gokvdb/internal/pkg/storage"
)

//Get is the GET command
type Get struct {
	strg dataStore
}

//Name implements Name of Command interface
func (c *Get) Name() string {
	return "GET"
}

//Help implements Help of Command interface
func (c *Get) Help() string {
	return `Usage: GET key
Get the value by key.
If provided key does not exist NIL will be returned.`
}

//Execute implements Execute of Command interface
func (c *Get) Execute(args ...string) Result {
	if len(args) != 1 {
		return ErrResult{ErrWrongArgsNumber}
	}
	value, err := c.strg.Get(storage.Key(args[0]))
	if err != nil {
		if err == storage.ErrKeyNotExists {
			return NilReply{}
		}
		return ErrResult{err}
	}
	if value.Type() != storage.StringDataType {
		return ErrResult{ErrWrongTypeOp}
	}
	return StringReply{Value: value.Data().(string)}
}
