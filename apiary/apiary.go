package apiary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"gopkg.in/urfave/cli.v1"
)

const (
	defaultbaseURL = "https://api.apiary.io/me/apis"
)

// Client does API calls Apiary.
type Client struct {
	// client performs HTTP calls.
	http *http.Client

	// BaseURL for the API.
	BaseURL string
}

// ResponseList holds information from a response listing APIs.
type ResponseList struct {
	APIName             string `json:"apiName"`
	APIDocumentationURL string `json:"apiDocumentationUrl"`
	APISubdomain        string `json:"ApiSubdomain"`
	APIisPublic         bool   `json:"apiIsPublic"`
	APIisPrivate        bool   `json:"apiIsPrivate"`
	APIisTeam           bool   `json:"apiIsTeam"`
	APIisPersonal       bool   `json:"apiIsPersonal"`
}

// NewClient returns a new client.
func NewClient() *Client {
	c := &Client{
		http: &http.Client{
			Timeout: time.Second * 10,
		},
		BaseURL: defaultbaseURL,
	}
	return c
}

// List gets a list of APIs for the current user.
func (c *Client) List(l *cli.Context) error {

	// Response holds the final answer
	type Response struct {
		APIS []ResponseList
	}

	r := Response{}
	apiaryKey := os.Getenv("APIARY_API_KEY")

	uri, err := url.Parse(c.BaseURL)
	if err != nil {
		return err
	}

	h := http.Header{}
	h.Set("Authorization", fmt.Sprintf("bearer %s", apiaryKey))

	req := &http.Request{
		URL:    uri,
		Method: http.MethodGet,
		Header: h,
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
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
