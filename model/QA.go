package model

import (
	"time"
)

type Answer struct {
	ID          uint      `gorm:"primarykey" db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	DeletedAt   time.Time `gorm:"index" db:"deleted_at"`
	Answerer    string    `db:"answerer"`
	Message     string    `db:"message" form:"message" json:"message" binding:"required"`
	Question_id uint      `db:"question_id" form:"question_id" json:"question_id" binding:"required"`
	Comments    []Comment `db:"comments"`
}

type AnswerInfo struct {
	ID          uint      `gorm:"primarykey" db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	Answerer    string    `db:"answerer"`
	Message     string    `db:"message" form:"message" json:"message" binding:"required"`
	Question_id uint      `db:"question_id" form:"question_id" json:"question_id" binding:"required"`
	Hots        int64
}

type Question struct {
	ID         uint      `gorm:"primarykey" db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	DeletedAt  time.Time `gorm:"index" db:"deleted_at"`
	Questioner string    `db:"questioner"`
	Title      string    `db:"title" form:"title" json:"title" binding:"required"`
	Message    string    `db:"message" form:"message" json:"message" binding:"required"`
	Answers    []Answer  `db:"answers"`
}

type Comment struct {
	ID        uint      `gorm:"primarykey" db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `gorm:"index" db:"deleted_at"`
	Commenter string    `db:"commenter"`
	Message   string    `db:"message" form:"message" json:"message" binding:"required"`
	Answer_id uint      `form:"id" json:"id" binding:"required"`
}

type QuestionInfo struct {
	ID         uint   `gorm:"primarykey" db:"id"`
	Title      string `db:"title" form:"title" json:"title" binding:"required"`
	Questioner string `db:"questioner"`
	Message    string `db:"message"`
	Hots       int64
}

type QuestionModify struct {
	ID      uint   `db:"id" form:"id" json:"id" binding:"required"`
	Message string `db:"message" form:"message" json:"message" binding:"required"`
}

type AnswerModify struct {
	ID      uint   `db:"id" form:"id" json:"id" binding:"required"`
	Message string `db:"message" form:"message" json:"message" binding:"required"`
}

type Delete struct {
	ID uint `db:"id" form:"id" json:"id" binding:"required"`
}
type Praise struct {
	ID uint `db:"id" form:"id" json:"id" binding:"required"`
}
type Search struct {
	Message string `db:"message" form:"message" json:"message" binding:"required"`
}
type SeeInformation struct {
	ID     uint   `db:"id" form:"id" json:"id" binding:"required"`
	Secret string `db:"secret" form:"secret" json:"secret" binding:"required"`
}
