package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

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

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			if os.IsNotExist(err) {
				http.Error(w, "File not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error reading file", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write(data)
	})

	fmt.Println("Backend server v2.0 with CORS running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
