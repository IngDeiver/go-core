package wshHandlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	wsDomain "github.com/ingdeiver/go-core/src/ws/domain"
)
var l = logger.Get()
var upgrader = websocket.Upgrader{}

type WebSocketHandlerManager struct {
	WsDomain *wsDomain.WebSocketDomain
}

func New(wsDomain *wsDomain.WebSocketDomain) WebSocketHandlerManager{
	return WebSocketHandlerManager{wsDomain}
}

func ( manager *WebSocketHandlerManager) handleConnection(conn *websocket.Conn) {
	defer conn.Close()
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			switch {
			case websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived):
				l.Info().Msg("The connection was closed")
				//then remove connection
			default:
				l.Error().Msgf("read message error: %v", err)
			}

			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			l.Error().Msgf("write message error: %v", err)
			return
		}
		l.Info().Msgf("Received message: '%s' of type: %v \n", p, messageType)
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
			l.Error().Msgf("new connection error => %v", err)
			return
		}

		manager.addNewConnection(conn)
	}
}