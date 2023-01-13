package api

import (
	"CSAwork/dao"
	"CSAwork/global"
	"CSAwork/model"
	"CSAwork/utils"
	"CSAwork/utils/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PraiseQuestion(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	//flag1 := dao.SelectQuestion(id)
	//if !flag1 {
	//	utils.RespFail(c, "Question is not exist")
	//	return
	//}
	// 验证用户是否点赞
	flag2 := dao.SelectPraiseuser(id, userId)
	// 用户点过赞则退出
	if flag2 {
		utils.RespFail(c, "您已点赞,无需重复点赞!")
		return
	}
	dao.Praiseadd(id, userId)
	utils.RespSuccess(c, "点赞成功!")
	ID := dao.GetIdFromUsername(username.(string))
	middleware.Producer("Subscribe", strconv.Itoa(int(ID))+" 给:"+strconv.Itoa(int(praiseform.ID))+"问题点赞了哦")
}

func PraiseAnswer(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 1)
	flag1 := dao.SelectAnswer(id)
	global.RedisDb.Do("select", 0)
	if !flag1 {
		utils.RespFail(c, "Answer is not exist")
		return
	}
	// 验证用户是否点赞

	global.RedisDb.Do("select", 1)
	flag2 := dao.SelectPraiseuser(id, userId)
	global.RedisDb.Do("select", 0)
	// 用户点过赞则退出
	if flag2 {
		utils.RespFail(c, "您已点赞,无需重复点赞!")
		return
	}
	global.RedisDb.Do("select", 1)
	dao.Praiseadd(id, userId)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "点赞成功!")
	ID := dao.GetIdFromUsername(username.(string))
	middleware.Producer("Subscribe", strconv.Itoa(int(ID))+" 给:"+strconv.Itoa(int(praiseform.ID))+"回答点赞了哦")
}

func PraiseComment(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 2)
	flag1 := dao.SelectComment(id)
	global.RedisDb.Do("select", 0)
	if !flag1 {
		utils.RespFail(c, "Answer is not exist")
		return
	}
	// 验证用户是否点赞

	global.RedisDb.Do("select", 2)
	flag2 := dao.SelectPraiseuser(id, userId)
	global.RedisDb.Do("select", 0)
	// 用户点过赞则退出
	if flag2 {
		utils.RespFail(c, "您已点赞,无需重复点赞!")
		return
	}

	global.RedisDb.Do("select", 2)
	dao.Praiseadd(id, userId)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "点赞成功!")
}
func PraiseArticle(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 3)
	flag1 := dao.SelectArticle(id)
	global.RedisDb.Do("select", 0)
	if !flag1 {
		utils.RespFail(c, "Article is not exist")
		return
	}
	// 验证用户是否点赞
	global.RedisDb.Do("select", 3)
	flag2 := dao.SelectPraiseuser(id, userId)
	global.RedisDb.Do("select", 0)
	// 用户点过赞则退出
	if flag2 {
		utils.RespFail(c, "您已点赞,无需重复点赞!")
		return
	}
	global.RedisDb.Do("select", 3)
	dao.Praiseadd(id, userId)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "点赞成功!")
	ID := dao.GetIdFromUsername(username.(string))
	middleware.Producer("Subscribe", strconv.Itoa(int(ID))+" 给:"+strconv.Itoa(int(praiseform.ID))+"文章点赞了哦")
}
func PraiseArticleComment(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 4)
	flag1 := dao.SelectArticleComment(id)
	global.RedisDb.Do("select", 0)
	if !flag1 {
		utils.RespFail(c, "ArticleComment is not exist")
		return
	}
	// 验证用户是否点赞
	global.RedisDb.Do("select", 4)
	flag2 := dao.SelectPraiseuser(id, userId)
	global.RedisDb.Do("select", 0)
	// 用户点过赞则退出
	if flag2 {
		utils.RespFail(c, "您已点赞,无需重复点赞!")
		return
	}
	global.RedisDb.Do("select", 4)
	dao.Praiseadd(id, userId)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "点赞成功!")
}
func CancelPraiseQuestion(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))

	flag1 := dao.SelectQuestion(id)
	if !flag1 {
		utils.RespFail(c, "问题不存在")
		return
	}
	// 验证用户是否点赞
	flag2 := dao.SelectPraiseuser(id, userId)
	// 用户没点过赞则退出
	if !flag2 {
		utils.RespFail(c, "你都没点赞呢,想点踩是吧")
		return
	}
	dao.CancelPraise(id, userId)
	utils.RespSuccess(c, "取消点赞成功!")
}
func CancelPraiseArticle(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 3)
	flag1 := dao.SelectArticle(id)
	global.RedisDb.Do("select", 0)
	if !flag1 {
		utils.RespFail(c, "文章不存在")
		return
	}
	// 验证用户是否点赞
	global.RedisDb.Do("select", 3)
	flag2 := dao.SelectPraiseuser(id, userId)
	global.RedisDb.Do("select", 0)
	// 用户没点过赞则退出
	if !flag2 {
		utils.RespFail(c, "你都没点赞呢,想点踩是吧")
		return
	}
	global.RedisDb.Do("select", 3)
	dao.CancelPraise(id, userId)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "取消点赞成功!")
}
func CancelPraiseArticleComment(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 4)
	flag1 := dao.SelectArticleComment(id)
	global.RedisDb.Do("select", 0)
	if !flag1 {
		utils.RespFail(c, "文章评论不存在")
		return
	}
	// 验证用户是否点赞
	global.RedisDb.Do("select", 4)
	flag2 := dao.SelectPraiseuser(id, userId)
	global.RedisDb.Do("select", 0)
	// 用户没点过赞则退出
	if !flag2 {
		utils.RespFail(c, "你都没点赞呢,想点踩是吧")
		return
	}
	global.RedisDb.Do("select", 4)
	dao.CancelPraise(id, userId)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "取消点赞成功!")
}

