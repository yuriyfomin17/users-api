package repository

import (
	"github.com/gin-gonic/gin"
	"usersRestAPI/internal/app/repository/model"
)

type UserRepository interface {
	FindUserById(ctx *gin.Context, userId string) (model.User, error)
	FindAllUsers(ctx *gin.Context) ([]model.User, error)
	CreateUser(ctx *gin.Context, user model.User) (model.User, error)
	UpdateUser(ctx *gin.Context, user model.User) (model.User, error)
	DeleteUserById(ctx *gin.Context, userId string) error
}
