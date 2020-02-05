package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (h Handler) validateAuth(r *http.Request, url string) bool {
	if h.Auth.Type == "jwt" {
		for _, path := range h.Auth.Path {
			if path == url {
				jwt := r.Header.Get("token")
				url = h.Auth.Upstream

				tr := &http.Transport{
					MaxIdleConns:       10,
					IdleConnTimeout:    30 * time.Second,
					DisableCompression: true,
				}

				client := &http.Client{Transport: tr}
				req, _ := http.NewRequest("GET", url, nil)
				req.Header.Set("Token", jwt)

				resp, err := client.Do(req)
				if err != nil {
					//http.Error(w, err.Error(), http.StatusInternalServerError)
					return false
				}

				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(body)
				return true
			}
		}
	}
	return false
}
