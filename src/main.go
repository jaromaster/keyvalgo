package main

import (
	"os"
)

func main() {
	const PORT int = 8000
	const PASSWD_ENV_NAME string = "KEYVALGO_PW"

	// get password from env
	var password string = os.Getenv(PASSWD_ENV_NAME)
	if password == "" {
		panic("no password set, please use $" + PASSWD_ENV_NAME)
	}

	// create database
	var database = New(PORT, password)

	// start database server
	database.HandleConnections()
}
