package user

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/config"
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
	tx := u.db.Preload("Books").Where("users.uid = ?", uid).First(&user)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	if err != nil {
		fmt.Println(err)
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
