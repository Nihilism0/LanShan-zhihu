package service

import (
	"CSAwork/dao"
	"CSAwork/global"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/nsqio/go-nsq"
	"strings"
	"time"
)

// NSQ Consumer Demo

// MyHandler 是一个消费者类型
type MyHandler struct {
	Title string
}

// HandleMessage 是需要实现的处理消息的方法
func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	TwoString := strings.SplitN(string(msg.Body), " ", 2)
	global.RedisDb.Do("select", 5)
	followers := dao.GetFollowersByString(TwoString[0])
	global.RedisDb.Do("select", 0)
	for _, v := range followers {
		flag := false
		for _, client := range SubsManager.SubsClients {
			if client.ID != v {
				continue
			}
			flag = true
			ReplyMsg := &ReplyMsg{
				Code:    50001,
				Content: string(msg.Body),
			}
			message, _ := json.Marshal(ReplyMsg)
			client.Socket.WriteMessage(websocket.TextMessage, message)
			sqlStr := "insert into Subscribe(sender_id, message, time, receiver_id,`read`)  values (?,?,?,?,?)"
			global.GlobalDb1.Exec(sqlStr, TwoString[0], string(msg.Body), time.Now(), v, 1)
			break
		}
		if flag == false {
			sqlStr := "insert into Subscribe(sender_id, message, time, receiver_id,`read`)  values (?,?,?,?,?)"
			global.GlobalDb1.Exec(sqlStr, TwoString[0], string(msg.Body), time.Now(), v, 0)
		}
	}
	return
}

// 初始化消费者
func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	consumer := &MyHandler{
		Title: "订阅消息",
	}
	c.AddHandler(consumer)
	if err := c.ConnectToNSQD(address); err != nil { // 直接连NSQD
		//if err := c.ConnectToNSQLookupd(address); err != nil { // 通过lookupd查询
		return err
	}
	return nil
}

func (manager *SubsClientManager) ConsumerOfSubscribe(topic, channel string) {
	err := initConsumer(topic, channel, "49.234.42.190:4150")
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}
	for {
		fmt.Println("监听连接")
		select {
		case conn := <-SubsManager.Register:
			fmt.Printf("subs有新连接:%v ", conn.ID)
			SubsManager.SubsClients[conn.ID] = conn //把连接放到用户管理上
			ReplyMsg := ReplyMsg{
				Code:    50002,
				Content: "subs已连接服务器",
			}
			msg, _ := json.Marshal(ReplyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-SubsManager.Unregister:
			fmt.Printf("subs连接中断%s, ", conn.ID)
			if _, ok := SubsManager.SubsClients[conn.ID]; ok {
				ReplyMsg := &ReplyMsg{
					Code:    50003,
					Content: "subs连接中断",
				}
				msg, _ := json.Marshal(ReplyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				delete(SubsManager.SubsClients, conn.ID)
			}
		}
	}
}
