package main

import (
	"os"

	"github.com/libidev/requtrap.go/cli"
)

func main() {

	cli.Parse(os.Args[1:])
}
