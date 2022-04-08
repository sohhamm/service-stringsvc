package main

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

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

// * Request and Responses (Endpoint layer)
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't json unmarshal so we use strings
}

type countRequest struct {
	S string `json:"s"`
}
type countResponse struct {
	V int `json:"v"`
}

// ? go kit endpoint function signature
// type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
// Adapter will take the StringService interface and returns and endpoint corresponding to that one method

// * Adapters

func makeUpperCaseEndpoint(svc StringService) endpoint.Endpoint {

	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}

		return uppercaseResponse{v, ""}, nil

	}

}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {

	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}

}

func main() {
	log.Println("service stringsvc")
}
