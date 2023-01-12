package dao

import (
	"CSAwork/global"
	"log"
	"strconv"
)

func SubsPeople(hostId, followerId uint) {
	global.RedisDb.SAdd(strconv.Itoa(int(hostId)), followerId)
	global.RedisDb.SAdd("hostids", hostId)
}
func CancelSubsPeople(hostId, followerId uint) {
	global.RedisDb.SRem(strconv.Itoa(int(hostId)), followerId)
}

func SeePeopleFollowerNum(hostId uint) int64 {
	result, err := global.RedisDb.SCard(strconv.Itoa(int(hostId))).Result()
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

func JudgeFollower(hostId, followerId uint) bool {
	flag, _ := global.RedisDb.SIsMember(strconv.Itoa(int(hostId)), followerId).Result()
	return flag
}
func GetFollowers(hostId uint) []string {
	followers, _ := global.RedisDb.SMembers(strconv.Itoa(int(hostId))).Result()
	return followers
}
func GetFollowersByString(hostId string) []string {
	followers, _ := global.RedisDb.SMembers(hostId).Result()
	return followers
}
func GetSubsFromID(id uint) []string {
	var realids []string
	hostIds, _ := global.RedisDb.SMembers("hostids").Result()
	flag := false
	for _, v := range hostIds {
		flag, _ = global.RedisDb.SIsMember(v, id).Result()
		if flag {
			realids = append(realids, v)
		}
	}
	return realids
}
