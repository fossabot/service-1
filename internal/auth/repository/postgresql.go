package repository

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/perfolio/service/internal/auth/model"
	"gorm.io/gorm"
)

type postgres struct {
	db     *gorm.DB
	logger log.Logger
}

func NewPostgres(db *gorm.DB, logger log.Logger) Repository {
	return &postgres{db: db, logger: log.With(logger, "repo", "postgres")}
}

func (p *postgres) CreateUser(ctx context.Context, user model.User) error {
	logger := log.With(p.logger, "method", "CreateUser")

	err := p.db.Create(&user).Error
	if err != nil {
		return err
	}
	level.Info(logger).Log("createdUseer", user.ID)

	return nil

}
