package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth/model"
	"go.uber.org/zap"
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

func (p *postgres) ChangeEmail(ctx context.Context, id uuid.UUID, newEmail string) error {
	var user model.User
	err := p.db.First(&user, id).Error
	if err != nil {
		p.logger.Error("Could not find user in db", zap.String("id", id.String()))
		return err
	}

	err = p.db.Model(&user).Updates(model.User{Email: newEmail, ConfirmedAt: nil}).Error
	if err != nil {
		p.logger.Error("Could not update user", zap.String("id", id.String()))
		return err
	}
	return nil

}
