package storage

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		s := &storage{
			logger: nil,
			values: make(map[string]string),
		}

		st := New(nil)
		assert.Equal(t, s, st)
	})
}

func Test_storage_Get(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		s := New(nil)
		_, ok := s.Get("1")
		assert.False(t, ok, true)
	})

	tests := []struct {
		name         string
		setupStorage func(mock *MockStorage)
	}{
		{
			name: "simple positive test",
			setupStorage: func(mock *MockStorage) {
				mock.
					On("Get", "args").
					Return("1", true)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strg := NewMockStorage(t)
			if tt.setupStorage != nil {
				tt.setupStorage(strg)
			}

			result, ok := strg.Get("args")
			if !ok {
				require.Error(t, errors.New("failed"))
			}

			assert.Equal(t, "1", result)
		})
	}
}

func Test_storage_Set(t *testing.T) {
	t.Run("functional test", func(t *testing.T) {
		s := New(nil)
		s.Set("key", "value")
		result, ok := s.Get("key")
		assert.True(t, ok, true)
		assert.Equal(t, result, "value")
	})

	tests := []struct {
		name         string
		setupStorage func(mock *MockStorage)
	}{
		{
			name: "simple mocked test",
			setupStorage: func(mock *MockStorage) {
				mock.
					On("Set", "key", "value").
					Return(true)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strg := NewMockStorage(t)
			if tt.setupStorage != nil {
				tt.setupStorage(strg)
			}

			strg.Set("key", "value")
		})
	}
}
