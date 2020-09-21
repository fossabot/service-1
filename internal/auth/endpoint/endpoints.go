package endpoint

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/perfolio/service/internal/auth"
)

type Endpoints struct {
	CreateUser   endpoint.Endpoint
	DeleteUser   endpoint.Endpoint
	ChangeEmail  endpoint.Endpoint
	ConfirmEmail endpoint.Endpoint
}

// New returns all endpoints with their middleware already configured
func New(srv auth.Service) Endpoints {
	return Endpoints{
		CreateUser:   makeCreateUserEndpoint(srv),
		DeleteUser:   makeDeleteUserEndpoint(srv),
		ChangeEmail:  makeChangeEmailEndpoint(srv),
		ConfirmEmail: makeConfirmEmailEndpoint(srv),
	}
}
