package wshHandlers

import (
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	logger "github.com/ingdeiver/go-core/src/commons/infrastructure/logs"
	wsDomain "github.com/ingdeiver/go-core/src/ws/domain"
)

var l = logger.Get()
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
        originAllowed := os.Getenv("CORS_ORIGIN")
		origin := r.Header.Get("Origin")

		if originAllowed == "*" || origin == originAllowed {
            return true
        }

        return false
    },
}

type WebSocketController struct {
	WsDomain *wsDomain.WebSocketDomain
	Mutex    sync.Mutex
}

func New(wsDomain *wsDomain.WebSocketDomain) WebSocketController{
	return WebSocketController{WsDomain: wsDomain}
}


func (manager *WebSocketController) handleConnection(conn *websocket.Conn) {
	defer func() {
		manager.removeConnection(conn)
		conn.Close()
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			switch {
			case websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived):
				l.Info().Msg("The connection was closed")
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


func (manager *WebSocketController) addNewConnection(conn *websocket.Conn) {
	manager.Mutex.Lock()
	manager.WsDomain.Connections = append(manager.WsDomain.Connections, conn)
	manager.Mutex.Unlock()
	go manager.handleConnection(conn)
}


func (manager *WebSocketController) removeConnection(conn *websocket.Conn) {
	manager.Mutex.Lock()
	defer manager.Mutex.Unlock()
	for i, c := range manager.WsDomain.Connections {
		if c == conn {
			manager.WsDomain.Connections = append(manager.WsDomain.Connections[:i], manager.WsDomain.Connections[i+1:]...)
			break
		}
	}
}

func (manager *WebSocketController) Handler() func(*gin.Context) {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			l.Error().Msgf("new connection error => %v", err)
			return
		}

		manager.addNewConnection(conn)
	}
}
