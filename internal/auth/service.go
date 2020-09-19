package auth

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/perfolio/service/internal/auth/model"
	"github.com/perfolio/service/internal/auth/repository"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (model.User, error)
}

type service struct {
	repository repository.Repository
	logger     log.Logger
}

func NewService(repository repository.Repository, logger log.Logger) Service {
	return service{repository, logger}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (model.User, error) {

	user, err := model.NewUser(email, password)
	err = s.repository.CreateUser(ctx, user)
	if err != nil {
		level.Error(s.logger).Log("err", err)
		return model.User{}, err

	}

	return user, err
}
