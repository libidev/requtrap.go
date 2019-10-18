package action

import (
	"fmt"
	"gopkg.in/yaml.v3"
  "github.com/libidev/requtrap.go/cli/config"
)

func Help() {
	fmt.Println("help")
}

func Start(conf *config.ConfigYaml) {
	fmt.Println("start")
	d, _ := yaml.Marshal(conf)
	fmt.Println(d)
}
