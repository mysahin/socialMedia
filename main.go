package main

import (
	"socialMedia/Database"
	"socialMedia/Router"
)

func main() {

	Database.Connect()
	app := Router.Routes()

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
