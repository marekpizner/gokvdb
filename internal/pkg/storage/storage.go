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
	//BitMapDataType is the bitmap data type. Stored as int64 integer
	BitMapDataType DataType = "bitmap"
	//ListDataType is the list data type. Stored as slice of string
	ListDataType DataType = "list"
	//MapDataType is the hash map type. Stored as map[string][string]
	MapDataType DataType = "map"
)

func (dt DataType) String() string {
	return string(dt)
}

//Key represents key to value
type Key string

//Value represents a single value of a storage
type Value struct {
	data     interface{}
	dataType DataType
}

//Storage represents storage
type Storage interface {
	Put(key Key, setter ValueSetter) error
	Get(key Key) (*Value, error)
	Del(key Key) error
	Keys() ([]Key, error)
}

type ValueSetter func(old *Value) (new *Value, err error)
