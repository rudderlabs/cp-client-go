package rudderclient

import (
        "time"

        "github.com/hashicorp/terraform-plugin-framework/types"
)

// Source Definition -
type SourceDefinition struct {
        ID            types.String                             `json:"id"`
        Name          types.String                             `json:"name"`
        Category      types.String                             `json:"category"`
        CreatedAt     time.Time                                `json:"createdAt"`
        UpdatedAt     time.Time                                `json:"updatedAt"`

        Config        SourceDefinitionConfig                   `json:"config"`
}

// Destination Definition -
type DestinationDefinition struct {
        ID            types.String                             `json:"id"`
        Name          types.String                             `json:"name"`
        Category      types.String                             `json:"category"`
        CreatedAt     time.Time                                `json:"createdAt"`
        UpdatedAt     time.Time                                `json:"updatedAt"`

        Config        DestinationDefinitionConfig              `json:"config"`
}

// Sources -
type Source struct {
        ID            string                                   `json:"id"`
        Name          string                                   `json:"name"`
        Type          string                                   `json:"type"`
        CreatedAt     time.Time                                `json:"createdAt"`
        UpdatedAt     time.Time                                `json:"updatedAt"`

        Config        SourceConfig                             `json:"config"`
}

// Destinations -
type Destination struct {
        ID            string                                   `json:"id"`
        Name          string                                   `json:"name"`
        Type          string                                   `json:"type,omitempty"`
        CreatedAt     time.Time                                `json:"createdAt"`
        UpdatedAt     time.Time                                `json:"updatedAt"`

        Config        DestinationConfig                        `json:"config"`
}

// Connection between a source and a destinations -
type Connection struct {
        ID            string                                   `json:"id"`
        SourceID      string                                   `json:"sourceId"`
        DestinationID string                                   `json:"destinationId"`
}

type SourceConfig struct {
        ID            int                                      `json:"id"`
}

type DestinationConfig struct {
        ID            int                                      `json:"id"`
}

type SourceDefinitionConfig struct {
}

type DestinationDefinitionConfig struct {
}

// API responses we handle are unmarshalled into this object first -
// Only one field is usually present in a typical API call.
type ApiResponseWrapper struct {
	Source        Source                                   `json:"source,omitempty"`
	Sources       []Source                                 `json:"sources,omitempty"`
	Destination   Destination                              `json:"destination,omitempty"`
	Destinations  []Destination                            `json:"destinations,omitempty"`
	Connection    Connection                               `json:"connection,omitempty"`
}
