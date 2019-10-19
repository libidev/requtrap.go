package http


import (
  "fmt"
  "strconv"
  "net/http"
  "github.com/libidev/requtrap.go/cli/config"
)

func StartingServe(conf *config.ConfigYaml){
  var port = ":" + strconv.Itoa(conf.Port)

  fmt.Printf("running server in %s%s\n",conf.Host,port)
  http.ListenAndServe(port,nil)
}


func CheckRequestMethod(w http.ResponseWriter, r *http.Request){
  
  if r.Method == http.MethodGet{
    fmt.Println("====> get")
    //TODO
  }else if r.Method == http.MethodPost{
    fmt.Println("====> post")
    //TODO
  }else if r.Method == http.MethodPut{
    fmt.Println("====> put")
    //TODO
  }else if r.Method == http.MethodDelete{
    fmt.Println("====> delete")
    //TODO
  }else{
    fmt.Println("====> method undefined")
  }
}
