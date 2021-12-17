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

// func deleteTodo(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	sqlStatement := `DELETE FROM public.todo WHERE todo_id = $1;`
// 	param := r.URL.Query().Get("todo_id")
// 	result, err := conn.Exec(sqlStatement, param)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("successfully deleted")
// 	json.NewEncoder(w).Encode(result)
// }
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := conn.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err, w)
		// return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
