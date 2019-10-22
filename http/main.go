package http

import (
	"fmt"
  "log"
  "strconv"
  "net/http"
  "github.com/libidev/requtrap.go/cli/config"
)

func Serve(conf *config.ConfigYaml){
  var host = conf.Host
  var port = strconv.Itoa(conf.Port)
  var uri  = host + ":" + port

  fmt.Printf("%s running on http://%s\n", conf.Name, uri)

  handler := &HttpHandler{}
  for _, service := range conf.Services {
    handler.AddRoute(service)
  }

  err := http.ListenAndServe(uri, handler)
  if err != nil{
    log.Fatal(err)
  }
}
