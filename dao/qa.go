package dao

import (
	"CSAwork/global"
	"CSAwork/model"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

func FindQuestionSubmited(username string) []model.Question {
	var questions []model.Question
	sqlStr := "select id, title,created_at, updated_at, questioner, message from questions where questioner = ?"
	err := global.GlobalDb1.Select(&questions, sqlStr, username)
	if err != nil {
		panic(err)
	}
	var answers []model.Answer
	for k, v := range questions {
		sqlStr = "select id, created_at, updated_at,answerer, message, question_id from answers where question_id = ?"
		err = global.GlobalDb1.Select(&answers, sqlStr, v.ID)
		v.Answers = answers
		questions[k] = v
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("find question error!:", questions)
	//global.GlobalDb1.Model(&model.Question{}).Preload("Answers").Preload("Comments").Where("Questioner = ?", username).Find(&questions)
	return questions
}

func FindAnswerSubmited(username string) []model.Answer {
	var answers []model.Answer
	sqlStr := "select id, created_at, updated_at, answerer, message, question_id from answers where answerer = ?"
	_ = global.GlobalDb1.Select(&answers, sqlStr, username)
	var comments []model.Comment
	for k, v := range answers {
		sqlStr = "select id, created_at, updated_at,commenter, message, answer_id from comments where answer_id = ?"
		err := global.GlobalDb1.Select(&comments, sqlStr, v.ID)
		v.Comments = comments
		answers[k] = v
		if err != nil {
			log.Println("find answer error!:", err)
		}
	}
	//global.GlobalDb1.Model(&model.Answer{}).Preload("Comments").Where("Answerer = ?", username).Find(&answers)
	return answers
}

func JudgeQuestion(username string, id uint) bool {
	OK := false
	var question model.Question
	sqlStr := "select questioner from questions where id = ?"
	_ = global.GlobalDb1.Get(&question, sqlStr, id)
	//global.GlobalDb1.Model(&model.Question{}).Where("questioner = ? AND id = ?", username, id).First(&question)
	if question.Questioner == username {
		OK = true
	}
	return OK
}

func QuestionModify(message string, id uint) {
	sqlStr := "update questions set message=?  where id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, message, id)
	//global.GlobalDb1.Model(&model.Question{}).Where("id = ?", id).Update("message", message)
	timenow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr = "update questions set updated_at=? where id = ?"
	global.GlobalDb1.Exec(sqlStr, timenow, id)
}

func JudgeAnswer(username string, id uint) bool {
	OK := false
	var answer model.Answer
	sqlStr := "select answerer from answers where id = ?"
	_ = global.GlobalDb1.Get(&answer, sqlStr, id)
	//global.GlobalDb1.Model(&model.Answer{}).Where("answerer = ? AND id = ?", username, id).First(&answer)
	if answer.Answerer == username {
		OK = true
	}
	return OK
}

func AnswerModify(message string, id uint) {
	TimeNow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "update answers set message=? and updated_at=? where id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, message, TimeNow, id)
	//global.GlobalDb1.Model(&model.Answer{}).Where("id = ?", id).Update("message", message)
	timenow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr = "update answers set updated_at=? where id = ?"
	global.GlobalDb1.Exec(sqlStr, timenow, id)
}

func AnswerDelete(id uint) {
	//var answer model.Answer
	//var comments []model.Comment
	//sqlStr := "delete from comments where answer_id = ?"
	//_, _ = global.GlobalDb1.Exec(sqlStr, id)
	sqlStr := "delete from answers where id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, id)
	//global.GlobalDb1.Model(&model.Comment{}).Unscoped().Where("answer_id = ?", id).Delete(&comments)
	//global.GlobalDb1.Model(&model.Answer{}).Unscoped().Where("id = ?", id).Delete(&answer)
}
func QuestionDelete(id uint) {
	//var question model.Question
	//var answer []model.Answer
	//var comments []model.Comment

	//sqlStr := "delete from comments where answer_id = ?"
	//_, _ = global.GlobalDb1.Exec(sqlStr, id)
	//sqlStr = "delete from answers where question_id = ?"
	//_, _ = global.GlobalDb1.Exec(sqlStr, id)
	sqlStr := "delete from questions where id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, id)
	//global.GlobalDb1.Model(&model.Comment{}).Unscoped().Where("answer_id = ?", id).Delete(&comments)
	//global.GlobalDb1.Model(&model.Answer{}).Unscoped().Where("question_id = ?", id).Delete(&answer)
	//global.GlobalDb1.Model(&model.Question{}).Unscoped().Where("id = ?", id).Delete(&question)
}

func SearchAll(message string) ([]model.Question, []model.Answer, []model.Article) {
	var questions []model.Question
	var answers []model.Answer
	var articles []model.Article
	part := fmt.Sprint("%", message, "%")
	sqlStr := "select id, created_at, updated_at, questioner, title, message from questions where message LIKE ? or title LIKE ?"
	err := global.GlobalDb1.Select(&questions, sqlStr, part, part)
	if err != nil {
		fmt.Println(err)
	}
	sqlStr = "select id, created_at, updated_at,answerer, message, question_id from answers where message LIKE ?"
	err = global.GlobalDb1.Select(&answers, sqlStr, part)
	if err != nil {
		fmt.Println(err)
	}
	sqlStr = "select id, created_at, updated_at, articler, title, message from articles where message LIKE ? or title LIKE ? "
	err = global.GlobalDb1.Select(&articles, sqlStr, part, part)
	if err != nil {
		fmt.Println(err)
	}
	return questions, answers, articles
}
func SearchTitle(message string) ([]model.Question, []model.Article) {
	var questions []model.Question
	var articles []model.Article
	part := fmt.Sprint("%", message, "%")
	sqlStr := "select id, created_at, updated_at, questioner, title, message from questions where title LIKE ?"
	err := global.GlobalDb1.Select(&questions, sqlStr, part)
	if err != nil {
		fmt.Println(err)
	}
	sqlStr = "select id, created_at, updated_at, articler, title, message from articles where title LIKE ? "
	err = global.GlobalDb1.Select(&articles, sqlStr, part)
	if err != nil {
		fmt.Println(err)
	}
	return questions, articles
}
func GetHotAnswers() []model.AnswerInfo {
	var AllAnswers []model.AnswerInfo
	var HotAnswers []model.AnswerInfo
	sqlStr := "select id, created_at, answerer, message, question_id from answers"
	global.GlobalDb1.Select(&AllAnswers, sqlStr)
	global.RedisDb.Do("select", 1)
	for k, v := range AllAnswers {
		var sum int64
		sqlStr = "select count(*) from comments where answer_id=?"
		global.GlobalDb1.Get(&sum, sqlStr, v.ID)
		likes := SeePraise(strconv.Itoa(int(v.ID)))
		//我在知乎随便看了一篇文章3764赞同 174评论==>所以设赞同170热度,评论376热度
		AllAnswers[k].Hots = likes*170 + sum*376
	}
	global.RedisDb.Do("select", 0)
	sort.SliceStable(AllAnswers, func(i, j int) bool {
		return AllAnswers[i].Hots > AllAnswers[j].Hots
	})
	if len(AllAnswers) > 10 {
		HotAnswers = AllAnswers[:10]
	} else {
		HotAnswers = AllAnswers
	}
	return HotAnswers
}
func GetHotArticles() []model.ArticleInfo {
	var AllArticles []model.ArticleInfo
	var HotArticles []model.ArticleInfo
	sqlStr := "select id, articler, title, message from articles"
	global.GlobalDb1.Select(&AllArticles, sqlStr)
	global.RedisDb.Do("select", 3)
	for k, v := range AllArticles {
		var sum int64
		sqlStr = "select count(*) from article_comments where article_id=?"
		global.GlobalDb1.Get(&sum, sqlStr, v.ID)
		likes := SeePraise(strconv.Itoa(int(v.ID)))
		AllArticles[k].Hots = likes*170 + sum*376
	}
	global.RedisDb.Do("select", 0)
	sort.SliceStable(AllArticles, func(i, j int) bool {
		return AllArticles[i].Hots > AllArticles[j].Hots
	})
	if len(AllArticles) > 10 {
		HotArticles = AllArticles[:10]
	} else {
		HotArticles = AllArticles
	}
	return HotArticles
}
func GetHotQuestions() []model.QuestionInfo {
	var AllQuestions []model.QuestionInfo
	var HotQuestions []model.QuestionInfo
	sqlStr := "select id, questioner, title, message from questions"
	global.GlobalDb1.Select(&AllQuestions, sqlStr)
	for k, v := range AllQuestions {
		var sum int64
		sqlStr = "select count(*) from answers where question_id = ?"
		global.GlobalDb1.Get(&sum, sqlStr, v.ID)
		likes := SeePraise(strconv.Itoa(int(v.ID)))
		AllQuestions[k].Hots = likes*170 + sum*376
	}
	sort.SliceStable(AllQuestions, func(i, j int) bool {
		return AllQuestions[i].Hots > AllQuestions[j].Hots
	})
	if len(AllQuestions) > 10 {
		HotQuestions = AllQuestions[:10]
	} else {
		HotQuestions = AllQuestions
	}
	return HotQuestions
}
func FindQuestion(id uint) model.Question {
	var question model.Question
	sqlStr := "select id, title,created_at, updated_at, questioner, message from questions where id = ?"
	err := global.GlobalDb1.Get(&question, sqlStr, id)
	if err != nil {
		log.Println(err)
	}
	var answers = []model.Answer{}
	sqlStr = "select id, created_at,created_at, updated_at,answerer, message, question_id from answers where question_id = ?"
	err = global.GlobalDb1.Select(&answers, sqlStr, question.ID)
	for k, v := range answers {
		var comments = []model.Comment{}
		sqlStr = "select id, created_at, updated_at,commenter, message, answer_id from comments where answer_id = ?"
		err = global.GlobalDb1.Select(&comments, sqlStr, v.ID)
		answers[k].Comments = comments
		if err != nil {
			log.Println("find answer error!:", err)
		}
	}
	question.Answers = answers
	return question
}
func RandQuestion() []model.Question {
	questions := []model.Question{}
	sqlStr := "select id,title,created_at,questioner,message from questions order by rand()"
	err := global.GlobalDb1.Select(&questions, sqlStr)
	if err != nil {
		log.Println(err)
	}
	return questions
}
func GetQuestionFromAnswer(id uint) model.Question {
	var ID uint
	sqlStr := "select question_id from answers where id = ?"
	global.GlobalDb1.Get(&ID, sqlStr, id)
	var question model.Question
	sqlStr = "select id, created_at, updated_at, questioner, title, message from questions where id = ?"
	global.GlobalDb1.Get(&question, sqlStr, ID)
	return question
}
func GetUsernameFromQuestionId(id uint) string {
	var username string
	sqlStr := "select questioner from questions where id=?"
	global.GlobalDb1.Get(&username, sqlStr, id)
	return username
}
