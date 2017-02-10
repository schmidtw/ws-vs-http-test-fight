package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{} // use default options

func drop(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		/*
			fmt.Printf("recv: %s", message)
			err = c.WriteMessage(mt, message)
			if err != nil {
				fmt.Println("write:", err)
				break
			}
		*/
	}
}

func main() {
	http.HandleFunc("/", drop)
	http.ListenAndServe(":8080", nil)
}
