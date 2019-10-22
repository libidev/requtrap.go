package main

import (
  "log"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type Book struct{
  Name string `json:"name"`
  Pages int  `json:"pages"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request){
  var response []byte
  if r.Method == http.MethodGet {
    // GET
    js, err := json.Marshal(books)
    if err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.Fatal(err)
      return
    }
    response = js
  } else if r.Method == http.MethodPost {
    // POST
    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()
    if err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.Fatal(err)
      return
    }
    var book Book
    err = json.Unmarshal(body, &book)
    if err != nil{
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.Fatal(err)
      return
    }

    books = append(books, book)
    response = []byte("OK")
  }

  // send response
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}


func main(){
  http.HandleFunc("/books", getBooks)
  http.ListenAndServe(":8001", nil)
}
