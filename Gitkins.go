package main

import (
	"fmt"
	"github.com/p4tin/Gitkins/config"
	"github.com/p4tin/Gitkins/router"
	"log"

	"encoding/json"
)

func main() {
	log.Printf("Welcome to Gitkins %s\n", config.Version)

	if config.Config.Debug {
		res2B, _ := json.MarshalIndent(config.Config, "", "    ")
		fmt.Println(string(res2B))
	}

	router.Server()
}
