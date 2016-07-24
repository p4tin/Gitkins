package router

import (
	"log"
	"net/http"

	"github.com/p4tin/Gitkins/handlers"
	"github.com/p4tin/Gitkins/config"
)

func Server() {
	http.HandleFunc("/health", handlers.HealthEventHandler)
	http.HandleFunc("/event", handlers.GitEventHandler)

	//ngrok http -subdomain="urbn-ci" 8081
	log.Println("Listening at: http://0.0.0.0:" + config.Config.Port)
	log.Println("Health Check: http://0.0.0.0:" + config.Config.Port + "/health")
	log.Fatal(http.ListenAndServe("0.0.0.0:" + config.Config.Port, nil))
}
