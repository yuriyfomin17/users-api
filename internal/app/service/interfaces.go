package service

import (
	"github.com/gin-gonic/gin"
	"usersRestAPI/internal/app/dto"
)

type UserService interface {
	FindUserById(ctx *gin.Context, userId string) (dto.User, error)
	FindAllUsers(ctx *gin.Context) ([]dto.User, error)
	CreateUser(ctx *gin.Context, user dto.User) (dto.User, error)
	UpdateUser(ctx *gin.Context, user dto.User) (dto.User, error)
	DeleteUserById(ctx *gin.Context, userId string) error
	DeleteUserB(ctx *gin.Context, userId string) error
}
