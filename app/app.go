package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/killtheverse/nitd-results/app/db"
	"github.com/killtheverse/nitd-results/config"
)

// Structure for the app
// It contains the serverAddress, router, Database client and the logger
type App struct {
	serverAddress	string
	Router			*mux.Router
	DBClient		*mongo.Client
	logger			*log.Logger
}

// Configure the app and run
func ConfigAndRun(config *config.Config) {
	app := new(App)
	app.initialize(config)
	app.run()
}

// Initialize the app
func(app *App) initialize(config *config.Config) {
	app.serverAddress = config.ServerAddress
	app.Router = mux.NewRouter()
	app.logger = log.New(os.Stdout, "results-app ", log.LstdFlags)
	app.DBClient = db.Connect(config.DBName, config.DBURI, app.logger)
}

// Register the routes in the router
func(app *App) setupRouters() {
	// TODO:Implement
}

// Run will start the http server
func(app *App) run() {
	server := http.Server{
		Addr: app.serverAddress,		// configure the bind address
		Handler: app.Router,			// set the default handler
		ErrorLog: app.logger,			// set the logger for the server
		ReadTimeout: 5*time.Second,		// max time to read request from the client
		WriteTimeout: 10*time.Second,	// max time to write response to the client
		IdleTimeout: 120*time.Second,	// max time for conncections using TCP Keep-Alive
	}

	go func() {
		app.logger.Println("Starting the server on:", app.serverAddress)
		err := server.ListenAndServe()
		if err != nil {
			app.logger.Fatal("[ERROR] Can't start server:", err)
		}
	}()

	// Signals for shutting down the server
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	
	// Block until a signal is recieved
	sig := <-sigs
	app.logger.Printf("Trapped signal:%v\nShutting down the server", sig)

	// Disconnect the MongoDB client
	db.Disconnect(app.DBClient, app.logger)

	// Shutdown the server, waiting for max 30 seconds
	app.logger.Print("Gracefully stopping server")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}