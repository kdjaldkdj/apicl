package apicl

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/nephyer/apicl/apiary"
	"gopkg.in/urfave/cli.v1"
)

var c = apiary.NewClient()

func list(l *cli.Context) error {
	r, err := c.List()
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 14, 7, 0, ' ', 0)
	header := "APINAME\tDOCS\tSUBDOMAIN\tPRIVATE\tPUBLIC\tTEAM\tPERSONAL"
	fmt.Fprintf(w, "%s\n", header)
	for _, api := range r.APIS {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			api.APIName,
			api.APIDocumentationURL,
			api.APISubdomain,
			strconv.FormatBool(api.APIisPrivate),
			strconv.FormatBool(api.APIisPublic),
			strconv.FormatBool(api.APIisTeam),
			strconv.FormatBool(api.APIisPersonal),
		)
	}
	w.Flush()
	return nil
}
