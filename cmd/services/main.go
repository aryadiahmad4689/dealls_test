package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aryadiahmad4689/dealls_test/src/app"
	"github.com/aryadiahmad4689/dealls_test/src/modules"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var (
		app    = app.Init()
		router = app.GetHttpRouter()
		module = modules.Init(app)
	)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://", "http://"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	router.Mount("/v1/users", module.User.GetHttpRouter())
	router.Mount("/v1/swipes", module.Swipe.GetHttpRouter())
	router.Mount("/v1/packages", module.Package.GetHttpRouter())
	router.Mount("/v1/subscriptions", module.Subscription.GetHttpRouter())

	err = app.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
