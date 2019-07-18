package apiary_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/nephyer/apicl/apiary"
)

func TestList(t *testing.T) {

	// mock response
	var listResponse = `
{
	"apis": [
    {
      "apiName": "mock",
      "apiDocumentationUrl": "https://mock.docs.apiary.io/",
      "apiSubdomain": "mock",
      "apiIsPrivate": false,
      "apiIsPublic": true,
      "apiIsTeam": false,
      "apiIsPersonal": true
    }
  ]
} `
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := apiary.NewClient()
	client.BaseURL = server.URL
	os.Setenv("APIARY_API_KEY", "mockinapikey")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, listResponse)
	})
	apis, err := client.List()
	if err != nil {
		t.Errorf("client.List got %v\n err :%v\n", apis, err)
	}

	expected := &apiary.Response{
		APIS: []apiary.ResponseList{
			{
				APIName:             "mock",
				APIDocumentationURL: "https://mock.docs.apiary.io/",
				APISubdomain:        "mock",
				APIisPublic:         true,
				APIisPersonal:       true,
			},
		},
	}

	if !reflect.DeepEqual(apis, expected) {
		t.Errorf("client.List got %v expected %v\n", apis, expected)
	}

}
