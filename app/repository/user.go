package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
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
