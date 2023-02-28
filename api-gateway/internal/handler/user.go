package handler

import (
	"github.com/gin-gonic/gin"
)

type User interface {
	Register(ctx *gin.Context) error
	Login(ctx *gin.Context) error
}

type HandleModel struct {
	user User
}

func (h *HandleModel)Register(ctx *gin.Context) {
	
}