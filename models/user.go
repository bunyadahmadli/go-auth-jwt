package models

import (
	"gorm.io/gorm"
)

type UserRole int

const (
	SystemAdmin UserRole = 1
	BasicUser            = 2
)

type User struct {
	gorm.Model
	Name     string   `gorm:"not null" json:"name,omitempty"`
	Surname  string   `gorm:"unique not null" gorm:"not null" json:"surname,omitempty"`
	Email    string   `gorm:"unique not null" json:"email,omitempty"`
	Password string   `gorm:"not null" json:"password,omitempty"`
	Role     UserRole `gorm:"not null" json:"role,string,omitempty"`
}
