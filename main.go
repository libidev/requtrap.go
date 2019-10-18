package main

import (
  "log"
  "os"
  "github.com/libidev/requtrap.go/cli"
)

func isError(err error){
  if err != nil {
    log.Fatal("error: %v",err)
  }
}

func main() {
  var err error
  defer isError(err)

  cli.Parse(os.Args[1:])

  if err != nil {return} 
}
