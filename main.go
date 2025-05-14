package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"tesodev-product-api/config"
	"tesodev-product-api/routes"
)

func main() {
	// Logger setup
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// Initialize Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Database
	config.ConnectDatabase()

	// Setup Routes
	routes.SetupRoutes(e)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
