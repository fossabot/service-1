package server

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/perfolio/service/internal/auth/endpoint"
	"github.com/perfolio/service/internal/auth/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func CreateHandler(endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(jsonHeader)

	r.Methods("POST").Path("/user/create").Handler(
		httptransport.NewServer(
			endpoints.CreateUser,
			transport.DecodeCreateUserRequest,
			transport.EncodeResponse,
		),
	)
	r.Methods("POST").Path("/user/delete").Handler(
		httptransport.NewServer(
			endpoints.DeleteUser,
			transport.DecodeDeleteUserRequest,
			transport.EncodeResponse,
		),
	)
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	return r
}
