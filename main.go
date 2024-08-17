package main

import (
	"aristio-sagala-test/config"
	"aristio-sagala-test/routes"
)

func main() {
	config.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":3000")
}
