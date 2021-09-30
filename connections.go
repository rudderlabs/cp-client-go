package rudderclient

import (
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
	"strings"
)

// GetConnections - Returns list of connections.
func (c *Client) GetConnections() ([]Connection, error) {
	host := c.WorkspaceHost
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/connections", host.Url), nil)
	if err != nil {
		return nil, err
	}

	body, err := host.doRequest(req)
	if err != nil {
		return nil, err
	}

	connections := []Connection{}
	err = json.Unmarshal(body, &connections)
	if err != nil {
		return nil, err
	}

	return connections, nil
}

// GetConnection - Returns connection
func (c *Client) GetConnection(connectionID string) (*Connection, error) {
	host := c.WorkspaceHost

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/connections/%s", host.Url, connectionID), nil)
	if err != nil {
		return nil, err
	}

	body, err := host.doRequest(req)
	if err != nil {
		return nil, err
	}

	connection := Connection{}
	err = json.Unmarshal(body, &connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

// CreateConnection - Create new connection. 
func (c *Client) CreateConnection(connection Connection) (*Connection, error) {
	host := c.WorkspaceHost
	rb, err := json.Marshal(connection)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%sconnections/", host.Url)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := host.doRequest(req)
	if err != nil {
		return nil, err
	}

	apiResponse := ApiResponseWrapper{}
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse.Connection, nil
}

// DeleteConnection - Delete existing connection
func (c *Client) DeleteConnection(connectionId string) error {
	host := c.WorkspaceHost

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/connections/%d", host.Url, connectionId), nil)
	if err != nil {
		return err
	}

	body, err := host.doRequest(req)
	_ = body
	if err != nil {
		return err
	}

	return nil
}
