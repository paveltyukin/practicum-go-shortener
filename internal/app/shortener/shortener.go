package shortener

//go:generate mockery --name=Shortener --inpackage --testonly

import (
	"strconv"

	"github.com/paveltyukin/practicum-go-shortener/pkg/logger"
)

var _ Shortener = &shortener{}

type Shortener interface {
	Short(link string) string
}

type shortener struct {
	logger *logger.Logger
}

func (s *shortener) Short(link string) string {
	linkLength := len(link)
	return strconv.Itoa(linkLength)
}

func New(logger *logger.Logger) Shortener {
	return &shortener{
		logger: logger,
	}
}
