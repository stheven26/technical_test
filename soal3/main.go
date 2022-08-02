package main

import (
	"github.com/stheven26/technical_test/db"
	"github.com/stheven26/technical_test/routes"
)

func main() {
	db.SetupDB()
	routes.StartApplication()
}
