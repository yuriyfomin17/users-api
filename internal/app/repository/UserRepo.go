package repository

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
	"time"
	"usersRestAPI/internal/app/repository/model"
	"usersRestAPI/internal/pkg/pg"
)

type UserRepositoryImpl struct {
	Db *pg.DB
}

func (u UserRepositoryImpl) FindUserById(ctx *gin.Context, userId string) (model.User, error) {
	var user model.User
	err := u.Db.NewSelect().Model(&user).Where("id = ?", userId).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, ErrNotFound
		}
		return model.User{}, ErrNotFound
	}
	return user, nil
}

func (u UserRepositoryImpl) CreateUser(ctx *gin.Context, user model.User) (model.User, error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	_, err := u.Db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		if isDuplicateEmail(err) {
			return model.User{}, ErrDuplicatedUser
		}
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepositoryImpl) UpdateUser(ctx *gin.Context, user model.User) (model.User, error) {
	_, err := u.Db.NewUpdate().Model(&user).Where("id = ?", user.ID).Exec(ctx)
	if err != nil {
		if isDuplicateEmail(err) {
			return model.User{}, ErrDuplicatedUser
		}
		return model.User{}, err
	}
	return user, nil
}
func (u UserRepositoryImpl) FindAllUsers(ctx *gin.Context) ([]model.User, error) {
	var users []model.User
	err := u.Db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserRepositoryImpl) DeleteUserById(ctx *gin.Context, userId string) error {
	_, err := u.Db.NewDelete().Model(&model.User{}).Where("id = ?", userId).Exec(ctx)
	if err != nil {
		return ErrOnDelete
	}
	return nil
}

func isDuplicateEmail(err error) bool {
	return err != nil && strings.Contains(err.Error(), "23505")
}
