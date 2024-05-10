package main

import (
	"main/router"
)

func main() {

	r := router.SetupRouter()

	r.Run()
}
