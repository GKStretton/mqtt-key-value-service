package storage

import "fmt"

type FileStore struct {
	dir string
}

func (s *FileStore) Set(Key, Value) error {
	return fmt.Errorf("Set not implemented")
}

func (s *FileStore) Get(Key) (Value, error) {
	return nil, fmt.Errorf("get not implemented")
}
