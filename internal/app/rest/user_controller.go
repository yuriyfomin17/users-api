package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"usersRestAPI/internal/app/dto"
	"usersRestAPI/internal/app/repository"
	"usersRestAPI/internal/app/service"
)

type UserController struct {
	UserService service.UserService
}

func (uc *UserController) GetUserById(c *gin.Context) {
	userId := c.Query("user_id")
	userDto, currError := uc.UserService.FindUserById(c, userId)
	if currError != nil {
		if errors.Is(currError, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": currError.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, userDto)
}
func (uc *UserController) GetAllUsers(c *gin.Context) {
	userDto, currError := uc.UserService.FindAllUsers(c)
	if currError != nil {
		if errors.Is(currError, repository.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": currError.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, userDto)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var newUser dto.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	newUser, userCreationErr := uc.UserService.CreateUser(c, newUser)
	if userCreationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": userCreationErr.Error()})
	}
	c.JSON(http.StatusCreated, newUser)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var newUser dto.User
	var userId = c.Param("id")
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
	}
	newUser.ID = parsedUUID
	updatedUser, updateUserCreationErr := uc.UserService.UpdateUser(c, newUser)
	if updateUserCreationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": updateUserCreationErr.Error()})
	}
	c.JSON(http.StatusOK, updatedUser)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	var userId = c.Param("id")
	errorOnDelete := uc.UserService.DeleteUserById(c, userId)
	if errorOnDelete != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errorOnDelete.Error()})
	}
	c.Status(http.StatusOK)
}
