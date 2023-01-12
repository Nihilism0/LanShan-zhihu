package api

import (
	"CSAwork/dao"
	"CSAwork/model"
	"CSAwork/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddQuestionRecord(c *gin.Context) {
	form := model.Praise{}
	if err := c.ShouldBind(&form); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	id := form.ID
	ID := dao.GetIdFromUsername(username.(string))
	dao.AddQuestionRecord(ID, id)
	utils.RespSuccess(c, "add success")
}
func AddArticleRecord(c *gin.Context) {
	form := model.Praise{}
	if err := c.ShouldBind(&form); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	id := form.ID
	ID := dao.GetIdFromUsername(username.(string))
	dao.AddArticleRecord(ID, id)
	utils.RespSuccess(c, "add success")
}
func SeeRecord(c *gin.Context) {
	username, _ := c.Get("username")
	ID := dao.GetIdFromUsername(username.(string))
	records := dao.SeeRecords(ID)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"records": records,
	})
}

func DeleteAllRecords(c *gin.Context) {
	username, _ := c.Get("username")
	ID := dao.GetIdFromUsername(username.(string))
	dao.DeleteAllRecords(ID)
	utils.RespSuccess(c, "删除所有记录成功")
}
