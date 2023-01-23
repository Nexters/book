package repository

import (
	"errors"

	"github.com/chaewonkong/go-template/app/config"
	"github.com/chaewonkong/go-template/app/entity"
	"github.com/google/uuid"
)

type (
	UserRepository interface {
		CreateUser() (interface{}, error)
	}
	userRepository struct {
		db config.Database
	}
)

func NewUserRepository(db config.Database) UserRepository {
	return userRepository{db}
}

func (u userRepository) CreateUser() (interface{}, error) {
	s := uuid.New()
	user := entity.User{Uid: s.String()}

	res := u.db.Create(&user)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("User creation Failed")
	}

	return user, nil
}
