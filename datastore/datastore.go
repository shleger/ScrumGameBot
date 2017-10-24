package datastore

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"ScrumGameBot/domain"

	"cloud.google.com/go/datastore"
)

var (
	kind = "Task"
	err  error
	DB   PropsService
)

func init() {
	println("Init newDatastoreDB")

	//TODO rm init db
	DB, err = newDatastoreDB()
	log.Printf("Init Datastore done for GCLOUD_PROJECT=[%s]", gcid())

}

type datastoreDB struct {
	client *datastore.Client
}

type PropsService interface {
	GetKey(key string) string
	PutKey(key string, task *domain.Task)
	GetToken(key string, ns string) string
	EchoTask(key string) string
}

func (db *datastoreDB) EchoTask(key string) string {
	return "RealStub+" + key
}

func (db *datastoreDB) PutKey(key string, task *domain.Task) {
	ctx := context.Background()
	q := datastore.NameKey(kind, key, nil)
	// Save new entity.
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

func (db *datastoreDB) GetToken(key string, ns string) string {
	ctx := context.Background()
	q := datastore.NewQuery("Props").Namespace(ns).Filter("Key =", key)

	var prop []domain.Props

	if _, err := db.client.GetAll(ctx, q, &prop); err != nil {
		log.Fatalf("Failed  to get  prop [%s]  with datastoreDB: %v", key, err)

	}

	//TODO impl Error()
	return prop[0].Val

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
