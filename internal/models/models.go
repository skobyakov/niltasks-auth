package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	NickName string `gorm:"unique" json:"nickName"`
	Password string `json:"password"`
}

type AuthRequest struct {
	NickName string `json:"nickName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Session string `json:"session"`
}
