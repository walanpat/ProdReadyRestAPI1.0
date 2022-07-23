package main

import (
	"ProdReadyRestAPI1.0/internal/database"
	transportHTTP "ProdReadyRestAPI1.0/internal/transport/http"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

// App - struct which contains things like pointers
// to database connections
type App struct{}

// Run sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		fmt.Println("Error at NewDatabase")
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}
func main() {
	fmt.Println("Go Rest Api initilization start")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error at startup of REST API")
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	}
}
