package repository

import (
	"context"

	"go.uber.org/zap"
	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth/model"
	"gorm.io/gorm"
)

type postgres struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPostgres(db *gorm.DB, logger *zap.Logger) Repository {
	return &postgres{db: db, logger: logger}
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
