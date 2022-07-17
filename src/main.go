package main

func main() {
	const PORT int = 8000

	// create database
	var database = New(PORT)

	// start server
	HandleConnections(&database)
}
