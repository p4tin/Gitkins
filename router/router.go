package router

import (
	"log"
	"net/http"

	"github.com/p4tin/Gitkins/handlers"
)

func Server() {
	http.HandleFunc("/event", handlers.GitEventHandler)

	//ngrok http -subdomain="p4tin" 8081
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
