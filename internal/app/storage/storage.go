package storage

//go:generate mockery --name=Storage --testonly --inpackage --exported

import (
	"sync"

	"github.com/paveltyukin/practicum-go-shortener/pkg/logger"
)

var _ Storage = &storage{}

type Storage interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

type storage struct {
	logger *logger.Logger
	mx     sync.RWMutex
	values map[string]string
}

func (s *storage) Get(key string) (string, bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	value, ok := s.values[key]
	return value, ok
}

func (s *storage) Set(key, value string) {
	s.mx.Lock()
	s.values[key] = value
	s.mx.Unlock()
}

func New(logger *logger.Logger) Storage {
	return &storage{
		values: make(map[string]string),
		logger: logger,
	}
}
