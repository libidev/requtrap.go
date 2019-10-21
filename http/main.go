package http

import (
  "fmt"
  "log"
  "strconv"
  "net/http"
  "github.com/libidev/requtrap.go/cli/config"
)

type HttpHandler struct {
  Routes []config.ConfigService
}


func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  //fmt.Println(h.GetRequestMethod(r))

  if r.URL.Path != "favicon.ico"{
    for _, service := range h.Routes{
      if r.URL.Path == service.Path{
        //GATEWAY
        fmt.Printf("redirect to : %v\n",service.Upstream)
      }
    }
  }
}

func (h *HttpHandler) GetRequestMethod(r *http.Request) string {
  return r.Method
}

func (h *HttpHandler) AddRoute(service config.ConfigService) {
  h.Routes = append(h.Routes, service)
}

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
