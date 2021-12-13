package rudderclient

import (
    "encoding/json"
    "fmt"
    // "os"
    "log"
    "net/http"
    "strings"
)

// GetSources - Returns list of sources.
func (c *Client) GetSources() ([]Source, error) {
        host := c.WorkspaceHost
	url := fmt.Sprintf("%s/sources", host.Url)
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
                log.Println("UnMarshalling server response failed. Url=", url, "body=", string(body))
                return nil, err
        }

        return apiResponse.Sources, nil
}

// FilterSources - Returns list of sources, filtered by search params.
func (c *Client) FilterSources(tYpe string, name string) ([]Source, error) {
        host := c.WorkspaceHost
        url := fmt.Sprintf("%s/sources", host.Url)
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                return nil, err
        }

        // Add type and name ids from query.
        q := req.URL.Query()
        if (tYpe != ""){
                q.Add("type", tYpe)
        }
        if (name != ""){
                q.Add("name", name)
        }
        req.URL.RawQuery = q.Encode()

        body, err := host.doRequest(req)
        if err != nil {
                return nil, err
        }

        apiResponse := ApiResponseWrapper{}
        err = json.Unmarshal(body, &apiResponse)
        if err != nil {
                log.Println("UnMarshalling server response failed. Url=", url, "body=", string(body))
                return nil, err
        }

        return apiResponse.Sources, nil
}

// GetSource - Returns source
func (c *Client) GetSource(sourceID string) (*Source, error) {
        host := c.WorkspaceHost

        url := fmt.Sprintf("%s/sources/%s", host.Url, sourceID)
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
                log.Println("UnMarshalling server response failed. Url=", url, "body=", string(body))
                return nil, err
        }

        source := apiResponse.Source
        return &source, nil
}

// CreateSource - Create new source
func (c *Client) CreateSource(source Source) (*Source, error) {
        host := c.WorkspaceHost
        rb, err := json.Marshal(source)
        if err != nil {
                return nil, err
        }

        url := fmt.Sprintf("%s/sources", host.Url)
        bodySent := string(rb)
        req, err := http.NewRequest("POST", url, strings.NewReader(bodySent))
        if err != nil {
        log.Println("Source createrequest creation failed URL=", url, "body=", bodySent)
                return nil, err
        }

        body, err := host.doRequest(req)
        if err != nil {
        log.Println("Source creation failed URL=", url, "body=", bodySent)
                return nil, err
        }

        apiResponse := ApiResponseWrapper{}
        err = json.Unmarshal(body, &apiResponse)
        if err != nil {
                log.Println("UnMarshalling server response failed. Url=", url, "body=", string(body))
                return nil, err
        }

        return &apiResponse.Source, nil
}

// UpdateSource - Update new source. 
func (c *Client) UpdateSource(sourceId string, source Source) (*Source, error) {
        host := c.WorkspaceHost
        rb, err := json.Marshal(source)
        if err != nil {
                return nil, err
        }

        url := fmt.Sprintf("%s/sources/%s", host.Url, sourceId)
        bodySent := string(rb)
        req, err := http.NewRequest("PUT", url, strings.NewReader(bodySent))
        if err != nil {
                log.Println("Source updaterequest creation failed URL=", url, "body=", bodySent)
                return nil, err
        }

        body, err := host.doRequest(req)
        if err != nil {
                log.Println("Source updation failed URL=", url, "body=", bodySent)
                return nil, err
        }

        apiResponse := ApiResponseWrapper{}
        err = json.Unmarshal(body, &apiResponse)
        if err != nil {
                log.Println("UnMarshalling server response failed. Url=", url, "body=", string(body))
                return nil, err
        }

        return &apiResponse.Source, nil
}

// DeleteSource - Delete existing source
func (c *Client) DeleteSource(sourceId string) error {
        host := c.WorkspaceHost

        req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/sources/%s", host.Url, sourceId), nil)
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
