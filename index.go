package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conn = connectdb()

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todo", createTodo).Methods("POST")
	r.HandleFunc("/todo", readTodo).Methods("GET")
	r.HandleFunc("/todo/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))

}

func connectdb() *mongo.Collection {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("todo_list").Collection("todo")
	return collection
}
