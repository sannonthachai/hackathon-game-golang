package model

import "github.com/golang-jwt/jwt"

type User struct {
	ID       int
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Age      int    `json:"age" gorm:"column:age"`
	Point    int    `gorm:"column:point"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserPointByChapter struct {
	Username string `json:"username,omitempty" gorm:"column:username"`
	Chapter  int    `json:"chapter" gorm:"column:chapter"`
	Point    int    `json:"point,omitempty" gorm:"column:point"`
}

type UserPoint struct {
	Username string `json:"username" gorm:"column:username"`
	Point    int    `json:"point" gorm:"column:point"`
}

type Chapter struct {
	UserId  int `gorm:"column:user_id"`
	Chapter int `gorm:"column:chapter"`
	Point   int `gorm:"column:point"`
}

type Token struct {
	Token string `json:"token"`
}

type JwtCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}
