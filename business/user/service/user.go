package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	model "gitlab.com/sannonthachai/find-the-hidden-backend/model/user"
	"gitlab.com/sannonthachai/find-the-hidden-backend/util"
)

func (s *userService) Register(payload model.User) error {
	user := model.User{}
	user.Username = payload.Username
	user.Age = payload.Age

	password, err := util.HashPassword(payload.Password)
	if err != nil {
		fmt.Println("Error hash password", err)
		return err
	}

	user.Password = password

	if err := s.userRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (s *userService) Login(payload model.Login) (model.Token, error) {
	token := model.Token{}
	user, err := s.userRepo.FindUser(payload)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return token, errors.New("401")
		}
		return token, err
	}

	if !util.CheckPasswordHash(payload.Password, user.Password) {
		fmt.Println("Incorrect password")
		return token, errors.New("401")
	}

	claims := &model.JwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := jwtToken.SignedString([]byte(s.appConfig.Secret))
	if err != nil {
		fmt.Println("Error sign jwt token")
		return token, err
	}

	token.Token = t

	return token, nil
}

func (s *userService) UpdateUserPoint(userID int, payload model.UserPointByChapter) error {
	chapter := model.Chapter{}
	chapter.Chapter = payload.Chapter
	chapter.Point = payload.Point
	chapter.UserId = userID

	if err := s.userRepo.SaveUserPointByChapter(chapter); err != nil {
		return err
	}

	userPointAllChapter, err := s.userRepo.GetUserPoint(userID)
	if err != nil {
		return err
	}

	userPoint := 0

	for _, value := range userPointAllChapter {
		userPoint = userPoint + value.Point
	}

	if err := s.userRepo.UpdateUserPoint(userID, userPoint); err != nil {
		return err
	}

	return nil
}

func (s *userService) GetUserPointByChapter(userID, chapter int) (model.UserPointByChapter, error) {
	result := model.UserPointByChapter{}

	userPointByChapter, err := s.userRepo.GetUserPointByChapter(userID, chapter)
	if err != nil {
		return result, err
	}

	result.Chapter = userPointByChapter.Chapter
	result.Point = userPointByChapter.Point

	user, err := s.userRepo.FindUserByUserId(userID)
	if err != nil {
		return result, err
	}

	result.Username = user.Username

	return result, nil
}

func (s *userService) GetUserPoint(userID int) (model.UserPoint, error) {
	userPoint := model.UserPoint{}

	user, err := s.userRepo.FindUserByUserId(userID)
	if err != nil {
		return userPoint, err
	}

	userPoint.Username = user.Username

	userPointAllChapter, err := s.userRepo.GetUserPoint(userID)
	if err != nil {
		return userPoint, err
	}

	for _, value := range userPointAllChapter {
		userPoint.Point = userPoint.Point + value.Point
	}

	return userPoint, nil
}

func (s *userService) GetLeaderBoardByChapter(chapter int) ([]model.UserPoint, error) {
	userPoint, err := s.userRepo.GetTopUserPointByChapter(chapter)
	if err != nil {
		return userPoint, err
	}

	return userPoint, nil
}

func (s *userService) GetLeaderBoard() ([]model.UserPoint, error) {
	userPoint, err := s.userRepo.GetTopUserPoint()
	if err != nil {
		return userPoint, err
	}

	return userPoint, nil
}
