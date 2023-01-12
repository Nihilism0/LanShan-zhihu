package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func (manager *SubsClientManager) Start() {
	for {
		fmt.Println("监听连接")
		select {
		case conn := <-Manager.Register:
			fmt.Printf("subs有新连接:%v ", conn.ID)
			Manager.Clients[conn.ID] = conn //把连接放到用户管理上
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
				delete(Manager.Clients, conn.ID)
			}
		}
	}
}
