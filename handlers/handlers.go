package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shleger/ScrumGameBot/datastore"
	"github.com/shleger/ScrumGameBot/domain"
)

type App struct {
	DbSrv datastore.PropsService
}

//Echo is telegram echo handler
func (a *App) Echo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" {
		http.NotFound(w, r)
		return
	}

	resp := a.DbSrv.EchoTask("echoKey")
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
func (a *App) GetTask(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/taskget" {
		http.NotFound(w, r)
		return
	}

	t := a.DbSrv.GetToken("TEST_TOKEN", "auth")
	fmt.Fprint(w, t)

}

//PutTask -save task to db
func (a *App) PutTask(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/taskput" {
		http.NotFound(w, r)
		return
	}

	task := domain.Task{
		Description: "Buy milk3",
	}

	a.DbSrv.PutKey("sampletask3", &task)

	fmt.Fprint(w, "hello from db")

}

//HealthCheckHandler - check
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
