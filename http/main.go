package http

import (
  "fmt"
  "log"
  "bytes"
  "io/ioutil"
  "time"
  "strconv"
  "encoding/json"
  "net/http"
  "github.com/libidev/requtrap.go/cli/config"
  "github.com/libidev/requtrap.go/cli/errors"
)

type HttpHandler struct {
  Routes []config.ConfigService
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  var err error
  defer errors.IsError(err)

  if r.URL.Path != "favicon.ico"{
    for _, service := range h.Routes{
      if r.URL.Path == service.Path{
        tr := &http.Transport{
          MaxIdleConns:       10,
          IdleConnTimeout:    30 * time.Second,
          DisableCompression: true,
        }

        url := service.Upstream + service.Path
        client := &http.Client{Transport: tr}
        if h.GetRequestMethod(r) == "GET"{


          resp, err := client.Get(url)
          if err!=nil{return}

          defer resp.Body.Close()

          if resp.StatusCode == http.StatusOK{
            contents, err := ioutil.ReadAll(resp.Body)
            if err != nil{
              http.Error(w,err.Error(),http.StatusInternalServerError)
              return
            }

            var result []map[string]interface{}

            if err := json.Unmarshal(contents,&result); err != nil{
              http.Error(w,err.Error(),http.StatusInternalServerError)
              return
            }

            js, err := json.Marshal(result)

            fmt.Printf("\nredirect to : %s\n",service.Upstream)
            fmt.Println("response :")
            fmt.Println(string(js))

            w.Header().Set("Content-Type", "application/json")
            w.Write(js)
          }
        }else if h.GetRequestMethod(r) == "POST"{

          var jsonStr ,err = ioutil.ReadAll(r.Body)
          if err != nil {
            http.Error(w,err.Error(),http.StatusInternalServerError)
            return
          }

          req, _ := http.NewRequest("POST",url,bytes.NewBuffer([]byte(jsonStr)))
          req.Header.Set("X-Custom-Header","myvalue")
          req.Header.Set("Content-type","application/json")

          resp, err := client.Do(req)
          if err != nil{
            http.Error(w,err.Error(),http.StatusInternalServerError)
            return
          }
          defer resp.Body.Close()

          fmt.Println("response status : ",resp.Status)
          fmt.Println("response Header : ",resp.Header)
          body, _ := ioutil.ReadAll(resp.Body)
          fmt.Println("response body : ",string(body))
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
