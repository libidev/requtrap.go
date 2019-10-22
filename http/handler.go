package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/libidev/requtrap.go/cli/config"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type HttpHandler struct {
	Routes []config.ConfigService
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(h.GetRequestMethod(r))

	if r.URL.Path != "favicon.ico" {
		for _, service := range h.Routes {
			if r.URL.Path == service.Path {
				//GATEWAY
				tr := &http.Transport{
					MaxIdleConns:       10,
					IdleConnTimeout:    30 * time.Second,
					DisableCompression: true,
				}

				url := service.Upstream + service.Path
				client := &http.Client{Transport: tr}
				if h.GetRequestMethod(r) == "GET" {
					resp, err := client.Get(url)
					if err != nil {
						log.Fatal(err)
					} else {
						defer resp.Body.Close()

						if resp.StatusCode == http.StatusOK {
							contents, err := ioutil.ReadAll(resp.Body)
							if err != nil {
								log.Fatal(err)
							}

							var result []map[string]interface{}
							err = json.Unmarshal(contents, &result)
							if err != nil {
								log.Fatal(err)
							}
							fmt.Printf("\nredirect to : %s\n", service.Upstream)
							fmt.Println("response :")
							js, err := json.Marshal(result)
							fmt.Println(string(js))
							w.Header().Set("Content-Type", "application/json")
							w.Write(js)
						}
					}
				} else if h.GetRequestMethod(r) == "POST" {

					var jsonStr, err = ioutil.ReadAll(r.Body)
					if err != nil {
						log.Fatal(err)
					}

					req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
					req.Header.Set("X-Custom-Header", "myvalue")
					req.Header.Set("Content-type", "application/json")

					resp, err := client.Do(req)
					if err != nil {
						log.Fatal(err)
					}
					defer resp.Body.Close()

					fmt.Println("response status : ", resp.Status)
					fmt.Println("response Header : ", resp.Header)
					body, _ := ioutil.ReadAll(resp.Body)
					fmt.Println("response body : ", string(body))
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
