package web

import (
	"fmt"
	"log"
	"mmt/example/pkg/media"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func Serve(port int, mediaSources media.MediaCollection, postsPerCycle int) error {
	websocketConnections := websocketConnections{
		Broadcast: make(chan media.Media),
	}
	mux := createHttpMux(&websocketConnections)

	l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return err
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			mediaSources.Collect(websocketConnections.Broadcast, postsPerCycle)
		}
	}()

	return http.Serve(l, mux)
}

func createHttpMux(websocketConnections *websocketConnections) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		u := websocket.Upgrader{}
		conn, err := u.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		log.Println("Opening websocket connection")
		websocketConnections.add(conn)
	})

	mux.Handle("/", http.FileServer(http.Dir("public")))

	return mux
}
