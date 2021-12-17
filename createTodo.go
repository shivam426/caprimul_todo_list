package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

// // func createTodo(w http.ResponseWriter, r *http.Request) {
// // 	w.Header().Set("Content-Type", "application/json")
// // 	var todo Todo
// // 	err := json.NewDecoder(r.Body).Decode(&todo)

// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	sqlStatement := `INSERT INTO "public.todo" (title,todo_description,status) VALUES ($1,$2,$3)`
// // 	_, err = conn.Exec(sqlStatement, todo.Title, todo.Todo_description, todo.Status)
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	fmt.Println(todo)

// // 	w.WriteHeader(http.StatusOK)
// // 	json.NewEncoder(w).Encode(todo)
// }
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo

	_ = json.NewDecoder(r.Body).Decode(&todo)
	result, err := conn.InsertOne(context.TODO(), todo)

	if err != nil {
		log.Fatal(err)
		// return
	}

	json.NewEncoder(w).Encode(result)
}
