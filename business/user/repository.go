package user

import model "gitlab.com/sannonthachai/find-the-hidden-backend/model/user"

type Repository interface {
	CreateUser(payload model.User) error
	FindUser(payload model.Login) (model.User, error)
	FindUserByUserId(id int) (model.User, error)
	UpdateUserPoint(userID int, point int) error
	SaveUserPointByChapter(payload model.Chapter) error
	GetUserPointByChapter(userID, chapter int) (model.Chapter, error)
	GetUserPoint(userID int) ([]model.Chapter, error)
	GetTopUserPointByChapter(chapter int) ([]model.UserPoint, error)
	GetTopUserPoint() ([]model.UserPoint, error)
}
