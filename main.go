package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusBadRequest)
		return
	}
	var msg Message
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
	}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content/Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the Avneet Pandey's Website")

}
func main() {
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/post", postHandler)
	pNumber := ":8091"
	http.ListenAndServe(pNumber, nil)

}
