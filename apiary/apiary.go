package apiary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	defaultbaseURL = "https://api.apiary.io"
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

// Response holds the final answer
type Response struct {
	APIS []ResponseList
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
//func (c *Client) List(l *cli.Context) error {
func (c *Client) List() (*Response, error) {

	r := new(Response)
	apiaryKey := os.Getenv("APIARY_API_KEY")

	uri, err := url.Parse(c.BaseURL + "/me/apis")
	if err != nil {
		return nil, err
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
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// FetchResponse stores API information
type FetchResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// Fetch a specific  API.
func (c *Client) Fetch(name string) (*FetchResponse, error) {

	r := new(FetchResponse)
	apiaryKey := os.Getenv("APIARY_API_KEY")

	uri, err := url.Parse(c.BaseURL + "/blueprint/get/" + name)
	if err != nil {
		return nil, err
	}

	h := http.Header{}
	h.Set("Content-Type", "application/json; charset=utf-8")
	h.Set("Authentication", fmt.Sprintf("Token %s", apiaryKey))

	req := &http.Request{
		URL:    uri,
		Method: http.MethodGet,
		Header: h,
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return r, nil
}
