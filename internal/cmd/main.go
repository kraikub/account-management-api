package main

import (
	"github.com/kraikub/account-manager-api/internal/config"
	"github.com/kraikub/account-manager-api/internal/controllers/handlers"
	"github.com/kraikub/account-manager-api/internal/controllers/routers"
	"github.com/kraikub/account-manager-api/internal/servers"
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