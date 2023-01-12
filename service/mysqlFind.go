package service

import (
	"CSAwork/global"
	"CSAwork/model/ws"
	"log"
	"time"
)

func InsertMessage(id string, content string, read uint) error {
	Timenow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "insert into Chat(time, direction, `read`, content) values (?,?,?,?)"
	_, err := global.GlobalDb1.Exec(sqlStr, Timenow, id, read, content)
	if err != nil {
		return err
	}
	return nil
}

func FindManyMysql(sendId string, id string) (results []ws.Result, err error) {
	var resultsMe []ws.Trainer
	var resultsYou []ws.Trainer
	sqlStr := "select time, `read`, content from Chat where direction = ?"
	err = global.GlobalDb1.Select(&resultsMe, sqlStr, id)
	if err != nil {
		log.Println(err)
	}
	sqlStr = "update Chat set `read`=1 where direction = ?"
	_, err = global.GlobalDb1.Exec(sqlStr, sendId)
	if err != nil {
		return nil, err
	}
	sqlStr = "select time, `read`, content from Chat where direction = ?"
	err = global.GlobalDb1.Select(&resultsYou, sqlStr, sendId)
	if err != nil {
		return nil, err
	}
	results, err = AppendAndSort(resultsMe, resultsYou)
	return
}
func FindUnreadFunc(search string) uint {
	var count uint
	sqlStr := "select count(*) from Chat where direction LIKE ? and `read`=0"
	global.GlobalDb1.Get(&count, sqlStr, search)
	return count
}
