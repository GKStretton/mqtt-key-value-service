package storage

type Key string
type Value []byte

type KeyValueStore interface {
	Set(Key, Value) error
	Get(Key) (Value, error)
}

func NewMemoryStore() KeyValueStore {
	return &MemoryStore{}
}

func NewFileStore(dir string) KeyValueStore {
	return &FileStore{
		dir: dir,
	}
}
