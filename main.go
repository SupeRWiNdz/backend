package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		err = ioutil.WriteFile("data.txt", body, 0644)
		if err != nil {
			http.Error(w, "Error saving data", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Data saved successfully!")
	})

	fmt.Println("Backend server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
