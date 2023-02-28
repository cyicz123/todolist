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
	Delete(id int) error
	GetByID(id int) (*User, error)
	GetByName(name string) (*User, error)
}

type DBFactory interface {
	New(logger.Interface) (UserRepository,error)
}