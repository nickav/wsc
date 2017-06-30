package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [URL]", os.Args[0])
		os.Exit(1)
	}
	ws := connect(os.Args[1])
	ws.Write([]byte(os.Args[2]))
	defer ws.Close()
	os.Exit(0)
}

func connect(addr string) *websocket.Conn {
	ws, err := websocket.Dial(addr, "", addr)
	if err != nil {
		log.Fatal(err)
	}
	return ws
}