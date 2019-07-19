package apicl

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

var fetchUsage = `
NAME:
   apicl fetch - Fetch API information

USAGE:
   apicl fetch [command options] [arguments...]

OPTIONS:
   --api-name value, -a value 
`

func fetch(l *cli.Context) error {
	apiName := l.String("api-name")
	if apiName == "" {
		fmt.Printf("%s", fetchUsage)
		os.Exit(1)
	}
	r, err := c.Fetch(apiName)
	if err != nil {
		return err
	}
	fmt.Println(r.Code)
	return nil
}
