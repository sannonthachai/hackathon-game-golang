package repository

import (
	"fmt"

	model "gitlab.com/sannonthachai/find-the-hidden-backend/model/user"
)

func (r *userRepository) CreateUser(payload model.User) error {
	if err := r.userDB.Table("users").Create(&payload).Error; err != nil {
		fmt.Println("Error Repo CreateUser: ", err)
		return err
	}

	return nil
}

func (r *userRepository) FindUser(payload model.Login) (model.User, error) {
	user := model.User{}
	if err := r.userDB.Table("users").Where("username = ?", payload.Username).Find(&user).Error; err != nil {
		fmt.Println("Error Repo FindUser: ", err)
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindUserByUserId(id int) (model.User, error) {
	user := model.User{}

	if err := r.userDB.Table("users").Where("id = ?", id).Find(&user).Error; err != nil {
		fmt.Println("Error Repo FindUserByUserId: ", err)
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUserPoint(userID int, point int) error {
	if err := r.userDB.Table("users").Where("id = ?", userID).Update("point", point).Error; err != nil {
		fmt.Println("Error Repo SaveUserPoint: ", err)
		return err
	}

	return nil
}

func (r *userRepository) SaveUserPointByChapter(payload model.Chapter) error {
	if r.userDB.Table("chapters").Where("user_id = ? AND chapter = ?", payload.UserId, payload.Chapter).Update("point", payload.Point).RowsAffected == 0 {
		if err := r.userDB.Table("chapters").Create(&payload).Error; err != nil {
			fmt.Println("Error Repo SaveUserPointByChapter: ", err)
			return err
		}
	}

	return nil
}

func (r *userRepository) GetUserPointByChapter(userID, chapter int) (model.Chapter, error) {
	userPoint := model.Chapter{}

	if err := r.userDB.Table("chapters").Where("user_id = ? AND chapter = ?", userID, chapter).Find(&userPoint).Error; err != nil {
		fmt.Println("Error Repo GetUserPointByChapter: ", err)
		return userPoint, err
	}

	return userPoint, nil
}

func (r *userRepository) GetUserPoint(userID int) ([]model.Chapter, error) {
	userPoint := []model.Chapter{}

	if err := r.userDB.Table("chapters").Where("user_id = ?", userID).Find(&userPoint).Error; err != nil {
		fmt.Println("Error Repo GetUserPoint: ", err)
		return userPoint, err
	}

	return userPoint, nil
}

func (r *userRepository) GetTopUserPointByChapter(chapter int) ([]model.UserPoint, error) {
	topUserPoint := []model.UserPoint{}

	if err := r.userDB.Table("chapters").
		Select("chapters.point, users.username").
		Joins("left join users on users.id = chapters.user_id").
		Where("chapter = ?", chapter).
		Order("point desc").
		Limit(10).
		Scan(&topUserPoint).
		Error; err != nil {
		fmt.Println("Error Repo GetTopUserPointByChapter: ", err)
		return topUserPoint, err
	}

	return topUserPoint, nil
}

func (r *userRepository) GetTopUserPoint() ([]model.UserPoint, error) {
	topUserPoint := []model.UserPoint{}

	if err := r.userDB.Table("users").Order("point desc").Limit(10).Find(&topUserPoint).Error; err != nil {
		fmt.Println("Error Repo GetTopUserPointByChapter: ", err)
		return topUserPoint, err
	}

	return topUserPoint, nil
}
