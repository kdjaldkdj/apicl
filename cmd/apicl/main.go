package main

import (
	"fmt"
	"os"

	"github.com/nephyer/apicl"
	cli "gopkg.in/urfave/cli.v1"
)

const usage = "Manage apiary"

var version = "0.1.0"

// GitHash contains the hash of the buid.
var GitHash string

// GitTag contains the tag of build if defined.
var GitTag string

func main() {

	app := cli.NewApp()
	app.Name = "apicl"
	app.Usage = usage
	app.Commands = apicl.Commands()

	if GitHash != "" || GitTag != "" {
		version = fmt.Sprintf("%s-%s", GitHash, GitTag)
	}
	app.Version = version
	if !apicl.CheckConfigEnv() {
		fmt.Fprintf(os.Stderr, "Please, configure your APIARY_API_KEY before using this tool.\n")
		os.Exit(1)
	}
	app.Run(os.Args)
}
