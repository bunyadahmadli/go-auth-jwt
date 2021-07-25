package models

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username string `gorm:"unique not null" gorm:"not null" json:"username"`
	Name     string	`gorm:"not null" json:"name"`
	Email    string `gorm:"unique not null" json:"email"`
	Password []byte  `gorm:"not null" json:"password"`
}

type UserRoles struct {
	gorm.Model
	userId uint `gorm:"not null" json:"user_id"`
	roleId uint `gorm:"not nul" json:"role_id"`
}

type Roles struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	ResourceId uint `gorm:"not null" json:"resource_id"`
	ResourceTypeId uint `gorm:"not null" json:"resource_type_id"`
}