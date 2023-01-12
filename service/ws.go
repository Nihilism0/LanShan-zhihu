package service

import (
	"CSAwork/dao"
	"CSAwork/global"
	"CSAwork/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"

	"time"
)

const month = 60 * 60 * 24 * 30

type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}
type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}
type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}
type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}
type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func CreateID(uid, toUid string) string {
	return uid + "->" + toUid
}

func ChatHandler(c *gin.Context) {
	uid := c.Query("uid")
	toUid := c.Query("toUid")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//创建一个用户实例
	client := &Client{
		ID:     CreateID(uid, toUid), //1->2
		SendID: CreateID(toUid, uid), //2->1
		Socket: conn,
		Send:   make(chan []byte),
	}
	//用户注册到用户管理上
	Manager.Register <- client
	go client.Read()
	go client.Write()
}
func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg)
		if err != nil {
			fmt.Println("数据格式不正确")
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}
		if sendMsg.Type == 1 {
			r1, _ := global.RedisDb.Get(c.ID).Result()
			r2, _ := global.RedisDb.Get(c.SendID).Result()
			//防止添狗骚扰女神
			if r1 > "3" && r2 == "" {
				ReplyMsg := ReplyMsg{
					Code:    50006,
					Content: "达到3条限制",
				}
				msg, _ := json.Marshal(ReplyMsg) //序列化
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			} else {
				global.RedisDb.Incr(c.ID)
				_ = global.RedisDb.Expire(c.ID, time.Hour*24*30*3)
			}
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMsg.Content), //发过来的消息
			}
		} else if sendMsg.Type == 2 { //拉取历史消息
			//results, _ := FindMany("zhihu", c.SendID, c.ID, 10)
			results, err := FindManyMysql(c.SendID, c.ID)
			if err != nil {
				log.Println(err)
			}
			for _, result := range results {
				replyMsg := ReplyMsg{
					From:    result.From,
					Content: fmt.Sprintf("%s", result.Msg),
				}
				msg, _ := json.Marshal(replyMsg)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			ReplyMsg := ReplyMsg{
				Code:    50001,
				Content: fmt.Sprintf("%s", string(message)),
			}
			msg, _ := json.Marshal(ReplyMsg)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func SeeAllUnread(c *gin.Context) {
	username, _ := c.Get("username")
	Id := dao.GetIdFromUsername(username.(string))
	search := fmt.Sprint("%", strconv.Itoa(int(Id)))
	count := FindUnreadFunc(search)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"count":  count,
	})
}

type normal struct {
	ID uint `db:"id" form:"id" json:"id" binding:"required"`
}

func SeeTAUnread(c *gin.Context) {
	TAform := normal{}
	if err := c.ShouldBind(&TAform); err != nil {
		utils.RespFail(c, "Incorrect form are submitted!")
		return
	}
	username, _ := c.Get("username")
	Id := dao.GetIdFromUsername(username.(string))
	search := fmt.Sprint(TAform.ID, "->", strconv.Itoa(int(Id)))
	count := FindUnreadFunc(search)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"count":  count,
	})
}
