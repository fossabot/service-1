package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth"
)

type ChangeEmailRequest struct {
	ID       uuid.UUID `json:"id"`
	NewEmail string    `json:"newEmail"`
}
type ChangeEmailResponse struct {
	Message string `json:"message"`
}

func makeChangeEmailEndpoint(service auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(ChangeEmailRequest)

		err := service.ChangeEmail(ctx, request.ID, request.NewEmail)
		if err != nil {
			return ChangeEmailResponse{Message: err.Error()}, err
		}
		return ChangeEmailResponse{Message: "Updated email"}, err
	}
}
