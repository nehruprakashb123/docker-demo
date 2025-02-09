package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello from Golang backend!"}`))
}

func main() {
	http.HandleFunc("/", handler)
	port := 8080
	fmt.Printf("Server running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
