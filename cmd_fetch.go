package apicl

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func fetch(l *cli.Context) error {
	apiName := l.String("api-name")
	if apiName == "" {
		fmt.Println("--api-name must be provided. Please use --help")
		os.Exit(1)
	}
	r, err := c.Fetch(apiName)
	if err != nil {
		return err
	}
	fmt.Println(r.Code)
	return nil
}
