package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
)

type (
	// UserRepository User Repository Interface
	UserRepository interface {
		CreateUser() (user entity.User, err error)
		FindUserByUID(uid string) (entity.User, error)
	}

	// userRepository userRepository Struct
	userRepository struct {
		db config.Database
	}
)

// NewUserRepository 생성자
func NewUserRepository(db config.Database) UserRepository {
	return userRepository{db}
}

// FindUserByUID UID로 사용자 조회
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

// CreateUser 사용자 생성
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
