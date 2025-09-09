package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод", http.StatusMethodNotAllowed)
		return
	}

	var RequestBody requestBody
	err := json.NewDecoder(r.Body).Decode(&RequestBody)
	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}
	task = RequestBody.Task

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Запрос изменен"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод", http.StatusMethodNotAllowed)
		return
	}
	getTask := task
	response := fmt.Sprintf("Hello, %s", getTask)
	w.Write([]byte(response))

}

func main() {
	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getHandler(w, r)
		case http.MethodPost:
			postHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
