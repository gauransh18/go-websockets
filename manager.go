package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

var (

	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {


}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new connection")

	conn, err := websocketUpgrader.Upgrade(w, r, nil) //upgrade rugular http connection to websocket connection
	if err != nil {
		log.Println(err)
		return
	}

	conn.Close()
}