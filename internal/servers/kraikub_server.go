package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kraikub/account-manager-api/internal/config"
)

type kraikubServer struct {
	router *gin.Engine
	port int
	name string
}

func NewKraikubServer(conf config.Config) kraikubServer {

	r := gin.Default()
	return kraikubServer{
		router: r,
		port: conf.Server.Port,
		name: conf.Server.Name,
	}
}

func (serv *kraikubServer) Router() *gin.Engine {
	return serv.router
}

func (serv *kraikubServer) Start() {

	srv := &http.Server{
		Handler: serv.router,
		Addr:    fmt.Sprintf(":%d", serv.port),
	}

	go func() {
		// enable service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("Server was successfully shutdown")
	}
	log.Println("Server exiting")
}
