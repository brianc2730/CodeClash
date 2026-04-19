package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// Test For Now
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error")
	}

	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error")
			break
		}

		fmt.Println(msg)
		err2 := conn.WriteMessage(msgType, msg)
		if err2 != nil {
			fmt.Println("Error")
			break
		}
	}
}

func main() {
	router := gin.Default()
	router.GET("/ws", wsHandler)
	router.Run()
}