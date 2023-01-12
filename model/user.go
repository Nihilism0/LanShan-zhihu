package model

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primarykey" db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `gorm:"index" db:"deleted_at"`
	Email       string    `db:"email" form:"email" json:"email" binding:"required"`
	HeadPhoto   string    `db:"headphoto" form:"headphoto" json:"headphoto"`
	PhoneNumber uint      `db:"phonenumber" form:"phonenumber" json:"phonenumber" binding:"required"`
	UserName    string    `db:"username" form:"username" json:"username" binding:"required"`
	Password    string    `db:"password" form:"password" json:"password" binding:"required"`
	Gender      string    `db:"gender" form:"gender" json:"gender"`
	Gold        uint      `db:"gold" form:"gold" json:"gold"`
	Sign        string    `db:"sign" form:"sign" json:"sign"`
}

type UserInfo struct {
	UserName  string `db:"username" form:"username" json:"username"`
	Gender    string `db:"gender" form:"gender" json:"gender"`
	Sign      string `db:"sign" form:"sign" json:"sign"`
	HeadPhoto string `db:"headphoto" form:"headphoto" json:"headphoto"`
}

type UserRegister struct {
	Code        string `db:"code" form:"code" json:"code" binding:"required"`
	Email       string `db:"email" form:"email" json:"email" binding:"required"`
	PhoneNumber uint   `db:"phonenumber" form:"phonenumber" json:"phonenumber" binding:"required"`
	UserName    string `db:"username" form:"username" json:"username" binding:"required"`
	Password    string `db:"password" form:"password" json:"password" binding:"required"`
}
type AccountLogin struct {
	UserName string `db:"username" form:"username" json:"username" binding:"required"`
	Password string `db:"password" form:"password" json:"password" binding:"required"`
}
type PhoneLogin struct {
	PhoneNumber uint   `db:"phonenumber" form:"phonenumber" json:"phonenumber" binding:"required"`
	Code        string `db:"code" form:"code" json:"code" binding:"required"`
}
type PhonePasswordLogin struct {
	PhoneNumber uint   `db:"phonenumber" form:"phonenumber" json:"phonenumber" binding:"required"`
	Password    string `db:"password" form:"password" json:"password" binding:"required"`
}
type EmailLogin struct {
	Email    string `db:"email" form:"email" json:"email" binding:"required"`
	Password string `db:"password" form:"password" json:"password" binding:"required"`
}
type CodeNumber struct {
	PhoneNumber uint `db:"phonenumber" form:"phonenumber" json:"phonenumber" binding:"required"`
}
type GetPassword struct {
	Password string `db:"password" form:"password" json:"password"`
	Salt     string `db:"salt" form:"salt" json:"salt"`
}
type ChangePassword struct {
	NewPassword string `db:"newpassword" form:"newpassword" json:"newpassword" binding:"required"`
	Confirm     string `db:"confirm" form:"confirm" json:"confirm" binding:"required"`
	Code        string `db:"code" form:"code" json:"code" binding:"required"`
}
