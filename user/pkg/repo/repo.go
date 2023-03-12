// Package repo provides an interface and implementation for a user repository.
// The repository uses GORM as the ORM library and supports the CRUD (Create, Read, Update, Delete) operations for User entity.
package repo
//go:generate mockgen -source=./repo.go -destination=../../mocks/repo_mock.go -package=mocks

import (
	"github.com/cyicz123/todolist/user/pkg/logger"

	"github.com/spf13/viper"
)

// User represents a user entity with the fields such as UserID, UserName, NickName, and PasswordDigest.
type User struct {
	UserID         uint      `gorm:"primarykey"`	// UserID represents the unique identifier for a user.
	UserName       string    `gorm:"unique"`		// UserName represents the unique username for a user.
	NickName       string							// NickName represents the nickname for a user.
	PasswordDigest string							// PasswordDigest represents the hashed password for a user.
}

// UserRepository provides an interface for a user repository with the methods such as Create, Update, Delete, GetByID, and GetByName.
type UserRepository interface {
	// Create creates a new user in the repository and returns an error if any occurred.
	Create(user *User) error

	// Update updates an existing user in the repository and returns an error if any occurred.
	Update(user *User) error

	// Delete deletes a user from the repository by the provided name and returns an error if any occurred.
	Delete(name string) error

	// GetByID retrieves a user from the repository by the provided user ID and returns the user and an error if any occurred.
	GetByID(id uint) (*User, error)

	// GetByName retrieves a user from the repository by the provided username and returns the user and an error if any occurred.
	GetByName(name string) (*User, error)
}

// DBFactory provides an interface for creating a UserRepository with the given logger and viper configuration.
type DBFactory interface {
	New(l logger.Interface, v *viper.Viper) (UserRepository,error)
}