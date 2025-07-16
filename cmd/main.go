// Swagger
//
//  @title                       Subscribers API
//  @version                     1.0
//  @description                 A comprehensive API for managing pets, offering endpoints for creation, update, deletion, and retrieval of pet data.
//  @termsOfService              http://petmanagement.com/terms
//  @contact.name                API Support Team
//  @contact.url                 http://petmanagement.com/support
//  @contact.email               support@petmanagement.com
//  @license.name                Apache 2.0
//  @license.url                 http://www.apache.org/licenses/LICENSE-2.0.html
//  @host                        petmanagement.com
//  @BasePath                    /api/v1
//  @schemes                     http https
//  @securityDefinitions.apiKey  JWT
//  @in                          header
//  @name                        Authorization
//  @description                 JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	_ "subscribers/docs"
	database "subscribers/internal"
	_ "subscribers/model"

	"syscall"
	"time"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	server := database.InitServer()

	fmt.Println("server is running: ", server.Addr)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	fmt.Println("Server is running at ", server.Addr)

	err := server.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
}
