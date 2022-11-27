package main

import (
	"context"
	"fmt"

	"github.com/kraikub/account-management-api/api/v1/internal/config"
	"github.com/kraikub/account-management-api/api/v1/internal/controllers"
	"github.com/kraikub/account-management-api/servers"
)

func main() {

	config, err := config.GetRuntimeConfig()
	if err != nil {
		panic(err)
	}

	kraikub := servers.NewKraikubServer(config.Server.Name, config.Server.Port)
	controllers.AssignRouter(kraikub.Router())

	// No need any go routines
	kraikub.StartWithGraceFullShutdown(func(cancel context.CancelFunc) {
		fmt.Println("stop!")
		cancel()
	})
}
