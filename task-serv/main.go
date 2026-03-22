package main

import (
	"fmt"
	"net/http"
)

func tasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "tasks list")
}

func main() {
	http.HandleFunc("/tasks", tasks)

	fmt.Println("Task service running on :8081")
	http.ListenAndServe(":8081", nil)
}