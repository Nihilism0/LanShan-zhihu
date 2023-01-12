package dao

import (
	"CSAwork/global"
	"CSAwork/model"
	"log"
)

func SelectArticle(id string) bool {
	flag, _ := global.RedisDb.SIsMember("articleids", id).Result()
	return flag
}
func SelectArticleComment(id string) bool {
	flag, _ := global.RedisDb.SIsMember("articlecommentids", id).Result()
	return flag
}
func FindArticleSubmited(username string) []model.Article {
	var articles []model.Article
	sqlStr := "select * from articles where articler = ?"
	_ = global.GlobalDb1.Select(&articles, sqlStr, username)
	var comments []model.ArticleComment
	for k, v := range articles {
		sqlStr = "select id, created_at, updated_at,commenter, message, article_id from article_comments where article_id = ?"
		err := global.GlobalDb1.Select(&comments, sqlStr, v.ID)
		v.Comments = comments
		articles[k] = v
		if err != nil {
			log.Println("find answer error!:", err)
		}
	}
	return articles
}
func FindArticleInformation(id uint) model.Article {
	var article model.Article
	sqlStr := "select * from articles where id = ?"
	err := global.GlobalDb1.Get(&article, sqlStr, id)
	if err != nil {
		log.Println(err)
	}
	var comments []model.ArticleComment
	sqlStr = "select * from article_comments where article_id = ?"
	err = global.GlobalDb1.Select(&comments, sqlStr, article.ID)
	article.Comments = comments
	return article
}
func RandArticles() []model.ArticleInfo {
	articles := []model.ArticleInfo{}
	sqlStr := "select id,title,articler,message from articles order by rand()"
	err := global.GlobalDb1.Select(&articles, sqlStr)
	if err != nil {
		log.Println(err)
	}
	return articles
}
func JudgeArticle(username string, articleId uint) bool {
	OK := false
	var article model.ArticleInfo
	sqlStr := "select articler from articles where id = ?"
	_ = global.GlobalDb1.Get(&article, sqlStr, articleId)
	//global.GlobalDb1.Model(&model.Question{}).Where("questioner = ? AND id = ?", username, id).First(&question)
	if article.Articler == username {
		OK = true
	}
	return OK
}
func JudgeArticleComment(username string, commentId uint) bool {
	OK := false
	var comment model.ArticleComment
	sqlStr := "select commenter from article_comments where id = ?"
	_ = global.GlobalDb1.Get(&comment, sqlStr, commentId)
	//global.GlobalDb1.Model(&model.Question{}).Where("questioner = ? AND id = ?", username, id).First(&question)
	if comment.Commenter == username {
		OK = true
	}
	return OK
}
func ArticleDelete(articleId uint) {
	sqlStr := "delete from articles where id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, articleId)
}
func ArticleCommentDelete(commentId uint) {
	sqlStr := "delete from article_comments where id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, commentId)
}
func GetUsernameFromArticleId(id uint) string {
	var username string
	sqlStr := "select articler from articles where id=?"
	global.GlobalDb1.Get(&username, sqlStr, id)
	return username
}
