package main

import (
	"log"

	"github.com/paveltyukin/practicum-go-shortener/internal/app/server"
	"github.com/paveltyukin/practicum-go-shortener/internal/app/shortener"
	"github.com/paveltyukin/practicum-go-shortener/internal/app/storage"
	"github.com/paveltyukin/practicum-go-shortener/pkg/logger"
)

func main() {
	l := logger.InitLogger()
	s := shortener.New(l)
	st := storage.New(l)
	err := server.Serve(s, st)
	if err != nil {
		log.Fatal(err)
	}
}
