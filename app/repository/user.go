package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
)

type (
	UserRepository interface {
		CreateUser() (user entity.User, err error)
		FindUserByUID(uid string) (entity.User, error)
	}
	userRepository struct {
		db config.Database
	}
)

func NewUserRepository(db config.Database) UserRepository {
	return userRepository{db}
}

func (u userRepository) FindUserByUID(uid string) (user entity.User, err error) {
	user = entity.User{Uid: uid}
	res := u.db.Where(&user).First(&user)
	if res.Error != nil {
		err = res.Error
		return
	}
	if res.RowsAffected == 0 {
		err = errors.New("User not found")
		return
	}
	return
}

func (u userRepository) CreateUser() (user entity.User, err error) {
	s := uuid.New()
	user = entity.User{Uid: s.String()}

	res := u.db.Create(&user)

	if res.Error != nil {
		err = res.Error
		return
	}

	if res.RowsAffected == 0 {
		err = errors.New("User creation Failed")
		return
	}
	return
}
