package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Email             string     `json:"email" gorm:"unique"`
	EncryptedPassword string     `json:"-"`
	InvitedAt         *time.Time `json:"invited_at,omitempty"`
	ConfirmedAt       *time.Time `json:"confirmed_at,omitempty"`
	LastSignInAt      *time.Time `json:"last_sign_in_at,omitempty"`
}

func NewUser(email string, password string) (User, error) {
	encryptedPassword, err := hashPassword(password)
	if err != nil {
		return User{}, err
	}
	return User{
		ID:                uuid.New(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Email:             email,
		EncryptedPassword: encryptedPassword,
	}, nil
}

// IsConfirmed checks if a user has already being
// registered and confirmed.
func (u *User) IsConfirmed() bool {
	return u.ConfirmedAt != nil
}

// hashPassword generates a hashed password from a plaintext string
func hashPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}

// Authenticate a user from a password
func (u *User) Authenticate(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	return err == nil
}
