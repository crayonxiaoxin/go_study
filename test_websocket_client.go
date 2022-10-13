// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/url"
// 	"os"
// 	"os/signal"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// func main() {

// 	interrupt := make(chan os.Signal, 1)
// 	signal.Notify(interrupt, os.Interrupt)

// 	u := url.URL{Scheme: "ws", Host: "localhost:8081", Path: "/websocket"}
// 	fmt.Printf("connecting to : %v\n", u.String())
// 	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
// 	if err != nil {
// 		log.Fatal("dial:", err)
// 	}
// 	defer c.Close()

// 	done := make(chan struct{})
// 	go func() {
// 		defer close(done)
// 		for {
// 			_, message, err := c.ReadMessage()
// 			if err != nil {
// 				return
// 			}
// 			fmt.Printf("received: %v\n", string(message))
// 		}
// 	}()

// 	ticker := time.NewTicker(time.Second)
// 	// 每秒发送一条信息
// 	for {
// 		select {
// 		case <-done: // 关闭
// 			return
// 		case t := <-ticker.C:
// 			fmt.Printf("t: %v\n", t)
// 			err := c.WriteMessage(websocket.TextMessage, []byte("hello"))
// 			if err != nil {
// 				fmt.Printf("write close: %v\n", err)
// 				return
// 			}
// 		case <-interrupt:
// 			log.Println("interrupt")

// 			// Cleanly close the connection by sending a close message and then
// 			// waiting (with timeout) for the server to close the connection.
// 			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
// 			if err != nil {
// 				log.Println("write close:", err)
// 				return
// 			}
// 			select {
// 			case <-done:
// 			case <-time.After(time.Second):
// 			}
// 			return
// 		}

// 	}

// }
