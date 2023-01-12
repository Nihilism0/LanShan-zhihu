package model

import "time"

type Article struct {
	ID        uint             `gorm:"primarykey" db:"id"`
	CreatedAt time.Time        `db:"created_at"`
	UpdatedAt time.Time        `db:"updated_at"`
	Articler  string           `db:"articler"`
	Title     string           `db:"title" form:"title" json:"title" binding:"required"`
	Message   string           `db:"message" form:"message" json:"message" binding:"required"`
	Comments  []ArticleComment `db:"answers"`
}
type ArticleComment struct {
	ID         uint      `gorm:"primarykey" db:"id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Commenter  string    `db:"commenter"`
	Message    string    `db:"message" form:"message" json:"message" binding:"required"`
	Article_id uint      `form:"id" json:"id" binding:"required"`
}
type ArticleInfo struct {
	ID       uint   `gorm:"primarykey" db:"id"`
	Title    string `db:"title" form:"title" json:"title" binding:"required"`
	Articler string `db:"articler"`
	Message  string `db:"message" form:"message" json:"message" binding:"required"`
	Hots     int64
}
