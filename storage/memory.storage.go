package storage

import "fmt"

type MemoryStore struct{}

func (s *MemoryStore) Set(Key, Value) error {
	return fmt.Errorf("Set not implemented")
}

func (s *MemoryStore) Get(Key) (Value, error) {
	return nil, fmt.Errorf("get not implemented")
}
