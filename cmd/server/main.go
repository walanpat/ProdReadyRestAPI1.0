package main

import "fmt"

// App - struct which contains things like pointers
// to database connections
type App struct{}

// Run sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")
	return nil
}
func main() {
	fmt.Println("Go Rest Api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error at startup of REST API")
	}
}
