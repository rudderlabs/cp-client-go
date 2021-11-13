package rudderclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Auth kinds using in RudderStack APIs.
type AuthKind int

const (
	BasicAuth AuthKind = iota
	TokenAuth          = iota
)

// Client -
type Client struct {
	HTTPClient    *http.Client
	WorkspaceHost HostAccessStruct
	SchemaHost    HostAccessStruct
}

// HostAccessStruct -
type HostAccessStruct struct {
	HTTPClient *http.Client
	Url        string   `json:"hosturl"`
	Token      string   `json:"token"`
	AuthKind   AuthKind `json:"authKind"`
}

// NewClient -
func NewClient(workspaceHost, workspaceToken, schemaHost, schemaToken *string) (*Client, error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}
	c := Client{
		HTTPClient: httpClient,
		WorkspaceHost: HostAccessStruct{
			HTTPClient: httpClient,
			Url:        *workspaceHost,
			Token:      *workspaceToken,
			AuthKind:   TokenAuth,
		},
		SchemaHost: HostAccessStruct{
			HTTPClient: httpClient,
			Url:        *schemaHost,
			Token:      *schemaToken,
			AuthKind:   BasicAuth,
		},
	}

	return &c, nil
}

func (ha *HostAccessStruct) doRequest(req *http.Request) ([]byte, error) {
	if ha.AuthKind == BasicAuth {
		req.SetBasicAuth(ha.Token, "")
	} else {
		req.Header.Set("Authorization", "Bearer " + ha.Token)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	response, err := ha.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	statusOK := response.StatusCode >= 200 && response.StatusCode < 300
	if !statusOK {
		return nil, fmt.Errorf("status: %d, body: %s", response.StatusCode, body)
	}

	return body, err
}
