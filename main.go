package main

import (
	"CSAwork/api"
	"CSAwork/boot"
	"CSAwork/service"
)

func main() {
	boot.ViperSetup("./config/config.yaml")
	boot.LoggerSetup()
	boot.InitDB()
	boot.RedisSetup()
	go service.Manager.Start()
	go service.SubsManager.ConsumerOfSubscribe("Subscribe", "1")
	api.InitRouter()
}

//boot.MongoDB()
