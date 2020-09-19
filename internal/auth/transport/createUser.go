package transport

import (
	"context"
	"encoding/json"
	"github.com/perfolio/service/internal/auth/endpoint"
	"net/http"
)

func DecodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
