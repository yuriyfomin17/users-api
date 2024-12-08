package service

import (
	"github.com/gin-gonic/gin"
	"usersRestAPI/internal/app/converter"
	"usersRestAPI/internal/app/dto"
	"usersRestAPI/internal/app/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func (s UserServiceImpl) FindUserById(ctx *gin.Context, userId string) (dto.User, error) {
	modelUser, err := s.UserRepository.FindUserById(ctx, userId)
	return converter.UserToDto(modelUser), err
}

func (s UserServiceImpl) FindAllUsers(ctx *gin.Context) ([]dto.User, error) {
	var dtoUsers = make([]dto.User, 0)
	users, err := s.UserRepository.FindAllUsers(ctx)
	if err != nil {
		return dtoUsers, err
	}
	for _, user := range users {
		dtoUsers = append(dtoUsers, converter.UserToDto(user))
	}
	return dtoUsers, err
}

func (s UserServiceImpl) CreateUser(ctx *gin.Context, user dto.User) (dto.User, error) {
	invalidUserError := user.Validate()
	if invalidUserError != nil {
		return dto.User{}, invalidUserError
	}
	createUser, invalidUserError := s.UserRepository.CreateUser(ctx, converter.DtoToUser(user))
	return converter.UserToDto(createUser), invalidUserError
}

func (s UserServiceImpl) UpdateUser(ctx *gin.Context, user dto.User) (dto.User, error) {
	invalidUserError := user.Validate()
	if invalidUserError != nil {
		return dto.User{}, invalidUserError
	}
	invalidUuid := user.ValidateID()
	if invalidUuid != nil {
		return dto.User{}, invalidUuid
	}

	updatedUser, invalidUpdateError := s.UserRepository.UpdateUser(ctx, converter.DtoToUser(user))
	return converter.UserToDto(updatedUser), invalidUpdateError
}

func (s UserServiceImpl) DeleteUserById(ctx *gin.Context, userId string) error {
	return s.UserRepository.DeleteUserById(ctx, userId)
}
