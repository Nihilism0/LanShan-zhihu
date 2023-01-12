package model

type Favorites struct {
	ID            uint   `gorm:"primarykey" db:"id"`
	UserName      string `db:"username" form:"username" json:"username"`
	FavoritesName string `db:"favoritesname" form:"favoritesname" json:"favoritesname" binding:"required"`
	Describe      string `db:"describe" form:"describe" json:"describe"`
	Private       uint   `db:"private" form:"private" json:"private" `
}
type FavoritesDelete struct {
	ID uint `form:"id" json:"id" binding:"required"`
}
type Collection struct {
	ID           uint `gorm:"primarykey" db:"id"`
	Favorites_id uint `db:"favorites_id" form:"favorites_id" json:"favorites_id" binding:"required"`
	Answer_id    uint `db:"answer_id" form:"answer_id" json:"answer_id" binding:"required"`
}
type RealCollection struct {
	ID           uint `gorm:"primarykey" db:"id"`
	Favorites_id uint `db:"favorites_id" form:"favorites_id" json:"favorites_id" binding:"required"`
	Answer_id    uint `db:"answer_id" form:"answer_id" json:"answer_id" binding:"required"`
	Article_id   uint `db:"article_id" form:"article_id" json:"article_id" binding:"required"`
}
type ArticleCollection struct {
	ID           uint `gorm:"primarykey" db:"id"`
	Favorites_id uint `db:"favorites_id" form:"favorites_id" json:"favorites_id" binding:"required"`
	Article_id   uint `db:"article_id" form:"article_id" json:"article_id" binding:"required"`
}
type FavoritesID struct {
	ID uint `form:"id" json:"id" binding:"required"`
}
type FavoritesModify struct {
	ID            uint   `gorm:"primarykey"form:"id" json:"id" db:"id" binding:"required"`
	FavoritesName string `db:"favoritesname" form:"favoritesname" json:"favoritesname" `
	Private       uint   `db:"private" form:"private" json:"private" `
	Describe      string `db:"describe" form:"describe" json:"describe"`
}
