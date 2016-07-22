package main

import (
        "log"
	"github.com/p4tin/Gitkins/router"
)

var version = "0.0.1.Alpha-1"


func main() {
	log.Printf("Welcome to Gitkins %s\n", version)

	router.Server()
}
