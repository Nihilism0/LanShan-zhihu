package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type SubsClient struct {
	ID     string
	Socket *websocket.Conn
}
type SubsClientManager struct {
	SubsClients map[string]*SubsClient
	Register    chan *SubsClient
	Unregister  chan *SubsClient
}

var SubsManager = SubsClientManager{
	SubsClients: make(map[string]*SubsClient),
	Register:    make(chan *SubsClient),
	Unregister:  make(chan *SubsClient),
}

func SubsHandler(c *gin.Context) {
	id := c.Query("id")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	subsClient := &SubsClient{
		ID:     id,
		Socket: conn,
	}
	SubsManager.Register <- subsClient
	go subsClient.Blocking()
}
func (subsClient *SubsClient) Blocking() {
	defer func() {
		SubsManager.Unregister <- subsClient
		_ = subsClient.Socket.Close()
	}()
	for {
		subsClient.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := subsClient.Socket.ReadJSON(&sendMsg)
		if err != nil {
			break
		}
	}
}
