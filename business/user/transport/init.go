package transport

import "gitlab.com/sannonthachai/find-the-hidden-backend/business/user"

type Handler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) Handler {
	return Handler{
		userService: userService,
	}
}
