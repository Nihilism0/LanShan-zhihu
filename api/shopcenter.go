package api

import (
	"CSAwork/dao"
	"CSAwork/global"
	"CSAwork/model"
	"CSAwork/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGoods(c *gin.Context) {
	var addgoods model.AddGoods
	if err := c.ShouldBind(&addgoods); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	if addgoods.Secret != "123456" {
		utils.RespFail(c, "secret error!")
		return
	}
	dao.AddGoods(addgoods.GoodsName, addgoods.Describe, addgoods.Cost)
	utils.RespSuccess(c, "Add goods success!")
}

func SeeGoods(c *gin.Context) {
	goods := dao.SeeGoods()
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"goods":  goods,
	})
}
func BuyVip(c *gin.Context) {
	username, _ := c.Get("username")
	if dao.JudgeBuy(username.(string), 1) {
		utils.RespFail(c, "You already buy it.")
		return
	}
	if dao.JudgeCanBuy(username.(string), 100) == false {
		utils.RespFail(c, "Your GO Coin is not enough!")
		return
	}
	dao.Buy(username.(string), 1, 100)
	sqlStr := "update users set vip=1 where username = ?"
	global.GlobalDb1.Exec(sqlStr, username.(string))
	utils.RespSuccess(c, "You get VIP NOW!!!")
}
