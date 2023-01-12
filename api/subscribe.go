package api

import (
	"CSAwork/dao"
	"CSAwork/global"
	"CSAwork/model"
	"CSAwork/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SubscribePeople(c *gin.Context) {
	form := model.Praise{}
	if err := c.ShouldBind(&form); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	id := dao.GetIdFromUsername(username.(string))
	global.RedisDb.Do("select", 5)
	dao.SubsPeople(form.ID, id)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "subs success")
}
func CancelSubscribePeople(c *gin.Context) {
	form := model.Praise{}
	if err := c.ShouldBind(&form); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	id := dao.GetIdFromUsername(username.(string))
	global.RedisDb.Do("select", 5)
	dao.CancelSubsPeople(form.ID, id)
	global.RedisDb.Do("select", 0)
	utils.RespSuccess(c, "cancel subs success")
}
func SeePeopleFollowers(c *gin.Context) {
	form := model.Praise{}
	if err := c.ShouldBind(&form); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	global.RedisDb.Do("select", 5)
	num := dao.SeePeopleFollowerNum(form.ID)
	followers := dao.GetFollowers(form.ID)
	global.RedisDb.Do("select", 0)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"count":     num,
		"followers": followers,
	})
}
func JudgeFollower(c *gin.Context) {
	form := model.Praise{}
	if err := c.ShouldBind(&form); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")

	id := dao.GetIdFromUsername(username.(string))
	global.RedisDb.Do("select", 5)
	flag := dao.JudgeFollower(form.ID, id)
	global.RedisDb.Do("select", 0)
	if flag {
		utils.RespSuccess(c, "已关注")
	} else {
		utils.RespFail(c, "未关注")
	}
}
func SeeMySubs(c *gin.Context) {
	username, _ := c.Get("username")
	ID := dao.GetIdFromUsername(username.(string))
	global.RedisDb.Do("select", 5)
	ids := dao.GetSubsFromID(ID)
	global.RedisDb.Do("select", 0)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"ids":    ids,
	})
}
