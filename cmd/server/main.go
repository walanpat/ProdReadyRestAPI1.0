package main

import (
	transportHTTP "ProdReadyRestAPI1.0/internal/transport/http"
	"fmt"
	"net/http"
)

// App - struct which contains things like pointers
// to database connections
type App struct{}

// Run sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")
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