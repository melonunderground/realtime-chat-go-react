package main

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"net/http"
// 	"os"
// 	"strconv"
// )

// func httphandler(w http.ResponseWriter, r *http.Request) {
// 	ipAddress, _, _ := net.SplitHostPort(r.RemoteAddr)
// 	fmt.Fprintf(w, "%s", ipAddress)
// }

// func main() {
// 	port, err := strconv.Atoi(os.Getenv("WHATISMYIP_PORT"))
// 	if err != nil {
// 		log.Fatalf("Please make sure the environment variable WHATISMYIP_PORT is defined and is a valid integer [1024-65535], error: %s", err)
// 	}

// 	listener := fmt.Sprintf(":%d", port)

// 	http.HandleFunc("/", httphandler)
// 	log.Fatal(http.ListenAndServe(listener, nil))
// }
