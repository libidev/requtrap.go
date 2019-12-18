package cli

import (
	"fmt"
	"log"
  "github.com/libidev/requtrap.go/cli/action"
  "github.com/libidev/requtrap.go/cli/config"
)

//func isError(err error){
//  if err != nil {
//    log.Fatal("error: %v",err)
//  }
//}

func Parse(args []string){
	var err error
	defer isError(err)

	if len(args) >= 1 {
		if args[0] == "help" {
			action.Help()
		} else if args[0] == "start" {
			if len(args) == 2 {
				config, err := config.Parse(args[1])
				if err != nil {return}
				action.Start(config)
			} else {
				fmt.Println("Config file not specified")
			}
		}
	} else {
		action.Help()
	}
}
