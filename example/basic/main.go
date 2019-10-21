package main

import (
  "log"
  "net/http"
  "encoding/json"
)

type Books struct{
  Name string
  Pages int
}

func getBooks(w http.ResponseWriter,r *http.Request){
  books := Books{"Hello",18}

  js, err := json.Marshal(books)
  if err != nil{
    http.Error(w,err.Error(),http.StatusInternalServerError)
    log.Fatal(err)
    return
  }

  w.Header().Set("Content-Type","application/json")
  w.Write(js)
}


func main(){
  http.HandleFunc("/books",getBooks)
  http.ListenAndServe(":8001",nil)
}
