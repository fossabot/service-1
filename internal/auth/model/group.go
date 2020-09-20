package model

import (
	"github.com/google/uuid"
	"time"
)

type Group struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name string `json:"name" gorm:"unique"`
}




func NewGroup(name string, uuid UUID, time Time) (Group, error) {
	return Group{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}, nil
}
