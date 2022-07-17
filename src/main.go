package main

import "fmt"

func main() {
	const PORT int = 8000

	// create database
	var database = New(PORT)

	database.Set("test", "test")

	fmt.Println(database.Size())
}
