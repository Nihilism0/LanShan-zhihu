package api

import (
	"CSAwork/dao"
	"CSAwork/global"
	"CSAwork/model"
	"CSAwork/utils"
	"CSAwork/utils/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Qcreate(c *gin.Context) {
	global.Bucket.Take(1)
	question := model.Question{}
	if err := c.ShouldBind(&question); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	//TestQuestion := model.Question{
	//	Questioner: username.(string),
	//	Message:    question.Message,
	//}
	TimeNow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "insert into questions(created_at,updated_at,questioner,title,message) values (?,?,?,?,?)"
	ret, err := global.GlobalDb1.Exec(sqlStr, TimeNow, TimeNow, username.(string), question.Title, question.Message)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	//global.GlobalDb1.Model(&model.Question{}).Create(&TestQuestion)

	id := strconv.Itoa(int(theID))
	err = global.RedisDb.SAdd("questionids", id).Err()
	if err != nil {
		log.Fatal("add error", zap.Error(err))
	}
	dao.AddGod(username.(string), 3)
	utils.RespSuccess(c, "亲爱的"+username.(string)+",您成功提问了一条,GO币加3==>"+question.Message)
	ID := dao.GetIdFromUsername(username.(string))
	middleware.Producer("Subscribe", strconv.Itoa(int(ID))+" 提了个问题:"+question.Title)
}

func Acreate(c *gin.Context) {
	global.Bucket.Take(1)
	answer := model.Answer{}
	if err := c.ShouldBind(&answer); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	message := answer.Message
	questionID := answer.Question_id
	TimeNow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "insert into answers(created_at, updated_at, answerer, message, question_id) values (?,?,?,?,?)"
	ret, _ := global.GlobalDb1.Exec(sqlStr, TimeNow, TimeNow, username.(string), message, questionID)
	TheId, _ := ret.LastInsertId()
	global.RedisDb.Do("select", 1)
	global.RedisDb.SAdd("answerids", TheId)
	global.RedisDb.Do("select", 0)
	//global.GlobalDb1.Model(&model.Answer{}).Create(&TestAnswer)
	dao.AddGod(username.(string), 5)
	utils.RespSuccess(c, "亲爱的"+username.(string)+",您成功回答了问题,GO币+5==>"+message)
	ID := dao.GetIdFromUsername(username.(string))
	middleware.Producer("Subscribe", strconv.Itoa(int(ID))+" 发布了个回答,快去看看")
}

func Qsubmited(c *gin.Context) {
	username, _ := c.Get("username")
	questions := dao.FindQuestionSubmited(username.(string))
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": questions,
	})
}

func Asubmited(c *gin.Context) {
	username, _ := c.Get("username")
	answers := dao.FindAnswerSubmited(username.(string))
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": answers,
	})
}

func Qmodify(c *gin.Context) {
	questionmodify := model.QuestionModify{}
	if err := c.ShouldBind(&questionmodify); err != nil {
		utils.RespFail(c, "Incorrect form are fucking submitted!")
		return
	}
	username, _ := c.Get("username")
	question := questionmodify.Message
	questionid := questionmodify.ID
	OK := dao.JudgeQuestion(username.(string), questionid)
	if !OK {
		utils.RespFail(c, "这个问题不属于你")
		return
	}
	dao.QuestionModify(question, questionid)
	utils.RespSuccess(c, "修改问题成功!!!")
}

func Amodify(c *gin.Context) {
	modifyform := model.AnswerModify{}
	if err := c.ShouldBind(&modifyform); err != nil {
		utils.RespFail(c, "Incorrect form are fucking submitted!")
		return
	}
	username, _ := c.Get("username")
	answer := modifyform.Message
	answerid := modifyform.ID
	OK := dao.JudgeAnswer(username.(string), answerid)
	if !OK {
		utils.RespFail(c, "这个回答不属于你")
		return
	}
	dao.AnswerModify(answer, answerid)
	utils.RespSuccess(c, "修改回答成功!!!")
}

func Qdelete(c *gin.Context) {
	deleteform := model.Delete{}
	if err := c.ShouldBind(&deleteform); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	questionid := deleteform.ID
	OK := dao.JudgeQuestion(username.(string), questionid)
	if !OK {
		utils.RespFail(c, "这个问题不属于你")
		return
	}
	dao.QuestionDelete(questionid)
	utils.RespSuccess(c, "删除问题成功!!!")
}

func Adelete(c *gin.Context) {
	deleteform := model.Delete{}
	if err := c.ShouldBind(&deleteform); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	answerid := deleteform.ID
	OK := dao.JudgeAnswer(username.(string), answerid)
	if !OK {
		utils.RespFail(c, "这个回答不属于你")
		return
	}
	dao.AnswerDelete(answerid)
	dao.CancelAnswerInFavorites(answerid)
	utils.RespSuccess(c, "删除回答成功!!!")
}

func Acomment(c *gin.Context) {
	global.Bucket.Take(1)
	comment := model.Comment{}
	if err := c.ShouldBind(&comment); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	message := comment.Message
	answerid := comment.Answer_id
	TimeNow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "insert into comments(created_at, updated_at, commenter, message, answer_id) values (?,?,?,?,?)"
	ret, _ := global.GlobalDb1.Exec(sqlStr, TimeNow, TimeNow, username.(string), message, answerid)
	TheId, _ := ret.LastInsertId()
	//global.GlobalDb1.Model(&model.Comment{}).Create(&TestComment)
	global.RedisDb.Do("select", 2)
	global.RedisDb.SAdd("commentids", TheId)
	global.RedisDb.Do("select", 0)
	dao.AddGod(username.(string), 1)
	utils.RespSuccess(c, "亲爱的"+username.(string)+"用户,您的评论已发送.GO币+1==>"+message)
}
func SearchTitle(c *gin.Context) {
	search := model.Search{}
	if err := c.ShouldBind(&search); err != nil {
		fmt.Println(err)
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	message := search.Message
	questions, articles := dao.SearchTitle(message)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"questions": questions,
		"articles":  articles,
	})
}
func SearchAll(c *gin.Context) {
	search := model.Search{}
	if err := c.ShouldBind(&search); err != nil {
		fmt.Println(err)
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	message := search.Message
	questions, answers, articles := dao.SearchAll(message)
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"questions": questions,
		"answers":   answers,
		"articles":  articles,
	})
}
func HotAnswer(c *gin.Context) {
	answers := dao.GetHotAnswers()
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"answers": answers,
	})
}
func HotQuestion(c *gin.Context) {
	questions := dao.GetHotQuestions()
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"questions": questions,
	})
}
func HotArticle(c *gin.Context) {
	articles := dao.GetHotArticles()
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"articles": articles,
	})
}
func SeeQuestionInformation(c *gin.Context) {
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
	question := dao.FindQuestion(id)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"question": question,
	})
}
func RandQuestions(c *gin.Context) {
	questions := dao.RandQuestion()
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"random": questions,
	})
}
func GetQuestionFromAnswer(c *gin.Context) {
	form := model.Praise{}
	if err := c.ShouldBind(&form); err != nil {
		fmt.Println(err)
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	question := dao.GetQuestionFromAnswer(form.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"question": question,
	})
}
