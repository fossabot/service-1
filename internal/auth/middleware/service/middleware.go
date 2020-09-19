package service

import (
	"github.com/perfolio/service/internal/auth"
)

// Middleware describes a service middleware.
type Middleware func(auth.Service) auth.Service
