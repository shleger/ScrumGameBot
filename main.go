package main

import (
	"ScrumGameBot/datastore"
	"ScrumGameBot/handlers"
	"log"
	"net/http"
)

func main() {

	app := &handlers.App{DbSrv: datastore.DB}

	http.HandleFunc("/taskput", makeHandler((app.PutTask)))
	http.HandleFunc("/taskget", makeHandler((app.GetTask)))
	http.HandleFunc("/_ah/health", handlers.HealthCheckHandler)
	http.HandleFunc("/hook", handlers.Hook)
	http.HandleFunc("/echo", makeHandler((app.Echo)))

	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}
