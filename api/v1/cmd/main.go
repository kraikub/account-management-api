package main

import (
	"fmt"

	"github.com/kraikub/account-management-api/api/v1/config"
	"github.com/kraikub/account-management-api/api/v1/controllers/handlers"
	"github.com/kraikub/account-management-api/api/v1/controllers/routers"
	"github.com/kraikub/account-management-api/servers"
)

func main() {

	config, err := config.GetRuntimeConfig()
	if err != nil {
		panic(err)
	}

	kraikub := servers.NewKraikubServer(config.Server.Name, config.Server.Port)
	accountHandler := handlers.NewAccountHandler()
	routers.InitRouter(kraikub.Router(), accountHandler)

	// No need any go routines
	kraikub.StartWithGraceFullShutdown(func() {
		fmt.Println("stop!")
	})
}
