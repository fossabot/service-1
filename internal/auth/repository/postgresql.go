package repository

import (
	"context"

	"errors"
	"github.com/google/uuid"
	"github.com/perfolio/service/internal/auth/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	ERR_EMAIL_NOT_CHANGED = errors.New("New email must not be the same as the old email")
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

	if user.Email == newEmail {
		p.logger.Error(ERR_EMAIL_NOT_CHANGED.Error(), zap.String("id", id.String()))
		return ERR_EMAIL_NOT_CHANGED
	}

	err = p.db.Model(&user).Updates(map[string]interface{}{"Email": newEmail, "ConfirmedAt": nil}).Error
	if err != nil {
		p.logger.Error("Could not update user", zap.String("id", id.String()))
		return err
	}
	return nil

}

func (p *postgres) ConfirmEmail(ctx context.Context, id uuid.UUID) error {

	now := time.Now()

	err := p.db.Model(&model.User{}).Where("ID = ? ", id).Update("ConfirmedAt", &now).Error
	if err != nil {
		p.logger.Error("Could not update user", zap.String("id", id.String()))
		return err
	}
	return nil

}
