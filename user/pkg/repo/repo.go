package repo

import (
	"user/pkg/logger"
)

type User struct {
	UserID         uint      `gorm:"primarykey"`
	UserName       string    `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(name string) error
	GetByID(id uint) (*User, error)
	GetByName(name string) (*User, error)
}

type DBFactory interface {
	New(logger.Interface) (UserRepository,error)
}