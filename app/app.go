package app

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/killtheverse/nitd-results/app/db"
	"github.com/killtheverse/nitd-results/config"
)

type App struct {
	Router		*mux.Router
	DBClient	*mongo.Client
	logger		*log.Logger
}

func ConfigAndRun(config *config.Config) {
	app := new(App)
	app.Initialize(config)
}

func(app *App) Initialize(config *config.Config) {
	app.Router = mux.NewRouter()
	app.logger = log.New(os.Stdout, "results-app ", log.LstdFlags)
	app.DBClient = db.Connect(config.DBName, config.DBURI, app.logger)
}