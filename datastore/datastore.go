package datastore

import (
	"ScrumGameBot/domain"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

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
