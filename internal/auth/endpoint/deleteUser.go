package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth"
)

type DeleteUserRequest struct {
	ID uuid.UUID `json:"id"`
}
type DeleteUserResponse struct {
	Message string `json:"message"`
}

func makeDeleteUserEndpoint(service auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(DeleteUserRequest)

		err := service.DeleteUser(ctx, request.ID)
		if err != nil {
			return DeleteUserResponse{Message: err.Error()}, err
		}
		return DeleteUserResponse{Message: "Deleted user"}, err
	}
}
