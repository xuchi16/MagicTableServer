package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			// TODO: Check Origin before goes to production
			return true
		},
	}
	connections = make(map[*websocket.Conn]bool, 0)
	lock        sync.Mutex
)

func broadcastToAll(messageType int, message []byte) {
	lock.Lock()
	defer lock.Unlock()
	for conn := range connections {
		if err := conn.WriteMessage(messageType, message); err != nil {
			log.Println("Write message error:", err)
			conn.Close()
			delete(connections, conn)
		}
	}
}

func WebSocketHandle(c *gin.Context) {
	w := c.Writer
	r := c.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	lock.Lock()
	connections[conn] = true
	lock.Unlock()

	defer func() {
		lock.Lock()
		delete(connections, conn)
		lock.Unlock()
		conn.Close()
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read message error:", err)
			break
		}
		broadcastToAll(messageType, p)
	}
}

func SendMessage(c *gin.Context) {

}

func GetMessage(c *gin.Context) {

}
