package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/perfolio/service/internal/auth"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func makeCreateUserEndpoint(service auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(CreateUserRequest)

		return service.CreateUser(ctx, request.Email, request.Password)
	}
}
