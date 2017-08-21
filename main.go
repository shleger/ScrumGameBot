package main

import (
	"ScrumGameBot/domain"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

func main() {

	http.HandleFunc("/", put)
	http.HandleFunc("/tasks", get)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	http.HandleFunc("/hook", hook)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

var (
	kind = "Task"
	err  error
	DB   PropsService
)

func hook(w http.ResponseWriter, r *http.Request) {
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

func init() {
	DB, err = newDatastoreDB()
	log.Print("Init Datastore  Done")

}

type datastoreDB struct {
	client *datastore.Client
}

type PropsService interface {
	GetKey(key string) string
	PutKey(key string, task *domain.Task)
}

func (db *datastoreDB) PutKey(key string, task *domain.Task) {
	ctx := context.Background()
	q := datastore.NameKey(kind, key, nil)
	//eeee
	// Saves the new entity.
	if _, err := db.client.Put(ctx, q, task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

}
func (db *datastoreDB) GetKey(key string) string {

	ctx := context.Background()
	q := datastore.NameKey(kind, key, nil)

	task := domain.Task{}
	if err := db.client.Get(ctx, q, &task); err != nil {
		log.Fatalf("Failed  to get  task with datastoreDB: %v", err)

	}

	return task.Description

}

type Post struct {
	Title       string
	Body        string `datastore:",noindex"`
	PublishedAt time.Time
}

func newDatastoreDB() (PropsService, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, gcid())

	if err != nil {
		return nil, fmt.Errorf("Failed to create client: %v", err)

	}

	return &datastoreDB{
		client: client,
	}, nil
}

//get gcloud project id
func gcid() string {
	return os.Getenv("GCLOUD_PROJECT")
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tasks" {
		http.NotFound(w, r)
		return
	}

	t := DB.GetKey("sampletask3")
	fmt.Fprint(w, t)

}

func put(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	task := domain.Task{
		Description: "Buy milk3",
	}

	DB.PutKey("sampletask3", &task)

	fmt.Fprint(w, "hello from db")

}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
