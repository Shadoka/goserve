package main

import (
	"connection"
)

func main() {
	server := connection.Server{Connected: false}
	server.Start(":8081")
}