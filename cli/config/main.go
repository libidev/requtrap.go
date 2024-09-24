package config

import (
	"os"

	"github.com/libidev/requtrap.go/cli/errors"
	"gopkg.in/yaml.v3"
)

// Service Struct
type Service struct {
	Path            string `yaml:"path"`
	Upstream        string `yaml:"upstream"`
	MaxIdleConn     int    `yaml:"max_idle_conn"`
	IdleConnTimeout string `yaml:"idle_conn_timeout"`
	Cors            Cors   `yaml:"cors"`
}

// Cors Struct
type Cors struct {
	Enable        bool     `yaml:"enable"`
	Methods       []string `yaml:"methods"`
	Origins       []string `yaml:"origins"`
	ExposeHeaders []string `yaml:"expose-headers"`
}

// Authentication Struct
type Authentication struct {
	Type     string `yaml:"type"`
	Upstream string `yaml:"upstream"`
}

// Yaml Struct
type Yaml struct {
	Name           string         `yaml:"name"`
	Host           string         `yaml:"host"`
	Port           int            `yaml:"port"`
	Services       []Service      `yaml:"services"`
	Authentication Authentication `yaml:"authentication"`
	Cors           Cors           `yaml:"cors"`
}

var (
	// Default config
	Default = `
    name: book-store
    host: 127.0.0.1
    port: 8080
    services:
      - path: /books
        upstream: http://127.0.0.1:8001
      - path: /authors
				upstream: http://127.0.0.1:8002
		cors:
			enable: true
			methods:
				- GET
				- POST
				- PUT
				- DELETE
			origins:
				- http://localhost:3000
	`

	config = Yaml{}
)

// Parse is config parser function
func Parse(confile string) (*Yaml, error) {
	var err error
	defer errors.IsError(err)

	f, err := os.ReadFile(confile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(f), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
