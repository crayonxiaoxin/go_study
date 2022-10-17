// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	// websocket: request origin not allowed by Upgrader.CheckOrigin
// 	CheckOrigin: func(r *http.Request) bool { return true },
// }

// func ws(c *gin.Context) {
// 	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		return
// 	}
// 	defer ws.Close()
// 	for {
// 		messageType, message, err := ws.ReadMessage()
// 		if err != nil {
// 			fmt.Printf("err: %v\n", err)
// 			break
// 		}
// 		if string(message) == "ping" {
// 			message = []byte("pong")
// 			ws.WriteMessage(messageType, message)
// 		} else if string(message) == "hello" {
// 			ws.WriteJSON(gin.H{"hello": "go websocket"})
// 		} else {
// 			ws.WriteMessage(messageType, message)
// 		}
// 	}

// }

// func main() {
// 	r := gin.Default()

// 	// ws://localhost:8081/websocket
// 	r.GET("/websocket", ws)

// 	r.GET("/api/:test", func(ctx *gin.Context) {
// 		ctx.JSON(200, gin.H{
// 			"rc":   0,
// 			"msg":  "success",
// 			"data": ctx.Param("test"),
// 		})
// 	})

// 	r.Run(":8081") // 监听并在 0.0.0.0:8081 上启动服务
// }
