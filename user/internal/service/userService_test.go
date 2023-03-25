package service

import (
	"context"
	"testing"

	"github.com/cyicz123/todolist/user/mocks"
	"github.com/cyicz123/todolist/user/pkg/e"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testCases struct {
	name     string
	input    *UserRequest
	expected interface{}
	mockExpected interface{}
}

func TestReq2UserModel(t *testing.T) {
	u := &UserService{
		user: nil,
	}
	req := &UserRequest{
		UserName: "username1",
		NickName: "nickname1",
		Password: "abc12345",
	}
	user, err := u.req2userModel(req)
	assert.NoError(t, err)
	assert.Equal(t, req.UserName, user.UserName)
	assert.Equal(t, req.NickName, user.NickName)
	assert.Equal(t, req.Password, user.PasswordDigest)
}

var loginTc = []*testCases{
    {
            name: "Login_Success",
            input:  &UserRequest{},
            expected: &UserDetailResponse{
                Code: uint32(e.SUCCESS.Code()),
            },
            mockExpected: nil,
        },
        {
            name:          "Login_Failed",
            input:         &UserRequest{},
            expected:      &UserDetailResponse{Code: e.ServiceInternalErr.Code()},
            mockExpected:  e.ServiceInternalErr,
    },
}

func TestUserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := mocks.NewMockUserInterface(ctrl)

    u := NewUserService(user)
    
    for _, tc := range loginTc {
		t.Run(tc.name, func(t *testing.T) {
			user.EXPECT().Login(gomock.Any()).Return(tc.mockExpected).Times(1)
			result, err := u.UserLogin(context.Background(),tc.input)
			assert.Equal(t, tc.expected, result)
            assert.Equal(t, tc.mockExpected, err)
        })
	}
}

var registerTc = []*testCases{
    {
            name: "Register_Success",
            input:  &UserRequest{},
            expected: &UserDetailResponse{
                Code: uint32(e.SUCCESS.Code()),
            },
            mockExpected: nil,
        },
        {
            name:          "Register_Failed",
            input:         &UserRequest{},
            expected:      &UserDetailResponse{Code: e.ServiceInternalErr.Code()},
            mockExpected:  e.ServiceInternalErr,
    },
}

func TestUserRegister(t *testing.T) {
    ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := mocks.NewMockUserInterface(ctrl)

    u := NewUserService(user)
    
    for _, tc := range registerTc {
		t.Run(tc.name, func(t *testing.T) {
			user.EXPECT().Register(gomock.Any()).Return(tc.mockExpected).Times(1)
			result, err := u.UserRegister(context.Background(),tc.input)
			assert.Equal(t, tc.expected, result)
            assert.Equal(t, tc.mockExpected, err)
        })
	} 
}

func TestUserLogout(t *testing.T) {
    mockReq := &UserRequest{}
    service := &UserService{}
    resp, err := service.UserLogout(context.Background(), mockReq)
    assert.NotNil(t, resp)
    assert.Nil(t, err)
}
