package main

import (
	"log"
	"os"
	"path"

	"github.com/degreane/ezekod.com/middleware/ezelogger"
	"github.com/degreane/ezekod.com/model"
	"github.com/degreane/ezekod.com/server"
	"github.com/gofiber/fiber/v2"
)

var (
	// Database
	DB model.DB

	// Server Fiber Application
	Server fiber.App
)

// Init function that runs for initialization options within the system
func init() {
	// Log: -> Initialize
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	logFile := path.Join(path.Clean(cwd), "logs", "ezekod.log")
	ezelogger.SetLogger(logFile)
	// Log: -> End Of Initialization

	// DB: -> Init Database MongoDB connection
	err = DB.Connect()
	if err != nil {
		ezelogger.Ezelogger.Fatalf("%+v", err)
	}
	// DB: -> End of Init Database Connection

	// Fibers: -> Initialize Application
	server.App = server.Init()
	server.App.SetRoutes()
	server.App.StartServer()
	// Fibers: -> End of Initialization
}
