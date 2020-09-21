package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ChangeEmail(ctx context.Context, id uuid.UUID, newEmail string) error
}
