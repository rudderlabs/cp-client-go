package rudderclient

import (
        "encoding/json"
        "fmt"
        "log"
        "net/http"
        "strings"
)

// GetDestinations - Returns list of destinations.
func (c *Client) GetDestinations() ([]Destination, error) {
        host := c.WorkspaceHost
        req, err := http.NewRequest("GET", fmt.Sprintf("%s/destinations", host.Url), nil)
        if err != nil {
                return nil, err
        }

        body, err := host.doRequest(req)
        if err != nil {
                return nil, err
        }

        destinations := []Destination{}
        err = json.Unmarshal(body, &destinations)
        if err != nil {
                return nil, err
        }

        return destinations, nil
}

// FilterDestinations - Returns list of destinations, filtered by search params.
func (c *Client) FilterDestinations(tYpe string, name string) ([]Destination, error) {
    host := c.WorkspaceHost
    req, err := http.NewRequest("GET", fmt.Sprintf("%s/destinations", host.Url), nil)
    if err != nil {
        return nil, err
    }

    // Add type and name ids from query.
    q := req.URL.Query()
    if (tYpe == ""){
        q.Add("type", tYpe)
    }
    if (name == ""){
        q.Add("name", name)
    }
    req.URL.RawQuery = q.Encode()

    body, err := host.doRequest(req)
    if err != nil {
        return nil, err
    }

    destinations := []Destination{}
    err = json.Unmarshal(body, &destinations)
    if err != nil {
        return nil, err
    }

    return destinations, nil
}

// GetDestination - Returns destination
func (c *Client) GetDestination(destinationID string) (*Destination, error) {
        host := c.WorkspaceHost
        url := fmt.Sprintf("%s/destinations/%s", host.Url, destinationID)

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

        destination := apiResponse.Destination
        return &destination, nil
}

// CreateDestination - Create new destination
func (c *Client) CreateDestination(destination Destination) (*Destination, error) {
        host := c.WorkspaceHost

        rb, err := json.Marshal(destination)
        if err != nil {
                return nil, err
        }

        url := fmt.Sprintf("%s/destinations", host.Url)
        bodySent := string(rb)
        req, err := http.NewRequest("POST", url, strings.NewReader(bodySent))
        if err != nil {
                log.Println("Destination createrequest creation failed URL=", url, "body=", bodySent)
                return nil, err
        }

        body, err := host.doRequest(req)
        if err != nil {
                log.Println("Destination creation failed URL=", url, "body=", bodySent)
                return nil, err
        }

        apiResponse := ApiResponseWrapper{}
        err = json.Unmarshal(body, &apiResponse)
        if err != nil {
                return nil, err
        }

        return &apiResponse.Destination, nil
}

// UpdateDestination - Update new destination. 
func (c *Client) UpdateDestination(destinationId string, destination Destination) (*Destination, error) {
        host := c.WorkspaceHost

        rb, err := json.Marshal(destination)
        if err != nil {
                return nil, err
        }

        url := fmt.Sprintf("%s/destinations/%s", host.Url, destinationId)
        bodySent := string(rb)
        req, err := http.NewRequest("PUT", url, strings.NewReader(bodySent))
        if err != nil {
                log.Println("Destination updaterequest creation failed URL=", url, "body=", bodySent)
                return nil, err
        }

        body, err := host.doRequest(req)
        if err != nil {
                log.Println("Destination updation failed URL=", url, "body=", bodySent)
                return nil, err
        }

        apiResponse := ApiResponseWrapper{}
        err = json.Unmarshal(body, &apiResponse)
        if err != nil {
                return nil, err
        }

        return &apiResponse.Destination, nil
}

// DeleteDestination - Delete existing destination
func (c *Client) DeleteDestination(destinationId string) error {
        host := c.WorkspaceHost

        req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/destinations/%d", host.Url, destinationId), nil)
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
