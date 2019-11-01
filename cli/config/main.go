package config

import (
	"github.com/libidev/requtrap.go/cli/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// ConfigService Struct
type ConfigService struct {
	Path     string `yaml:"path"`
	Upstream string `yaml:"upstream"`
}

// ConfigCors Struct
type ConfigCors struct {
	Enable  bool     `yaml:"enable"`
	Methods []string `yaml:"methods"`
	Origins []string `yaml:"origins"`
}

// ConfigYaml Struct
type ConfigYaml struct {
	Name     string          `yaml:"name"`
	Host     string          `yaml:"host"`
	Port     int             `yaml:"port"`
	Services []ConfigService `yaml:"services"`
	Cors     ConfigCors      `yaml:"cors"`
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

	config = ConfigYaml{}
)

// Parse is config parser function
func Parse(confile string) (*ConfigYaml, error) {
	var err error
	defer errors.IsError(err)

	f, err := ioutil.ReadFile(confile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(f), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
