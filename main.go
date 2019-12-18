package main

import (
  "log"
  "os"
  "github.com/libidev/requtrap.go/cli"
  "github.com/libidev/requtrap.go/cli/errors"
)

//func IsError(err error){
//  if err != nil{
//    log.Fatal("error : %v",err)
//  }
//}

func main() {
  var err error
  defer errors.IsError(err)

  cli.Parse(os.Args[1:])

  if err != nil {return} 
}
