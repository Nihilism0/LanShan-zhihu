package api

import (
	"CSAwork/dao"
	"CSAwork/global"
	"CSAwork/model"
	"CSAwork/utils"
	"CSAwork/utils/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateArticle(c *gin.Context) {
	global.Bucket.Take(1)
	article := model.Article{}
	if err := c.ShouldBind(&article); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	title := article.Title
	message := article.Message
	TimeNow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "insert into articles(created_at, updated_at, articler, title, message) values (?,?,?,?,?)"
	ret, _ := global.GlobalDb1.Exec(sqlStr, TimeNow, TimeNow, username.(string), title, message)
	TheId, _ := ret.LastInsertId()
	global.RedisDb.Do("select", 3)
	global.RedisDb.SAdd("articleids", TheId)
	global.RedisDb.Do("select", 0)
	dao.AddGod(username.(string), 5)
	utils.RespSuccess(c, "亲爱的"+username.(string)+",您成功发布文章,GO币+5==>"+title)
	ID := dao.GetIdFromUsername(username.(string))
	middleware.Producer("Subscribe", strconv.Itoa(int(ID))+" 发布了个文章!:"+article.Title)
}

func CreateArticleComment(c *gin.Context) {
	comment := model.ArticleComment{}
	if err := c.ShouldBind(&comment); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	message := comment.Message
	articleid := comment.Article_id
	TimeNow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "insert into article_comments(created_at, updated_at, commenter, message, article_id) values (?,?,?,?,?)"
	ret, err := global.GlobalDb1.Exec(sqlStr, TimeNow, TimeNow, username.(string), message, articleid)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "评论article出大问题")
		return
	}
	TheId, _ := ret.LastInsertId()
	global.RedisDb.Do("select", 4)
	global.RedisDb.SAdd("articlecommentids", TheId)
	global.RedisDb.Do("select", 0)
	dao.AddGod(username.(string), 1)
	utils.RespSuccess(c, "发送成功,GO币+1")
}
func ArticleSubmited(c *gin.Context) {
	username, _ := c.Get("username")
	articles := dao.FindArticleSubmited(username.(string))
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": articles,
	})
}
func SeeArticleInformation(c *gin.Context) {
	seeinformation := model.SeeInformation{}
	if err := c.ShouldBind(&seeinformation); err != nil {
		fmt.Println(err)
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	if seeinformation.Secret != "123456" {
		utils.RespFail(c, "Secret wrong!Get out of my website!")
		return
	}
	id := seeinformation.ID
	article := dao.FindArticleInformation(id)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"question": article,
	})
}
func RandArticles(c *gin.Context) {
	articles := dao.RandArticles()
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"random": articles,
	})
}
func ArticleDelete(c *gin.Context) {
	deleteform := model.Delete{}
	if err := c.ShouldBind(&deleteform); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	articleId := deleteform.ID
	OK := dao.JudgeArticle(username.(string), articleId)
	if !OK {
		utils.RespFail(c, "这个文章不属于你")
		return
	}
	dao.ArticleDelete(articleId)
	dao.CancelArticleInFavorites(articleId)
	utils.RespSuccess(c, "删除文章成功!!!")
}
func ArticleCommentDelete(c *gin.Context) {
	deleteform := model.Delete{}
	if err := c.ShouldBind(&deleteform); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	commentId := deleteform.ID
	OK := dao.JudgeArticleComment(username.(string), commentId)
	if !OK {
		utils.RespFail(c, "这个文章不属于你")
		return
	}
	dao.ArticleCommentDelete(commentId)
	utils.RespSuccess(c, "删除文章成功!!!")
}