func SeePraiseQuestion(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	number := dao.SeePraise(id)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"PraiseSum": number,
	})
}

func SeePraiseAnswer(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 1)
	number := dao.SeePraise(id)
	global.RedisDb.Do("select", 0)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"PraiseSum": number,
	})
}

func SeePraiseComment(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 2)
	number := dao.SeePraise(id)
	global.RedisDb.Do("select", 0)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"PraiseSum": number,
	})
}
func SeePraiseArticle(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 3)
	number := dao.SeePraise(id)
	global.RedisDb.Do("select", 0)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"PraiseSum": number,
	})
}
func SeePraiseArticleComment(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 4)
	number := dao.SeePraise(id)
	global.RedisDb.Do("select", 0)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"PraiseSum": number,
	})
}
func CancelPraiseAnswer(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 1)
	flag1 := dao.SelectAnswer(id)
	global.RedisDb.Do("select", 0)
	if !flag1 {
		utils.RespFail(c, "文章不存在")
		return
	}
	// 验证用户是否点赞
	global.RedisDb.Do("select", 1)
	flag2 := dao.SelectPraiseuser(id, userId)
	global.RedisDb.Do("select", 0)
	// 用户没点过赞则退出
	if !flag2 {
		utils.RespFail(c, "你都没点赞呢,想点踩是吧")
		return
	}
	global.RedisDb.Do("select", 1)
	dao.CancelPraise(id, userId)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "取消点赞成功!")
}
func CancelPraiseComment(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	id2 := praiseform.ID
	id := strconv.Itoa(int(id2))
	global.RedisDb.Do("select", 2)
	flag1 := dao.SelectComment(id)
	global.RedisDb.Do("select", 0)
	if !flag1 {
		utils.RespFail(c, "评论不存在")
		return
	}
	// 验证用户是否点赞
	global.RedisDb.Do("select", 2)
	flag2 := dao.SelectPraiseuser(id, userId)
	global.RedisDb.Do("select", 0)
	// 用户没点过赞则退出
	if !flag2 {
		utils.RespFail(c, "你都没点赞呢,想点踩是吧")
		return
	}
	global.RedisDb.Do("select", 2)
	dao.CancelPraise(id, userId)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "取消点赞成功!")
}

func JudgePraiseQuestion(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	flag := dao.SelectPraiseuser(strconv.Itoa(int(praiseform.ID)), userId)
	if flag {
		utils.RespSuccess(c, "点了赞")
		return
	} else {
		utils.RespFail(c, "没点赞")
		return
	}
}
func JudgePraiseAnswer(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	global.RedisDb.Do("select", 1)
	flag := dao.SelectPraiseuser(strconv.Itoa(int(praiseform.ID)), userId)
	global.RedisDb.Do("select", 1)
	if flag {
		utils.RespSuccess(c, "点了赞")
		return
	} else {
		utils.RespFail(c, "没点赞")
		return
	}
}
func JudgePraiseComment(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))

	global.RedisDb.Do("select", 2)
	flag := dao.SelectPraiseuser(strconv.Itoa(int(praiseform.ID)), userId)
	global.RedisDb.Do("select", 2)
	if flag {
		utils.RespSuccess(c, "点了赞")
		return
	} else {
		utils.RespFail(c, "没点赞")
		return
	}
}
func JudgePraiseArticle(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	global.RedisDb.Do("select", 3)
	flag := dao.SelectPraiseuser(strconv.Itoa(int(praiseform.ID)), userId)
	global.RedisDb.Do("select", 3)
	if flag {
		utils.RespSuccess(c, "点了赞")
		return
	} else {
		utils.RespFail(c, "没点赞")
		return
	}
}
func JudgePraiseArticleComment(c *gin.Context) {
	praiseform := model.Praise{}
	if err := c.ShouldBind(&praiseform); err != nil {
		utils.RespFail(c, "wrong submit")
		return
	}
	username, _ := c.Get("username")
	userId := dao.GetIdFromUsername(username.(string))
	global.RedisDb.Do("select", 4)
	flag := dao.SelectPraiseuser(strconv.Itoa(int(praiseform.ID)), userId)
	global.RedisDb.Do("select", 4)
	if flag {
		utils.RespSuccess(c, "点了赞")
		return
	} else {
		utils.RespFail(c, "没点赞")
		return
	}
}
