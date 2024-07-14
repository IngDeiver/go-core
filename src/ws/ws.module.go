package ws

import (
	"github.com/ingdeiver/go-core/src/config"
	domain "github.com/ingdeiver/go-core/src/ws/domain"
	constrollers "github.com/ingdeiver/go-core/src/ws/infrastructure/controllers"
)

// ------------ ws config ------------
func InitWsModule(){
	router := config.GetRouter()
	server := config.GetServer()

	webSocketDomain := domain.New()
	webSocketController := constrollers.New(webSocketDomain)
	server.SetWebSocketHandler(webSocketController.Handler(), router)
}