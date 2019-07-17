package main

import (
	"fmt"
	"os"

	"github.com/nephyer/apicl"
	cli "gopkg.in/urfave/cli.v1"
)

const usage = "Manage apiary"

func main() {

	app := cli.NewApp()
	app.Name = "apicl"
	app.Usage = usage
	app.Commands = apicl.Commands()
	app.Version = "0.1.0"
	if !apicl.CheckConfigEnv() {
		fmt.Fprintf(os.Stderr, "Please, configure your APIARY_API_KEY before using this tool.\n")
		os.Exit(1)
	}
	app.Run(os.Args)
}
