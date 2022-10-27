package main

import (
	"golang_web_programming/server"
)

func main() {
	server := server.NewDefaultServer()
	server.Run()
}
