package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

func main() {

	buffer := []byte("Hello, world.")

	start := time.Now()
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/", nil)

	fmt.Printf("err: %v\n", err)
	for i := 0; i < 1000000; i++ {
		c.WriteMessage(websocket.BinaryMessage, buffer)
	}
	delta := time.Since(start)

	fmt.Printf("Delta: %s\n", delta.String())
}
