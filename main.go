package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

func main() {

	http.HandleFunc("/", handle)
	http.HandleFunc("/tasks", taskHandle)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

var (
	kind = "Task"
	err  error
	DB   PropsService
)

func init() {
	DB, err = newDatastoreDB()
	log.Print("Init Datastore  Done")

}

type datastoreDB struct {
	client *datastore.Client
}

type PropsService interface {
	GetKey(key string) string
}

func (db *datastoreDB) GetKey(key string) string {

	ctx := context.Background()
	q := datastore.NameKey(kind, key, nil)

	task := Task{}
	if err := db.client.Get(ctx, q, &task); err != nil {
		log.Fatalf("Failed  to get  task with datastoreDB: %v", err)

	}

	return task.Description

}

type Task struct {
	Description string
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

func taskHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tasks" {
		http.NotFound(w, r)
		return
	}

	t := DB.GetKey("sampletask2")
	fmt.Fprint(w, t)

}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := context.Background()

	for _, e := range os.Environ() {
		log.Println(e)
	}

	// Creates a client.
	client, err := datastore.NewClient(ctx, gcid())

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the kind for the new entity.
	kind := "Task"
	// Sets the name/ID for the new entity.
	name := "sampletask2"
	// Creates a Key instance.
	taskKey := datastore.NameKey(kind, name, nil)

	// Creates a Task instance.
	task := Task{
		Description: "Buy milk2",
	}

	// Saves the new entity.
	if _, err := client.Put(ctx, taskKey, &task); err != nil {
		log.Fatalf("Failed to save task: %v", err)
	}

	fmt.Fprint(w, "hello from db")

}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
