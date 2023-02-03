package main

import (
	"log"

	"github.com/paveltyukin/practicum-go-shortener/internal/app/server"
	"github.com/paveltyukin/practicum-go-shortener/internal/app/shortener"
	"github.com/paveltyukin/practicum-go-shortener/internal/app/storage"
	"github.com/paveltyukin/practicum-go-shortener/internal/config"
	"github.com/paveltyukin/practicum-go-shortener/pkg/logger"
)

func main() {
	l := logger.InitLogger()
	cfg := config.InitConfig(l)
	s := shortener.New(l)
	st := storage.New(l)
	err := server.Serve(cfg, s, st)
	if err != nil {
		log.Fatal(err)
	}
}
