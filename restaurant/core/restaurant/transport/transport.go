package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"restaurant/core/restaurant/endpoint"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Handle("/restaurant", kitHttp.NewServer(
		endpoints.Create,
		decodeCreateRestaurantRequest,
		encodeResponse,
	)).Methods("POST")
	r.Handle("/restaurant/{id}", kitHttp.NewServer(
		endpoints.Get,
		decodeGetRestaurantRequest,
		encodeResponse,
	)).Methods("GET")
	r.Handle("/restaurant/{id}", kitHttp.NewServer(
		endpoints.Update,
		decodeUpdateRestaurantRequest,
		encodeResponse,
	)).Methods("PUT")
	r.Handle("/restaurant/{id}", kitHttp.NewServer(
		endpoints.Delete,
		decodeDeleteRestaurantRequest,
		encodeResponse,
	)).Methods("DELETE")
	return r
}

func decodeCreateRestaurantRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req.Restaurant)
	return req, err
}

func decodeGetRestaurantRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.GetRequest{ID: vars["id"]}, nil
}

func decodeUpdateRestaurantRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateRequest
	vars := mux.Vars(r)
	req.ID = vars["id"]
	err := json.NewDecoder(r.Body).Decode(&req.Restaurant)
	return req, err
}

func decodeDeleteRestaurantRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return endpoint.DeleteRequest{ID: vars["id"]}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
