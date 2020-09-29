package storage

import "errors"

var (
	//ErrKeyNotExists means that key does not exist. Returns by Key method
	ErrKeyNotExists = errors.New("storage: key does not exist")
)

type DataType string

type Key struct {
	value    string
	dataType DataType
}

func (k Key) Val() string {
	return k.value
}

type Value interface{}

type Storage interface {
	Put(Key, Value) error
	Get(Key) (Value, error)
	Del(Key) error
	GetKey(string) (Key, error)
	Keys() ([]Key, error)
}
