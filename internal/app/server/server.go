package server

import (
	"io"
	"net/http"
	"time"
)

type handler struct {
	storage   storage
	shortener shortener
}

type storage interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

type shortener interface {
	Short(link string) string
}

// postHandler /
func (h handler) postHandler(w http.ResponseWriter, r *http.Request) {
	rawLink, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	link := string(rawLink)
	shortLink := h.shortener.Short(link)
	h.storage.Set(shortLink, link)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("http://localhost:8080/" + shortLink))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// getHandler /{string id}
func (h handler) getHandler(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Path[1:]
	link, ok := h.storage.Get(shortLink)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Location", link)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getHandler(w, r)
	case "POST":
		h.postHandler(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func Serve(addr string, s shortener, st storage) error {
	server := &http.Server{
		Addr:           addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	h := handler{
		shortener: s,
		storage:   st,
	}

	http.Handle("/", h)

	return server.ListenAndServe()
}
