package main

import (
	"os"

	"github.com/libidev/requtrap.go/cli"
	"github.com/libidev/requtrap.go/cli/errors"
)

func main() {
	var err error
	defer errors.IsError(err)

	cli.Parse(os.Args[1:])

	if err != nil {
		return
	}
}
