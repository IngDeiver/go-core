package ws

import (
	"github.com/ingdeiver/go-core/src/config"
	wsDomain "github.com/ingdeiver/go-core/src/ws/domain"
	wsHandlers "github.com/ingdeiver/go-core/src/ws/infrastructure/handlers"
)

// ------------ ws config ------------
func InitWsModule(){
	router := config.GetRouter()
	server := config.GetServer()


	webSocketDomain := wsDomain.New()
	webSocketManager := wsHandlers.New(webSocketDomain)
	server.SetWebSocketHandler(webSocketManager.Handler(), router)
}