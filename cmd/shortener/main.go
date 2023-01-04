package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"time"
)

type Links map[string]string

var l = make(Links)

func generateShortLink(link string) string {
	hash := md5.Sum([]byte(link))
	return hex.EncodeToString(hash[:])
}

// PostHandler /
func PostHandler(w http.ResponseWriter, r *http.Request) {
	rawLink, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	link := string(rawLink)
	shortLink := generateShortLink(link)
	l[shortLink] = link

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(shortLink))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetHandler /{string id}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Path[1:]
	link, ok := l[shortLink]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Location", link)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetHandler(w, r)
	case "POST":
		PostHandler(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/", Handler)

	server := &http.Server{
		Addr:           "127.0.0.1:8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
