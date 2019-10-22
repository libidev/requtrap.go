package http

import (
	"net/http"
	"strings"
	"github.com/libidev/requtrap.go/cli/config"
)

func EnableCors(w *http.ResponseWriter, conf config.ConfigCors) {
	origins := strings.Join(conf.Origins, ",")
	methods := strings.Join(conf.Methods, ",")

	(*w).Header().Set("Access-Control-Allow-Origin", origins)
	(*w).Header().Set("Access-Control-Allow-Methods", methods)
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}
