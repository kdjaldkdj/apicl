package apicl

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

func fetch(l *cli.Context) error {
	apiName := l.String("api-name")
	r, err := c.Fetch(apiName)
	if err != nil {
		return err
	}
	fmt.Println(r.Code)
	return nil
}
