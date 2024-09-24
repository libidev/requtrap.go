package http

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/libidev/requtrap.go/cli/config"
	"github.com/libidev/requtrap.go/pkg/constants"
	"github.com/libidev/requtrap.go/pkg/utils"
)

// Handler - main HTTP handler
type Handler struct {
	Routes  map[string]config.Service
	Cors    config.Cors
	Circuit CircuitBreaker
}

// GetRequestMethod - to get request method, eg. GET, POST, PUT, etc
func (h Handler) GetRequestMethod(r *http.Request) string {
	return r.Method
}

// AddRoute - to add service routes into route list
func (h *Handler) AddRoute(service config.Service) {
	h.Routes[service.Path] = service
}

// ServeHTTP - entry point for HTTP handler
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Cors.Enable {
		EnableCors(&w, h.Cors)
	}

	if r.Method == http.MethodOptions {
		return
	}

	if r.URL.Path != "favicon.ico" {

		path := r.URL.Path

		segments := strings.Split(path, "/")
		if len(segments) == 0 {
			return
		}

		service, ok := h.Routes["/"+segments[1]]
		if !ok {
			http.Error(w, "Service not found", http.StatusNotFound)
			return
		}

		fn := h.Gateway(w, r)
		err := h.Circuit.Call(fn, service, utils.StringToTimeDuration(service.IdleConnTimeout))
		if err != nil && errors.Is(err, constants.ERR_CIRCUIT_OPEN) {
			http.Error(w, "Service is not available", http.StatusServiceUnavailable)
		}
	}
}

// Gateway - frowarding client request to all services
func (h Handler) Gateway(w http.ResponseWriter, r *http.Request) func(config.Service) error {
	var reqbody []byte
	var err error

	if h.GetRequestMethod(r) == http.MethodPost {
		reqbody, err = io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return func(service config.Service) error {
		url := service.Upstream + service.Path

		if service.MaxIdleConn == 0 {
			service.MaxIdleConn = 10
		}

		tr := &http.Transport{
			MaxIdleConns:       service.MaxIdleConn,
			IdleConnTimeout:    utils.StringToTimeDuration(service.IdleConnTimeout),
			DisableCompression: true,
		}
		client := &http.Client{Transport: tr}
		req, err := http.NewRequest(h.GetRequestMethod(r), url, bytes.NewBuffer(reqbody))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		req.Header.Set("Content-type", "application/json")

		var setCors = func(wg *sync.WaitGroup, r *http.Request, exposeHeaders []string) {
			defer wg.Done()
			for _, header := range exposeHeaders {
				value := r.Header.Get(header)
				req.Header.Set(header, value)
			}
		}

		var wg sync.WaitGroup
		wg.Add(2)

		//Set header request from expose-headers list
		go setCors(&wg, r, h.Cors.ExposeHeaders)
		go setCors(&wg, r, service.Cors.ExposeHeaders)

		wg.Wait()

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}
		body, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		fmt.Println("")
		fmt.Println("Status :", resp.Status)
		fmt.Println("Header :")
		for k, v := range resp.Header {
			//Set header respone result from expose-headers list
			w.Header().Set(k, v[0])
			fmt.Println("  ", k+":", v[0])
		}
		fmt.Println("Body   :", string(body))

		w.Header().Set("Content-Type", "application/json")
		w.Write(body)

		return nil
	}
}
