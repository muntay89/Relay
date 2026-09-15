package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Topic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

type TopicStore struct {
	topics map[string]Topic
}

func (topic Topic) Validate() error {
	if topic.Name == "" {
		return errors.New("topic must have a name")
	}

	return nil
}

func (topic *Topic) Disable() {
	topic.Enabled = false
}

func topicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var topic Topic
	err := json.NewDecoder(r.Body).Decode(&topic)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err = topic.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(topic)
	if err != nil {
		fmt.Println("Failed to encode response:", err)
	}
}

func main() {
	http.HandleFunc("/topics", topicHandler)
	fmt.Println("Relay running on http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
