package handler
//go:generate mockgen -source=./user.go -destination=../../mocks/handler_mock.go -package=mocks

import (
	"regexp"

	"github.com/cyicz123/todolist/user/pkg/e"
	"github.com/cyicz123/todolist/user/pkg/logger"
	"github.com/cyicz123/todolist/user/pkg/repo"
)

// UserInterface is an interface for user management operations.
type UserInterface interface {
	// Register creates a new user.
	Register(info *repo.User) e.ErrInterface

	// Login returns the user after querying user info and checking the authentication.
	Login(info *repo.User) e.ErrInterface

	// UserDelete deletes the user with the specified name.
	UserDelete(name string) e.ErrInterface
}

// UserInfo provides methods for user management operations.
type UserInfo struct {
	r repo.UserRepository
	l logger.Interface
}

func NewUserInfo(r repo.UserRepository, l logger.Interface) *UserInfo {
	return &UserInfo{
		r: r,
		l: l,
	}
}

// Register creates a new user. Returns an error if the provided user name or password is invalid or if the user already exists.
func (u *UserInfo) Register(info *repo.User) e.ErrInterface {
	if !u.checkUserName(info.UserName) {
		u.l.Error(e.UserNameInvalid.Error())
		return e.UserNameInvalid
	}
	if u.checkUserExist(info.UserName) {
		u.l.Error(e.UserExist.Error())
		return e.UserExist
	}
	if !u.checkPassword(info.PasswordDigest) {
		u.l.Error(e.PasswordInvalid.Error())
		return e.PasswordInvalid
	}
	err := u.r.Create(info)
	if err != nil {
		u.l.Error(e.UserInternalErr.Error())
		return e.UserInternalErr
	}
	return nil
}

// Login returns the user with the specified name. Returns an error if the provided user name is invalid or if the user does not exist.
func (u *UserInfo) Login(info *repo.User) e.ErrInterface {
	if !u.checkUserName(info.UserName) {
		u.l.Error(e.UserNameInvalid.Error())
		return e.UserNameInvalid
	}
	user, err := u.r.GetByName(info.UserName)
	if err != nil {
		u.l.Error(err)
		return e.UserInternalErr
	}
	if user.PasswordDigest != info.PasswordDigest {
		u.l.Warn(e.UserPasswdNotMatch.Error())
		return e.UserPasswdNotMatch
	}
	*info = *user
	return nil
}

// UserDelete deletes the user with the specified name. Returns an error if the provided user name is invalid or if the user does not exist.
func (u *UserInfo) UserDelete(name string) e.ErrInterface {
	if !u.checkUserName(name) {
		u.l.Error(e.UserNameInvalid.Error())
		return e.UserNameInvalid
	}
	err := u.r.Delete(name)
	if err != nil {
		u.l.Error(e.UserInternalErr.Error())
		return e.UserInternalErr
	}
	return nil
}

// checkUserName checks whether the specified user name is valid.
func (u *UserInfo) checkUserName(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]{4,32}$`)
    return re.MatchString(name)
}

// checkUserExist checks whether a user with the specified name already exists.
func (u *UserInfo) checkUserExist(name string) bool {
	if _, err := u.r.GetByName(name); err != nil {
		return false
	}
	return true
}

// checkPassword checks whether the specified password is valid.
func (u *UserInfo) checkPassword(pw string) bool {
	return pw != ""
}