package main

import (
	"github.com/kraikub/account-management-api/api/v1/controllers/handlers"
	"github.com/kraikub/account-management-api/api/v1/controllers/routers"
	"github.com/kraikub/account-management-api/config"
	"github.com/kraikub/account-management-api/servers"
)

func main() {

	config, err := config.GetRuntimeConfig()
	if err != nil {
		panic(err)
	}

	kraikub := servers.NewKraikubServer(config)
	accountHandler := handlers.NewAccountHandler()
	routers.RegisterAccountRouter(kraikub.Router(), accountHandler)

	// No need any go routines
	kraikub.Start()
}
