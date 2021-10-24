package memory

import (
	"testing"

	"github.com/khan745/gokvdb/internal/pkg/storage"
	"github.com/stretchr/testify/assert"
)

func TestStorage_Put(t *testing.T) {
	type args struct {
		key    storage.Key
		setter storage.ValueSetter
	}
	tests := []struct {
		name    string
		strg    *Storage
		args    args
		wantErr bool
	}{
		{"Test1", New(), args{key: "Put", setter: storage.ValueSetter(func(*storage.Value) (*storage.Value, error) {
			val := storage.NewString("value")
			return val, nil
		})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.strg.Put(tt.args.key, tt.args.setter); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Put_WhenValueShouldBeAdded(t *testing.T) {
	strg := &Storage{
		items: make(map[storage.Key]*storage.Value),
	}

	assert.NoError(t, strg.Put(storage.Key("key"), func(*storage.Value) (*storage.Value, error) {
		val := storage.NewString("value")
		return val, nil
	}))

	_, ok := strg.items[storage.Key("key")]
	assert.True(t, ok)
}
