package main

import "fmt"

func main() {
	const PORT int = 8000

	// create database
	var database = Database{}.New(PORT)

	database.Set("test", "test")
	database.Delete("test")

	fmt.Println(database.Get("test"))
}
