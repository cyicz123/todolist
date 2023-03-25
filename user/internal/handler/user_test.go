package handler

import (
	"testing"

	"github.com/cyicz123/todolist/user/mocks"
	"github.com/cyicz123/todolist/user/pkg/e"
	"github.com/cyicz123/todolist/user/pkg/repo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testCases struct {
	name     string
	input    string
	expected interface{}
	mockExpected []interface{}
}

// var user repo.User

var userNameTc = []testCases{
	{
		name:     "Valid username with lowercase letters",
		input:    "john_doe",
		expected: true,
	},
	{
		name:     "Valid username with uppercase letters",
		input:    "John_Doe",
		expected: true,
	},
	{
		name:     "Valid username with digits",
		input:    "johndoe123",
		expected: true,
	},
	{
		name:     "Invalid username with special character",
		input:    "john@doe",
		expected: false,
	},
	{
		name:     "Invalid username that is too short",
		input:    "j",
		expected: false,
	},
	{
		name:     "Invalid username that is too long",
		input:    "this_username_is_way_too_long_to_be_valid",
		expected: false,
	},
}

func TestCheckUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mocks.NewMockInterface(ctrl)
	r := mocks.NewMockUserRepository(ctrl)

	u := &UserInfo{
		l: log,
		r: r,
	}

	for _, tc := range userNameTc {
		t.Run(tc.name, func(t *testing.T) {
			result := u.checkUserName(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

var passwordTc = []testCases{
	{
		name:     "Invalid password that is null",
		input:    "",
		expected: false,
	},
	{
		name:     "valid password",
		input:    "Abc12345",
		expected: true,
	},
}

func TestCheckPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mocks.NewMockInterface(ctrl)
	r := mocks.NewMockUserRepository(ctrl)

	u := &UserInfo{
		l: log,
		r: r,
	}

	for _, tc := range passwordTc {
		t.Run(tc.name, func(t *testing.T) {
			result := u.checkPassword(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

var userExistTc = []testCases{
	{
		name:     "Exist",
		input:    "username1",
		expected: true,
		mockExpected: []interface{}{nil, nil},
	},
	{
		name:     "Not exist",
		input:    "username2",
		expected: false,
		mockExpected: []interface{}{nil, e.UserNotExist},
	},
}

func TestCheckUserExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mocks.NewMockInterface(ctrl)
	r := mocks.NewMockUserRepository(ctrl)

	u := &UserInfo{
		l: log,
		r: r,
	}

	for _, tc := range userExistTc {
		t.Run(tc.name, func(t *testing.T) {
			r.EXPECT().GetByName(tc.input).Return(tc.mockExpected...).Times(1)
			result := u.checkUserExist(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

var registerTc = []struct{
	name		string
	input   	*repo.User
	expected	error
	mockExpected	[][]interface{}
}{
	{
		name: "Valid user",
		input: &repo.User{
			UserName:        "valid_user",
			PasswordDigest:  "Abc12345",
		},
		expected: nil,
		mockExpected: [][]interface{}{
			{
				nil,
				e.UserNotExist,
			},
			{nil},
		},
	},
	{
		name: "Invalid username",
		input: &repo.User{
			UserName:        "john@doe",
			PasswordDigest:  "Abc12345",
		},
		expected: e.UserNameInvalid,
		mockExpected: [][]interface{}{
			{
				nil, nil,
			},
			{nil},
		},
	},
	{
		name: "Existing user",
		input: &repo.User{
			UserName:        "existing_user",
			PasswordDigest:  "Abc12345",
		},
		expected: e.UserExist,
		mockExpected: [][]interface{}{
			{
				nil, nil,
			},
			{nil},
		},
	},
	{
		name: "Invalid password",
		input: &repo.User{
			UserName:        "valid_user",
			PasswordDigest:  "",
		},
		expected: e.PasswordInvalid,
		mockExpected: [][]interface{}{
			{
				nil, 
				e.UserNotExist,
			},
			{nil},
		},
	},
	{
		name: "Create user failed",
		input: &repo.User{
			UserName:        "valid_user",
			PasswordDigest:  "abc12345",
		},
		expected: e.UserInternalErr,
		mockExpected: [][]interface{}{
			{
				nil,
				e.UserNotExist,
			},
			{e.UserInternalErr},
		},
	},
}

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mocks.NewMockInterface(ctrl)
	r := mocks.NewMockUserRepository(ctrl)

	u := &UserInfo{
		l: log,
		r: r,
	}
	log.EXPECT().Error(gomock.Any()).AnyTimes()
	for _, tc := range registerTc {
		t.Run(tc.name, func(t *testing.T) {
			r.EXPECT().GetByName(tc.input.UserName).Return(tc.mockExpected[0]...).MaxTimes(1)
			r.EXPECT().Create(tc.input).Return(tc.mockExpected[1]...).MaxTimes(1)

			result := u.Register(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

var loginTc = []struct{
	name		string
	input		*repo.User
	expected	error
	mockExpected	[]interface{}
}{
	{
		name: "Valid user",
		input: &repo.User{
			UserName:        "valid_user",
			PasswordDigest:  "Abc12345",
		},
		expected: nil,
		mockExpected: []interface{}{
			&repo.User{
				UserID: 0,
				UserName: "valid_user",
				PasswordDigest: "Abc12345",
				NickName: "valid_user",
			},
			nil,
		},
	},
	{
		name: "Invalid username",
		input: &repo.User{
			UserName:        "john@doe",
			PasswordDigest:  "Abc12345",
		},
		expected: e.UserNameInvalid,
		mockExpected: []interface{}{
			nil, nil,
		},
	},
	{
		name: "Get user info failed",
		input: &repo.User{
			UserName:        "existing_user",
			PasswordDigest:  "Abc12345",
		},
		expected: e.UserInternalErr,
		mockExpected: []interface{}{
			nil, e.UserInternalErr,
		},
	},
	{
		name: "User and password not match",
		input: &repo.User{
			UserName:        "valid_user",
			PasswordDigest:  "abc12345",
		},
		expected: e.UserPasswdNotMatch,
		mockExpected: []interface{}{
			&repo.User{
				UserName:	"valid_user",
				PasswordDigest: "Abc12345",
			},
			nil,
		},
	},
}

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mocks.NewMockInterface(ctrl)
	r := mocks.NewMockUserRepository(ctrl)

	u := &UserInfo{
		l: log,
		r: r,
	}
	log.EXPECT().Error(gomock.Any()).AnyTimes()
	log.EXPECT().Warn(gomock.Any()).AnyTimes()
	for _, tc := range loginTc {
		t.Run(tc.name, func(t *testing.T) {
			r.EXPECT().GetByName(tc.input.UserName).Return(tc.mockExpected...).MaxTimes(1)

			result := u.Login(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

var userDeleteTc = []testCases{
	{
		name:     "success",
		input:    "username",
		expected: nil,
		mockExpected: []interface{}{nil},
	},
	{
		name:     "invalid username",
		input:    "j",
		expected: e.UserNameInvalid,
		mockExpected: []interface{}{nil},
	},
	{
		name:     "delete error",
		input:    "username",
		expected: e.UserInternalErr,
		mockExpected: []interface{}{e.UserInternalErr},
	},
}

func TestUserDelet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mocks.NewMockInterface(ctrl)
	r := mocks.NewMockUserRepository(ctrl)

	u := &UserInfo{
		l: log,
		r: r,
	}
	log.EXPECT().Error(gomock.Any()).AnyTimes()
	for _, tc := range userDeleteTc {
		t.Run(tc.name, func(t *testing.T) {
			r.EXPECT().Delete(tc.input).Return(tc.mockExpected...).MaxTimes(1)

			result := u.UserDelete(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}