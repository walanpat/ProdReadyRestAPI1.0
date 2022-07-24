package main

import (
	"ProdReadyRestAPI1.0/internal/comment"
	"ProdReadyRestAPI1.0/internal/database"
	transportHTTP "ProdReadyRestAPI1.0/internal/transport/http"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// App - contain application information
type App struct {
	Name    string
	Version string
}

// Run sets up our application
func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up Application")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Error at NewDatabase")
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}
	return nil
}
func main() {
	app := App{
		Name:    "Commenting Service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Error at startup of REST API")
		log.Fatal(err)
	}
}

//init - Reads our .env file, which has our database credentials
func init() {
	err := godotenv.Load()
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
}
