package main

import (
	"ScrumGameBot/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.PutTask)
	http.HandleFunc("/tasks", handlers.GetTask)
	http.HandleFunc("/_ah/health", handlers.HealthCheckHandler)
	http.HandleFunc("/hook", handlers.Hook)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
