package main

import (
  "log"
  "fmt"
  "os"

  "gopkg.in/yaml.v3"
)

var(
  Default = `
    name: book-store
    host: 127.0.0.1
    port: 8080
    services:
      - path: /books
        upstream: http://127.0.0.1:8001
      - path: /authors
        upstream: http://127.0.0.1:8002
  `
)

func isError(err error){
  if err != nil {
    log.Fatal("error: %v",err)
  }
}

type ConfigYaml struct{
  Host string `yaml:"host"`
  
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
