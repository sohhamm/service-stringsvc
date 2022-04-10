package services

import (
	"errors"
	"strings"
)

var Svc = stringService{}

// * Service layer
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct {
}

var errEmpty = errors.New("Empty String")

func (stringService) Uppercase(s string) (string, error) {

	if s == "" {
		return "", errEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
