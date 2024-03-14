package wshHandlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	wsDomain "github.com/ingdeiver/go-core/src/ws/domain"
)

var upgrader = websocket.Upgrader{}

type WebSocketHandlerManager struct {
	WsDomain wsDomain.WebSocketDomain
}

func New(wsDomain wsDomain.WebSocketDomain) WebSocketHandlerManager{
	return WebSocketHandlerManager{wsDomain}
}

func ( manager *WebSocketHandlerManager) handleConnection(conn *websocket.Conn) {
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			switch {
			case websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived):
				log.Println("The connection was closed")
				//then remove connection
			default:
				errorFormat := fmt.Errorf("read message error: %v", err)
				log.Println(errorFormat)
			}

			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			errorFormat := fmt.Errorf("write message error: %v", err)
			log.Println(errorFormat)
			return
		}
		fmt.Printf("Received message: '%s' of type: %v \n", p, messageType)
	}
}

func (manager *WebSocketHandlerManager) addNewConnection(conn *websocket.Conn) {
	manager.WsDomain.Connections = append(manager.WsDomain.Connections, conn)
	go manager.handleConnection(conn)
}

func (manager *WebSocketHandlerManager) Handler() func(*gin.Context) {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			errorFormat := fmt.Errorf("new connection error => %v", err)
			log.Println(errorFormat)
			return
		}

		manager.addNewConnection(conn)
	}
}