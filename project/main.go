package main

import (
	"flag"
	"fmt"
	"goProject/internal/app"
	"goProject/internal/routes"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Go backend port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		// panic will crash the app
		panic(err)
	}

	
	
	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Printf("We are running on port %d\n", port)
	
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
	defer app.DB.Close()

}

