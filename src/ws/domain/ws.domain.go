package webSocketDomain

import (
	"github.com/gorilla/websocket"
)

type WebSocketDomain struct {
	Connections []*websocket.Conn
}

func New() WebSocketDomain {
	connections := []*websocket.Conn{}
	return WebSocketDomain{Connections: connections}
}
