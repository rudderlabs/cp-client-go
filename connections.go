package rudderclient

import (
    "encoding/json"
    "fmt"
    "log"
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

// FilterConnections - Returns list of connections, filtered by source and destination.
func (c *Client) FilterConnections(sourceId string, destinationId string) ([]Connection, error) {
    host := c.WorkspaceHost
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/connections", host.Url), nil)
    if err != nil {
        return nil, err
    }

    // Add source and destination ids from query.
    q := req.URL.Query()
    if (sourceId == ""){
        q.Add("sourceId", sourceId)
    }
    if (destinationId == ""){
        q.Add("destinationId", destinationId)
    }
    req.URL.RawQuery = q.Encode()

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
    url := fmt.Sprintf("%s/connections/%s", host.Url, connectionID)

    req, err := http.NewRequest("GET", url, nil)
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

    connection := apiResponse.Connection
    return &connection, nil
}

// CreateConnection - Create new connection. 
func (c *Client) CreateConnection(connection Connection) (*Connection, error) {
    host := c.WorkspaceHost
    rb, err := json.Marshal(connection)
    if err != nil {
        return nil, err
    }

    url := fmt.Sprintf("%s/connections/", host.Url)
    bodySent := string(rb)
    req, err := http.NewRequest("POST", url, strings.NewReader(bodySent))
    if err != nil {
	log.Println("Connection createrequest creation failed URL=", url, "body=", bodySent)
        return nil, err
    }

    body, err := host.doRequest(req)
    if err != nil {
	log.Println("Connection creation failed URL=", url, "body=", bodySent)
        return nil, err
    }

    apiResponse := ApiResponseWrapper{}
    err = json.Unmarshal(body, &apiResponse)
    if err != nil {
        return nil, err
    }

    return &apiResponse.Connection, nil
}

// UpdateConnection - Create new connection. 
func (c *Client) UpdateConnection(connectionId string, connection Connection) (*Connection, error) {
    host := c.WorkspaceHost
    rb, err := json.Marshal(connection)
    if err != nil {
        return nil, err
    }

    url := fmt.Sprintf("%s/connections/%s", host.Url, connectionId)
    req, err := http.NewRequest("PUT", url, strings.NewReader(string(rb)))
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
