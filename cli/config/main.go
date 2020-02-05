package config

import (
	"io/ioutil"

	"github.com/libidev/requtrap.go/cli/errors"
	"gopkg.in/yaml.v3"
)

// Service Struct
type Service struct {
	Path     string `yaml:"path"`
	Upstream string `yaml:"upstream"`
}

// Cors Struct
type Cors struct {
	Enable  bool     `yaml:"enable"`
	Methods []string `yaml:"methods"`
	Origins []string `yaml:"origins"`
}

// Authentication Struct
type Authentication struct {
	Type     string   `yaml:"type"`
	Upstream string   `yaml:"upstream"`
	Path     []string `yaml:"path"`
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
		authentication:
			type: jwt
			upstream: http://127.0.0.1:8000/auth/
			path:
				- http://127.0.0.1:8001/index
				- http://127.0.0.1:8002/home
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
	//fmt.Println(confile)

	f, err := ioutil.ReadFile(confile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
