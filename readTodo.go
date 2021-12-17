package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// func readTodo(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	rows, err := conn.Query("SELECT * FROM public.todo LIMIT 1000")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	todos := []Todo{}

// 	for rows.Next() {
// 		todo := Todo{}
// 		err = rows.Scan(&todo.Todo_id, &todo.Title, &todo.Todo_description, &todo.Status)
// 		if err != nil {
// 			panic(err)
// 		}
// 		todos = append(todos, todo)
// 		fmt.Println(todo)
// 	}

// 	json.NewEncoder(w).Encode(todos)

// }
func readTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	todos := []Todo{}
	cur, err := conn.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
		// return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var todo Todo
		err := cur.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	json.NewEncoder(w).Encode(todos)

}

// func getProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product Products

// 	var params = mux.Vars(r)

// 	id, _ := primitive.ObjectIDFromHex(params["id"])

// 	filter := bson.M{"_id": id}
// 	err := collection.FindOne(context.TODO(), filter).Decode(&product)

// 	if err != nil {
// 		log.Fatal(err)
// 		// return
// 	}

// 	json.NewEncoder(w).Encode(product)
// }
