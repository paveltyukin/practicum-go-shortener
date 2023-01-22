package storage

import (
	"sync"
)

type Storage struct {
	mx     sync.RWMutex
	values map[string]string
}

func (s *Storage) Get(key string) (string, bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	value, ok := s.values[key]
	return value, ok
}

func (s *Storage) Set(key, value string) {
	s.mx.Lock()
	s.values[key] = value
	s.mx.Unlock()
}

func New() *Storage {
	return &Storage{
		values: make(map[string]string),
	}
}
