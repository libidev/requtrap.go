package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Book struct {
	Name  string `json:"name"`
	Pages int    `json:"pages"`
}

var books []Book = []Book{}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var response []byte
	if r.Method == http.MethodGet {
		// GET
		js, err := json.Marshal(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response = js
	} else if r.Method == http.MethodPost {
		// POST
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		var book Book
		err = json.Unmarshal(body, &book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		books = append(books, book)
		response = []byte("OK")
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	port := ":8001"
	http.HandleFunc("/books", getBooks)
	fmt.Printf("book service running on http://localhost:%s", port)
	http.ListenAndServe(port, nil)
}
