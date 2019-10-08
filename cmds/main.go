package main

import (
  "log"
  "fmt"
  "os"

  "gopkg.in/yaml.v3"
)

var(
  Default = `
  #server
  host: "8080"

  #gateway example 1
  request:
    url: /test
    addr: http://127.0.0.1:8000
  

  #gateway example 2
  request:
    url: /coba
    addr: http://127.0.0.1:8080


  `
)


func isError(err error){
  if err != nil {
    log.Fatal("error: %v",err)
  }
}

func main() {
  var err error
  defer isError(err)

  t := yaml.Node{}
  err = yaml.Unmarshal([]byte(Default),&t)
  if err != nil {return}

  b,err:= yaml.Marshal(&t) 
  if err != nil {return}

  f,err := os.Create("requtrap_example.yaml")
  if err != nil {return}
  fmt.Fprintf(f,string(b))
  err = f.Close() 
  if err != nil {return} 
}
