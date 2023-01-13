package dao

import (
	"CSAwork/global"
	"fmt"
	"log"
)

func SelectQuestion(id string) bool {
	flag, err := global.RedisDb.SIsMember("questionids", id).Result()
	if err != nil {
		fmt.Println(err)
	}
	return flag
}

func SelectAnswer(id string) bool {
	flag, _ := global.RedisDb.SIsMember("answerids", id).Result()
	return flag
}

func SelectComment(id string) bool {
	flag, _ := global.RedisDb.SIsMember("commentids", id).Result()
	return flag
}

func Praiseadd(id string, userId uint) {
	global.RedisDb.SAdd(id, userId)
}

func SelectPraiseuser(id string, userId uint) bool {
	flag, _ := global.RedisDb.SIsMember(id, userId).Result()
	return flag
}
func CancelPraise(id string, userId uint) {
	global.RedisDb.SRem(id, userId)
}

func SeePraise(id string) int64 {
	result, err := global.RedisDb.SCard(id).Result()
	if err != nil {
		log.Fatalln(err)
	}
	return result
}
