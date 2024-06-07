package main

import (
	"github.com/ingdeiver/go-core/src/auth"

	"github.com/ingdeiver/go-core/src/config"
	"github.com/ingdeiver/go-core/src/emails"

	"github.com/ingdeiver/go-core/src/users"

	"github.com/ingdeiver/go-core/src/ws"
)




func main() {
	initConfig()
	initApp()
}


func initConfig(){
	config.LoadEnv()
	config.InitMongoDB()
	config.CreateRouter()
	config.CreateServer()
}

func initApp(){
	

 	// -----------------------  1. load server config --------------------------------
	server := config.GetServer()
	router := config.GetRouter()
	
	//------------ static files ------------
	server.ConfigureStaticFiles("public", router) // optional, if don't need it, you can remove this line
	
	// ------------ set middlewares ------------
	server.ConfigGlobalMiddlewares(router)

	// ------------------------ 2. load modules ------------------------------------------------
	ws.InitWsModule() // optional, if don't need it, you can remove this line
	users.InitUsersModule()
	emails.InitEmailsModule()
	auth.InitAuthModule()

	// ------------------------- 3. start server -------------------------------------------
	server.StartServer()
}


