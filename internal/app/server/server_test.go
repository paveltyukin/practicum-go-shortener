package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handler_ServeHTTP(t *testing.T) {
	tests := []struct {
		name           string
		setupStorage   func(mock *mockStorage)
		setupShortener func(mock *mockShortener)
		assertResponse func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "simple positive test",
			setupStorage: func(mock *mockStorage) {
				mock.
					On("Get", "100").
					Return("123", true).
					Maybe()
				mock.
					On("Set", "123", "/counter/Counter/100").
					Return("123", true).
					Maybe()
			},
			setupShortener: func(mock *mockShortener) {
				mock.
					On("Short", "/counter/Counter/100").
					Return("123", true).
					Maybe()
			},
			assertResponse: func(t *testing.T, r *httptest.ResponseRecorder) {
				assert.Equal(t, "text/plain", r.Header().Get("Content-Type"))
				assert.Equal(t, http.StatusCreated, r.Code)
				rawLink, _ := io.ReadAll(r.Body)
				assert.Equal(t, "http://localhost:8080/123", string(rawLink))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strg := newMockStorage(t)
			if tt.setupStorage != nil {
				tt.setupStorage(strg)
			}

			shrtnr := newMockShortener(t)
			if tt.setupShortener != nil {
				tt.setupShortener(shrtnr)
			}

			h := handler{
				storage:   strg,
				shortener: shrtnr,
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/counter/Counter/100", strings.NewReader("/counter/Counter/100"))
			h.postHandler(w, r)
			tt.assertResponse(t, w)
		})
	}
}

func Test_handler_getHandler(t *testing.T) {
	tests := []struct {
		name           string
		setupStorage   func(mock *mockStorage)
		assertResponse func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "simple positive test",
			setupStorage: func(mock *mockStorage) {
				mock.
					On("Get", "100").
					Return("123", true)
			},
			assertResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				assert.Equal(t, "123", response.Header().Get("Location"))
				assert.Equal(t, http.StatusTemporaryRedirect, response.Code)
			},
		},
		{
			name: "simple negative test",
			setupStorage: func(mock *mockStorage) {
				mock.On("Get", "100").Return("", false)
			},
			assertResponse: func(t *testing.T, response *httptest.ResponseRecorder) {
				assert.Equal(t, "", response.Header().Get("Location"))
				assert.Equal(t, http.StatusNotFound, response.Code)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strg := newMockStorage(t)
			if tt.setupStorage != nil {
				tt.setupStorage(strg)
			}

			h := handler{
				storage: strg,
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/100", nil)
			h.getHandler(w, r)
			tt.assertResponse(t, w)
		})
	}
}

func Test_handler_postHandler(t *testing.T) {
	tests := []struct {
		name           string
		setupStorage   func(mock *mockStorage)
		setupShortener func(mock *mockShortener)
		assertResponse func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "simple positive test",
			setupStorage: func(mock *mockStorage) {
				mock.
					On("Set", "123", "/counter/Counter/100").
					Return("123", true)
			},
			setupShortener: func(mock *mockShortener) {
				mock.
					On("Short", "/counter/Counter/100").
					Return("123", true)
			},
			assertResponse: func(t *testing.T, r *httptest.ResponseRecorder) {
				assert.Equal(t, "text/plain", r.Header().Get("Content-Type"))
				assert.Equal(t, http.StatusCreated, r.Code)
				rawLink, _ := io.ReadAll(r.Body)
				assert.Equal(t, "http://localhost:8080/123", string(rawLink))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			strg := newMockStorage(t)
			if tt.setupStorage != nil {
				tt.setupStorage(strg)
			}

			shrtnr := newMockShortener(t)
			if tt.setupShortener != nil {
				tt.setupShortener(shrtnr)
			}

			h := handler{
				storage:   strg,
				shortener: shrtnr,
			}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/counter/Counter/100", strings.NewReader("/counter/Counter/100"))
			h.postHandler(w, r)
			tt.assertResponse(t, w)
		})
	}
}
