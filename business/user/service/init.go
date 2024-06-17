package service

import (
	"gitlab.com/sannonthachai/find-the-hidden-backend/business/user"
	"gitlab.com/sannonthachai/find-the-hidden-backend/model"
)

type userService struct {
	userRepo  user.Repository
	appConfig model.Config
}

func NewUserService(userRepo user.Repository, appConfig model.Config) user.Service {
	return &userService{
		userRepo:  userRepo,
		appConfig: appConfig,
	}
}
