package model

import "time"

type Record struct {
	UserId     uint      `db:"user_id" form:"user_id" json:"user_id" binding:"required"`
	Time       time.Time `db:"time" form:"time" json:"time" binding:"rtime"`
	QuestionId uint      `db:"question_id" form:"question_id" json:"question_id" binding:"required"`
	ArticleId  uint      `db:"article_id" form:"article_id" json:"article_id" binding:"required"`
}
