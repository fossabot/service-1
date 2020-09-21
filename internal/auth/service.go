package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth/model"
	"github.com/perfolio/service/internal/auth/repository"
	"go.uber.org/zap"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (model.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	ChangeEmail(ctx context.Context, id uuid.UUID, newEmail string) error
	ConfirmEmail(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repository repository.Repository
	logger     *zap.Logger
}

func NewService(repository repository.Repository, logger *zap.Logger) Service {
	return &service{repository, logger}
}

func (s *service) CreateUser(ctx context.Context, email string, password string) (model.User, error) {

	user, err := model.NewUser(email, password)
	if err != nil {
		s.logger.Error("error", zap.Error(err))
		return model.User{}, err

	}
	err = s.repository.CreateUser(ctx, *user)
	if err != nil {
		s.logger.Error("error", zap.Error(err))
		return model.User{}, err

	}

	return *user, err
}

func (s *service) DeleteUser(ctx context.Context, id uuid.UUID) error {

	err := s.repository.DeleteUser(ctx, id)
	if err != nil {
		s.logger.Error("error", zap.Error(err))

	}

	return err
}

func (s *service) ChangeEmail(ctx context.Context, id uuid.UUID, newEmail string) error {

	err := s.repository.ChangeEmail(ctx, id, newEmail)
	if err != nil {
		s.logger.Error("error", zap.Error(err))

	}

	return err
}

func (s *service) ConfirmEmail(ctx context.Context, id uuid.UUID) error {

	err := s.repository.ConfirmEmail(ctx, id)
	if err != nil {
		s.logger.Error("error", zap.Error(err))

	}

	return err
}
