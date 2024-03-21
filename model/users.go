package model

import (
	"github.com/gofrs/uuid"
)

type UserInfo interface {
	GetID() uuid.UUID
	GetDisplayName() string
}

// User userの構造体
type User struct {
	ID          uuid.UUID `gorm:"type:char(36);not null;primaryKey"`
	DisplayName string    `gorm:"type:varchar(32);not null;default:''"`
}

// GetID implements UserInfo interface
func (user *User) GetID() uuid.UUID {
	return user.ID
}

// GetDisplayName implements UserInfo interface
func (user *User) GetDisplayName() string {
	return user.DisplayName
}
