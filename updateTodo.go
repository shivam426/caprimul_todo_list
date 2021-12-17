package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func updateTodo(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	todo := Todo{}
// 	param := r.URL.Query().Get("product_id")
// 	err := json.NewDecoder(r.Body).Decode(&todo)
// 	sqlStatement := `UPDATE public.todo SET title=$2,todo_description= $3 WHERE todo_id = $1;`
// 	_, err = conn.Exec(sqlStatement, param, todo.Title, todo.Todo_description)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("successfully updated")
// 	json.NewEncoder(w).Encode(todo)

// }
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	todo := Todo{}

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&todo)

	update := bson.D{
		{"$set", bson.D{
			{"title", todo.Title},
		}},
	}

	err := conn.FindOneAndUpdate(context.TODO(), filter, update).Decode(&todo)

	if err != nil {
		log.Fatal(err)
		// return
	}

	todo.Id = id

	json.NewEncoder(w).Encode(todo)
}
