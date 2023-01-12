package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func (manager *ClientManager) Start() {
	for {
		fmt.Println("-----监听管道通信-----")
		select {
		case conn := <-Manager.Register:
			fmt.Printf("有新连接:%v ", conn.ID)
			Manager.Clients[conn.ID] = conn //把连接放到用户管理上
			ReplyMsg := ReplyMsg{
				Code:    50002,
				Content: "已连接服务器",
			}
			msg, _ := json.Marshal(ReplyMsg)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.Unregister:
			fmt.Printf("连接失败%s, ", conn.ID)
			if _, ok := Manager.Clients[conn.ID]; ok {
				ReplyMsg := &ReplyMsg{
					Code:    50003,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(ReplyMsg)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(Manager.Clients, conn.ID)
			}
		case broadcast := <-Manager.Broadcast:
			message := broadcast.Message
			sendId := broadcast.Client.SendID
			flag := false //默认对方是不在线的
			for id, conn := range Manager.Clients {
				if id != sendId {
					continue
				}
				select {
				case conn.Send <- message: //此时的conn是被发送消息的client
					flag = true
				default:
					close(conn.Send)
					delete(Manager.Clients, conn.ID)
				}
			}
			id := broadcast.Client.ID //1->2
			if flag {
				ReplyMsg := &ReplyMsg{
					Code:    50004,
					Content: "对方在线应答",
				}
				msg, _ := json.Marshal(ReplyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				//err := InsertMsg("zhihu", id, string(message), 1, int64(3*month)) //1 已经读了
				err := InsertMessage(id, string(message), 1)
				if err != nil {
					fmt.Println("InsertOne Err", err)
				}
			} else {
				fmt.Println("对方不在线")
				ReplyMsg := &ReplyMsg{
					Code:    50005,
					Content: "对方不在线应答",
				}
				msg, _ := json.Marshal(ReplyMsg)
				_ = broadcast.Client.Socket.WriteMessage(websocket.TextMessage, msg)
				err := InsertMessage(id, string(message), 0) //0 没有读
				if err != nil {
					fmt.Println("InsertOne Err", err)
				}
			}
		}
	}
}
