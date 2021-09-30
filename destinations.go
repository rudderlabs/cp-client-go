package rudderclient

import (
	"encoding/json"
	"fmt"
	// "log"
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

// GetDestination - Returns destination
func (c *Client) GetDestination(destinationID string) (*Destination, error) {
	host := c.WorkspaceHost

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/destinations/%s", host.Url, destinationID), nil)
	if err != nil {
		return nil, err
	}

	body, err := host.doRequest(req)
	if err != nil {
		return nil, err
	}

	destination := Destination{}
	err = json.Unmarshal(body, &destination)
	if err != nil {
		return nil, err
	}

	return &destination, nil
}

// CreateDestination - Create new destination
func (c *Client) CreateDestination(destination Destination) (*Destination, error) {
	host := c.WorkspaceHost
	rb, err := json.Marshal(destination)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%sdestinations/", host.Url)
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

	return &apiResponse.Destination, nil
}

// UpdateDestination - Create new destination. 
func (c *Client) UpdateDestination(destinationId string, destination Destination) (*Destination, error) {
	host := c.WorkspaceHost
	rb, err := json.Marshal(destination)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%sdestinations/%s", host.Url, destinationId)
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
