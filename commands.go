package apicl

import (
	"github.com/nephyer/apicl/apiary"
	"gopkg.in/urfave/cli.v1"
)

var c = apiary.NewClient()

// Commands returns actions for each of the commands:w
func Commands() []cli.Command {
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
			Action:  c.List,
		},
		{
			Name:    "publish",
			Aliases: []string{"p"},
			Usage:   "Publish API",
			Action:  nil,
		},
	}
}
