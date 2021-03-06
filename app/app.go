package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	gohandlers "github.com/gorilla/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/killtheverse/nitd-results/app/db"
	"github.com/killtheverse/nitd-results/app/handlers"
	logger "github.com/killtheverse/nitd-results/app/logging"
	"github.com/killtheverse/nitd-results/app/utils"
	"github.com/killtheverse/nitd-results/config"
)

// Structure for the app
// It contains the serverAddress, router, Database client, Database name and the logger
type App struct {
	serverAddress	string
	Router			*mux.Router
	DBClient		*mongo.Client
	DB				*mongo.Database
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
	app.DBClient = db.Connect(config.DBURI)
	app.DB = app.DBClient.Database(config.DBName)
	app.setupMiddlewares()
	app.createIndexes()
	app.setupRouters()
}

// setupMiddlewares will add middlewares to main router
func(app *App) setupMiddlewares() {
	app.Router.Use(utils.JSONContentTypeMiddleware)
	app.Router.Use(utils.LoggingMiddleware)
}

// createIndexes will create unique and index fields
func (app *App) createIndexes() {
	keys := bsonx.Doc{
		{Key: "roll_no", Value: bsonx.Int32(1)},
	}
	students := app.DB.Collection("students")
	db.SetIndexes(students, keys)
}

// Register the routes in the router
func(app *App) setupRouters() {
	doc_opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	doc_middleware := middleware.Redoc(doc_opts, nil)
	app.Router.Handle("/docs", doc_middleware)
	app.Router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	app.Router.Handle("/", http.HandlerFunc(func (rw http.ResponseWriter, r *http.Request) {
		logger.Write("in redirect")
		http.Redirect(rw, r, "https://"+r.Host+r.URL.String()+"/docs", http.StatusMovedPermanently)
	}))

	authRouter := app.Router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signin/", handlers.SignIn).Methods("POST")

	studentRouter := app.Router.PathPrefix("/students").Subrouter()
	studentRouter.HandleFunc("", app.handleRequest(handlers.GetStudents)).Methods("GET")
	studentRouter.HandleFunc("/{roll_number}", app.handleRequest(handlers.GetStudent)).Methods("GET")

	studentUpdateRouter := studentRouter.Methods(http.MethodPut).Subrouter()
	studentUpdateRouter.Use(utils.AuthenticationRequiredMiddleware)
	studentUpdateRouter.HandleFunc("/{roll_number}", app.handleRequest(handlers.UpdateStudent)).Methods("PUT")
}

// Run will start the http server
func(app *App) run() {

	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string {"*"}))

	server := http.Server{
		Addr: app.serverAddress,		// configure the bind address
		Handler: corsHandler(app.Router),			// set the default handler
		ReadTimeout: 5*time.Second,		// max time to read request from the client
		WriteTimeout: 10*time.Second,	// max time to write response to the client
		IdleTimeout: 120*time.Second,	// max time for conncections using TCP Keep-Alive
	}

	go func() {
		logger.Write("Starting the server on: %v", app.serverAddress)
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal("[ERROR] Can't start server: %v", err)
		}
	}()

	// Signals for shutting down the server
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	
	// Block until a signal is recieved
	sig := <-sigs
	logger.Write("Trapped signal:%v\nShutting down the server", sig)

	// Disconnect the MongoDB client
	db.Disconnect(app.DBClient)

	// Shutdown the server, waiting for max 30 seconds
	logger.Write("Gracefully stopping server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}

// RequestHandlerFunction is a custome type that help us to pass db arg to all endpoints
type RequestHandlerFunction func(db *mongo.Database, w http.ResponseWriter, r *http.Request)

// handleRequest is a middleware we create for pass in db connection to endpoints.
func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}