package repository

import (
	"context"

	"github.com/perfolio/service/internal/auth/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) error
}
