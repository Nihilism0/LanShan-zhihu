package dao

import (
	"CSAwork/global"
	"CSAwork/model"
	"fmt"
	"time"
)

func AddQuestionRecord(user_id, question_id uint) {
	var test uint
	sqlStr := "select user_id from record where user_id = ? and question_id = ?"
	global.GlobalDb1.Get(&test, sqlStr, user_id, question_id)
	if test != user_id {
		timenow := time.Now().Format("2006-01-02 15:04:05")
		sqlStr = "insert into record(user_id, time, question_id, article_id) values (?,?,?,?)"
		global.GlobalDb1.Exec(sqlStr, user_id, timenow, question_id, 0)
	}
}
func AddArticleRecord(user_id, article_id uint) {
	var test uint
	sqlStr := "select user_id from record where user_id = ? and article_id = ?"
	global.GlobalDb1.Get(&test, sqlStr, user_id, article_id)
	if test != user_id {
		timenow := time.Now().Format("2006-01-02 15:04:05")
		sqlStr = "insert into record(user_id, time, question_id, article_id) values (?,?,?,?)"
		global.GlobalDb1.Exec(sqlStr, user_id, timenow, 0, article_id)
	}
}
func SeeRecords(id uint) []model.Record {
	var records []model.Record
	sqlStr := "select user_id, time, question_id, article_id from record where user_id = ?"
	err := global.GlobalDb1.Select(&records, sqlStr, id)
	if err != nil {
		fmt.Println(err)
	}
	return records
}
func DeleteAllRecords(id uint) {
	sqlStr := "delete from record where user_id = ?"
	global.GlobalDb1.Exec(sqlStr, id)
}
