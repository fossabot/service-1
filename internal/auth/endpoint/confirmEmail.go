package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth"
)

type ConfirmEmailRequest struct {
	ID uuid.UUID `json:"id"`
}
type ConfirmEmailResponse struct {
	Message string `json:"message"`
}

func makeConfirmEmailEndpoint(service auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		request := req.(ConfirmEmailRequest)

		err := service.ConfirmEmail(ctx, request.ID)
		if err != nil {
			return ConfirmEmailResponse{Message: err.Error()}, err
		}
		return ConfirmEmailResponse{Message: "Updated email"}, err
	}
}
