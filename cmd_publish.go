package apicl

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

func publish(l *cli.Context) error {
	filename := l.String("filename")
	apiname := l.String("api-name")
	r, err := c.Publish(apiname, filename)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
