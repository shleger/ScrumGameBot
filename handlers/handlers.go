package handlers

import (
	"ScrumGameBot/datastore"
	"ScrumGameBot/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type App struct {
	db datastore.PropsService
}

//Echo is telegram echo handler
func (a *App) Echo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" {
		http.NotFound(w, r)
		return
	}

	resp := a.db.EchoTask("echoKey")
	fmt.Fprint(w, resp)
}

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

		mapping := domain.EchoResponce{}

		if err := json.Unmarshal(body, &mapping); err != nil {
			log.Printf("Unmarshall %v ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		log.Println(mapping.ID)
		log.Println(mapping.Message.From.FirstName)

	}
}

//GetTask -- get task by const
func GetTask(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tasks" {
		http.NotFound(w, r)
		return
	}

	t := datastore.DB.GetToken("TEST_TOKEN", "auth")
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
