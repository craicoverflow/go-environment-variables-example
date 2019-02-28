package main

import (
	"fmt"
	"log"

	"github.com/craicoverflow/go-environment-variables-example/config"
	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.New()

	// Print out environment variables
	fmt.Println(conf.GitHub.Username)
	fmt.Println(conf.GitHub.APIKey)
	fmt.Println(conf.DebugMode)
	fmt.Println(conf.MaxUsers)

	// Print out each role
	for _, role := range conf.UserRoles {
		fmt.Println(role)
	}
}
