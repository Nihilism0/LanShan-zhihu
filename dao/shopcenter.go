package dao

import (
	"CSAwork/global"
	"CSAwork/model"
	"fmt"
	"log"
	"time"
)

func AddGoods(name, describe string, cost uint) {
	sqlStr := "insert into shopcenter(goods_name, cost, `describe`) values (?,?,?)"
	global.GlobalDb1.Exec(sqlStr, name, cost, describe)
}
func SeeGoods() []model.Goods {
	goods := []model.Goods{}
	sqlStr := "select id, goods_name, sales_volume, cost, `describe` from shopcenter"
	err := global.GlobalDb1.Select(&goods, sqlStr)
	if err != nil {
		log.Println(err)
	}
	return goods
}
func AddGod(username string, num uint) {
	sqlStr := "update users set gold=gold+ ? where username=?"
	global.GlobalDb1.Exec(sqlStr, num, username)
}
func JudgeBuy(username string, goods_id uint) bool {
	var user []model.UserInfo
	sqlStr := "select username from shoplist where goods_id = ? and	username= ? "
	err := global.GlobalDb1.Select(&user, sqlStr, goods_id, username)
	if err != nil {
		fmt.Println(err)
	}
	if user[0].UserName == username {
		return true
	} else {
		return false
	}
}
func JudgeCanBuy(username string, cost uint) bool {
	var Gold []model.User
	sqlStr := "select gold from users where username= ? "
	err := global.GlobalDb1.Select(&Gold, sqlStr, username)
	if err != nil {
		fmt.Println(err)
	}
	if Gold[0].Gold >= cost {
		return true
	} else {
		return false
	}
}
func Buy(username string, goods_id uint, cost uint) {
	TimeNow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "update shopcenter set sales_volume=sales_volume+1  where id = ?"
	_, err := global.GlobalDb1.Exec(sqlStr, goods_id)
	if err != nil {
		fmt.Println(err)
	}
	sqlStr = "update users set gold=gold- ? where username = ?"
	global.GlobalDb1.Exec(sqlStr, cost, username)
	sqlStr = "insert into shoplist(username, goods_id, time) values (?,?,?)"
	global.GlobalDb1.Exec(sqlStr, username, goods_id, TimeNow)
}
