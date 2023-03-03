package repo

import (
	"testing"
	"user/pkg/logger"
	"user/config"
	
	"github.com/stretchr/testify/assert"
)

// TestUserRepository is a test function that tests the UserRepository interface and its implementations.
func TestUserRepository(t *testing.T) {
	// 实例化一个UserRepository
	dbFactoty := &MysqlFactory{}
	l := logger.New("repoTest")
	v := config.GetInstance()
	userRepo, err := dbFactoty.New(l, v)
	assert.NoError(t, err)

	// 创建一个用户
	user := &User{
		UserName:       "testuser",
		NickName:       "Test User",
		PasswordDigest: "testpassword",
	}
	err = userRepo.Create(user)
	assert.NoError(t, err)

	// 通过UserName查询用户
	foundUser, err := userRepo.GetByName("testuser")
	assert.NoError(t, err)
	assert.Equal(t, user.UserName, foundUser.UserName)
	assert.Equal(t, user.NickName, foundUser.NickName)
	assert.Equal(t, user.PasswordDigest, foundUser.PasswordDigest)

	// 通过UserID查询用户
	foundUser, err = userRepo.GetByID((user.UserID))
	assert.NoError(t, err)
	assert.Equal(t, user.UserName, foundUser.UserName)
	assert.Equal(t, user.NickName, foundUser.NickName)
	assert.Equal(t, user.PasswordDigest, foundUser.PasswordDigest)

	// 修改用户
	user.NickName = "Updated Test User"
	err = userRepo.Update(user)
	assert.NoError(t, err)

	// 再次查询，确认修改成功
	foundUser, err = userRepo.GetByName("testuser")
	assert.NoError(t, err)
	assert.Equal(t, user.UserName, foundUser.UserName)
	assert.Equal(t, user.NickName, foundUser.NickName)
	assert.Equal(t, user.PasswordDigest, foundUser.PasswordDigest)

	// 删除用户
	err = userRepo.Delete("testuser")
	assert.NoError(t, err)

	// 再次查询，确认删除成功
	_, err = userRepo.GetByName("testuser")
	assert.Error(t, err)
}
