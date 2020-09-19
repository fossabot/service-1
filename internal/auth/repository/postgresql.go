package repository

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
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
	err := p.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil

}

func (p *postgres) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := p.db.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}
	return nil

}
