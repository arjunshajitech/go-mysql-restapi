package main

import "go/packages/connection"
import "go/packages/routes"

func main() {
	connection.InitlizeMigration()
	routes.InitlizeRoute()
}
