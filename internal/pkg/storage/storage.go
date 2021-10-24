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

func (v *Value) Data() interface{} {
	return v.data
}

func (v *Value) Type() DataType {
	return v.dataType
}

func NewString(str string) *Value {
	return &Value{
		data:     str,
		dataType: StringDataType,
	}
}

//NewBitMap creates a new value of the BitMapDataType. Stored as uint64 integer.
func NewBitMap(value []uint64) *Value {
	return &Value{
		data:     value,
		dataType: BitMapDataType,
	}
}

//NewList creates a new value of the ListDataType. Stored as slice of strings.
func NewList(data []string) *Value {
	return &Value{
		data:     data,
		dataType: ListDataType,
	}
}

//NewMap creates a new value of the MapDataType.
func NewMap(val map[string]string) *Value {
	return &Value{
		data:     val,
		dataType: MapDataType,
	}
}
