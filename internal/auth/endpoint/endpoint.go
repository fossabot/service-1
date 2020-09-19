package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"github.com/perfolio/service/internal/auth"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
}

// New returns all endpoints with their middleware already configured
func New(srv auth.Service, logger log.Logger) Endpoints {
	return Endpoints{
		CreateUser: loggingMiddleware(log.With(logger, "method", "CreateUser"))(MakeCreateUserEndpoint(srv)),
	}
}
