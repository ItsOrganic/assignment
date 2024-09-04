package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var numbers []int32
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
	}
	var res int32

	for _, num := range numbers {
		res += num
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, res)
}
