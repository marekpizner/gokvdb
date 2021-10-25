package command

import "github.com/khan745/gokvdb/internal/pkg/storage"

type Del struct {
	strg dataStore
}

func (c *Del) Name() string {
	return "DEL"
}

func (c *Del) Help() string {
	return `Usage: DEL key
	Del the given key`
}

func (c *Del) ValidateArgs(args ...string) error {
	if len(args) != 1 {
		return ErrWrongArgsNumber
	}
	return nil
}

//Execute implements Execute of Command interface
func (c *Del) Execute(strg storage.Storage, args ...string) Response {
	err := strg.Del(storage.Key(args[0]))
	if err != nil {
		return ErrResult{err}
	}
	return OkResult{}
}
