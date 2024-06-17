package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/sannonthachai/find-the-hidden-backend/business/user"
)

type userRepository struct {
	userDB *gorm.DB
}

func NewUserRepository(userDB *gorm.DB) user.Repository {
	return &userRepository{
		userDB: userDB,
	}
}
