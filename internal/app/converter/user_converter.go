package converter

import (
	"usersRestAPI/internal/app/dto"
	"usersRestAPI/internal/app/repository/model"
)

func DtoToUser(userDto dto.User) model.User {
	return model.User{
		ID:        userDto.ID,
		FirstName: userDto.Firstname,
		LastName:  userDto.Lastname,
		Email:     userDto.Email,
		Age:       userDto.Age,
		CreatedAt: userDto.Created,
	}
}

func UserToDto(user model.User) dto.User {
	return dto.User{
		ID:        user.ID,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		Age:       user.Age,
		Created:   user.CreatedAt,
	}
}
