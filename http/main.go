package http

import (
  "fmt"
  "log"
  "io/ioutil"
  "time"
  "strconv"
  // "encoding/json"
  "net/http"
  "github.com/libidev/requtrap.go/cli/config"
)

type HttpHandler struct {
  Routes []config.ConfigService
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  //fmt.Println(h.GetRequestMethod(r))

  if r.URL.Path != "favicon.ico"{
    for _, service := range h.Routes{
      if r.URL.Path == service.Path{
        //GATEWAY
        tr := &http.Transport{
          MaxIdleConns:       10,
          IdleConnTimeout:    30 * time.Second,
          DisableCompression: true,
        }

        url := service.Upstream + service.Path
        client := &http.Client{Transport: tr}
        resp, err := client.Get(url)
        if err!=nil{
          log.Fatal(err)
        }else{
          defer resp.Body.Close()

          if resp.StatusCode == http.StatusOK{
            contents, err := ioutil.ReadAll(resp.Body)
            if err != nil{
              log.Fatal(err)
            }

            w.Header().Set("Content-Type", "application/json")
            w.Write([]byte(contents))
          }
        }
      }
    }
  }
}

func (h HttpHandler) GetRequestMethod(r *http.Request) string {
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
