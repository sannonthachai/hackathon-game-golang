package user

import model "gitlab.com/sannonthachai/find-the-hidden-backend/model/user"

type Service interface {
	Register(payload model.User) error
	Login(payload model.Login) (model.Token, error)
	UpdateUserPoint(userID int, payload model.UserPointByChapter) error
	GetUserPointByChapter(userID, chapter int) (model.UserPointByChapter, error)
	GetUserPoint(userID int) (model.UserPoint, error)
	GetLeaderBoardByChapter(chapter int) ([]model.UserPoint, error)
	GetLeaderBoard() ([]model.UserPoint, error)
}
