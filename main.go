package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"os"
)

var version string = "UNSPECIFIED"

func main() {
	app := cli.NewCLI("observe", version)
	app.Args = os.Args[1:]

	if len(app.Args) < 1 {
		panic("no URL specified")
	}

	fmt.Printf("Observing %s\n", app.Args[0])
}
