package main

import (
	"context"

	"github.com/kraikub/account-management-api/api/v1/internal/config"
	"github.com/kraikub/account-management-api/api/v1/internal/controllers"
	"github.com/kraikub/account-management-api/api/v1/internal/repositories"
	"github.com/kraikub/account-management-api/api/v1/internal/usecases"
	"github.com/kraikub/account-management-api/servers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	config, err := config.GetRuntimeConfig()
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithCancel(context.Background())

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Db.Mongo.Uri))
	if err != nil {
		panic(err)
	}

	db := mongoClient.Database(config.Db.Mongo.Name)

	userRepository := repositories.CreateUserRepository(db)

	userUseCase := usecases.CreateUserUseCase(userRepository)

	kraikub := servers.NewKraikubServer(config.Server.Name, config.Server.Port)
	controllers.AssignRouter(
		kraikub.Router(),
		userUseCase,
	)

	// No need any go routines
	kraikub.StartWithGraceFullShutdown(func(cancel context.CancelFunc) {
		// if err = mongoClient.Disconnect(context.TODO()); err != nil {
		// should not panic
		// 	log.Fatal(err)
		// }
		cancel()
	})
}
