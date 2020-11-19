package web

import (
	"fmt"
	"log"
	"sync"

	"mmt/example/pkg/media"

	"github.com/gorilla/websocket"
)

type websocketConnections struct {
	sync.Mutex
	Broadcast   chan media.Media
	connections []*websocket.Conn
}

func (wc *websocketConnections) add(conn *websocket.Conn) {
	wc.Lock()
	defer wc.Unlock()

	if wc.connections == nil {
		go wc.broadcastingHandler()
	}

	wc.connections = append(wc.connections, conn)
}

func (wc *websocketConnections) broadcastingHandler() {
	for m := range wc.Broadcast {
		log.Println(fmt.Sprintf("Broadcasting message '%s' (%s)", m.Title, m.ImageUrl))

		var deadConnections []*websocket.Conn

		for _, conn := range wc.connections {
			err := conn.WriteJSON(m)
			if err != nil {
				deadConnections = append(deadConnections, conn)
			}
		}

		for _, conn := range deadConnections {
			wc.remove(conn)
		}
	}
}

func (wc *websocketConnections) remove(conn *websocket.Conn) {
	wc.Lock()
	defer wc.Unlock()

	for i, connection := range wc.connections {
		if conn == connection {
			log.Println("Removing a dead connection")
			wc.connections[i] = wc.connections[len(wc.connections)-1]
			wc.connections[len(wc.connections)-1] = nil
			wc.connections = wc.connections[:len(wc.connections)-1]
			break
		}
	}
}
