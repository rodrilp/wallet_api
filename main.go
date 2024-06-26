package main

import (
	"fmt"
	"main/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	
	if err != nil {
	  fmt.Println("Error loading .env file.", err)
	}

	r := router.SetupRouter()

	err = r.Run()

	if err != nil {
		fmt.Println("Error running the Router")
	}
}
