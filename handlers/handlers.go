package handlers

import (
	"ScrumGameBot/datastore"
	"ScrumGameBot/domain"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hook - telegram WebHook
func Hook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path == "/hook" && r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		//		messg := EchoResponce{}
		//		json.Unmarshall(body, &messg)
		if err != nil {
			fmt.Fprintf(w, "Error %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		log.Println("HOOK_JSON:" + string(body))

	}
}

//GetTask -- get task by const
func GetTask(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tasks" {
		http.NotFound(w, r)
		return
	}

	t := datastore.DB.GetKey("sampletask3")
	fmt.Fprint(w, t)

}

//PutTask -save task to db
func PutTask(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	task := domain.Task{
		Description: "Buy milk3",
	}

	datastore.DB.PutKey("sampletask3", &task)

	fmt.Fprint(w, "hello from db")

}

//HealthCheckHandler - check
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
