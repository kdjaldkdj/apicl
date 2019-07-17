package main

import (
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

const usage = "Manage apiary"

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   "Fetch API information",
			Action:  nil,
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List APIs.",
			Action:  nil,
		},
		{
			Name:    "publish",
			Aliases: []string{"p"},
			Usage:   "Publish API",
			Action:  nil,
		},
	}
}
func main() {
	app := cli.NewApp()
	app.Name = "apicl"
	app.Usage = usage
	app.Commands = commands()
	app.Version = "0.1.0"
	app.Run(os.Args)
}
