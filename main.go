package main

import (
	"context"
	"final-project/connection"

	"final-project/service/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	HttpServer *http.Server
}

func main() {

	gormDB, err := connection.NewConnection()
	if err != nil {
		os.Exit(1)
	}

	// Init Router
	router := StartServer(gormDB)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Start the server in a separate goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// Graceful shutdown logic
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish the current request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}

	fmt.Println("Server exiting")
}

func StartServer(gormDB *connection.GormDB) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	routes.InitMygramRoute(api, gormDB).Routes()
	return router
}
