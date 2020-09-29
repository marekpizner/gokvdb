package storage

import "errors"

var (
	//ErrKeyNotExists means that key does not exist. Returns by Key method
	ErrKeyNotExists = errors.New("storage: key does not exist")
)

type DataType string

const (
	//StringDataType is the string data type
	StringDataType DataType = "string"
)

//Key represents key to value
type Key struct {
	value    string
	dataType DataType
}

//Val returns value of key
func (k Key) Val() string {
	return k.value
}

//NewStringKey creates a new key with StringDataType
func NewStringKey(key string) Key {
	return Key{
		value:    key,
		dataType: StringDataType,
	}
}

//Value represents stored data
type Value interface{}

//Storage represents storage
type Storage interface {
	Put(Key, Value) error
	Get(Key) (Value, error)
	Del(Key) error
	GetKey(string) (Key, error)
	Keys() ([]Key, error)
}
