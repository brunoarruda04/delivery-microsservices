// internal/restaurant/endpoint/endpoint.go
package endpoint

import (
	"context"
	"restaurant/core/restaurant/domain"
	"restaurant/core/restaurant/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Create endpoint.Endpoint
	Get    endpoint.Endpoint
	Update endpoint.Endpoint
	Delete endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.Create(req.Restaurant)
		return CreateResponse{ID: id, Err: err}, err
	}
}

func makeGetEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		restaurant, err := s.Get(req.ID)
		return GetResponse{Restaurant: restaurant, Err: err}, err
	}
}

func makeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		err := s.Update(req.ID, req.Restaurant)
		return UpdateResponse{Err: err}, err
	}
}

func makeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		err := s.Delete(req.ID)
		return DeleteResponse{Err: err}, err
	}
}

type CreateRequest struct {
	Restaurant domain.Restaurant
}

type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

type GetRequest struct {
	ID string
}

type GetResponse struct {
	Restaurant domain.Restaurant `json:"restaurant"`
	Err        error             `json:"error,omitempty"`
}

type UpdateRequest struct {
	ID         string
	Restaurant domain.Restaurant
}

type UpdateResponse struct {
	Err error `json:"error,omitempty"`
}

type DeleteRequest struct {
	ID string
}

type DeleteResponse struct {
	Err error `json:"error,omitempty"`
}
