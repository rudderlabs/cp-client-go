package rudderclient

import (
        "time"

        "github.com/hashicorp/terraform-plugin-framework/types"
)

// Each config property can be an arbitrary struct. Using interface{} for the same.
type SingleConfigPropertyValue interface {
}

// Source Definition -
type SourceDefinition struct {
        ID            types.String                             `json:"id,omitempty"`
        Name          types.String                             `json:"name"`
        Category      types.String                             `json:"category"`
        CreatedAt     time.Time                                `json:"createdAt"`
        UpdatedAt     time.Time                                `json:"updatedAt"`

        Config        map[string](SingleConfigPropertyValue)   `json:"config"`
}

// Destination Definition -
type DestinationDefinition struct {
        ID            types.String                             `json:"id,omitempty"`
        Name          types.String                             `json:"name"`
        Category      types.String                             `json:"category"`
        CreatedAt     time.Time                                `json:"createdAt"`
        UpdatedAt     time.Time                                `json:"updatedAt"`

        Config        map[string](SingleConfigPropertyValue)   `json:"config"`
}

// Sources -
type Source struct {
        ID            string                                   `json:"id,omitempty"`
        Name          string                                   `json:"name"`
        Type          string                                   `json:"type,omitempty"`
        IsEnabled     bool                                     `json:"enabled,omitempty"`
        IsDeleted     bool                                     `json:"deleted,omitempty"`
        CreatedAt     time.Time                                `json:"createdAt"`
        UpdatedAt     time.Time                                `json:"updatedAt"`

        Config        map[string](SingleConfigPropertyValue)   `json:"config"`
}

// Destinations -
type Destination struct {
        ID            string                                   `json:"id,omitempty"`
        Name          string                                   `json:"name"`
        Type          string                                   `json:"type,omitempty"`
        IsEnabled     bool                                     `json:"enabled,omitempty"`
        IsDeleted     bool                                     `json:"deleted,omitempty"`
        CreatedAt     time.Time                                `json:"createdAt"`
        UpdatedAt     time.Time                                `json:"updatedAt"`

        Config        map[string](SingleConfigPropertyValue)   `json:"config"`
}

// Connection between a source and a destinations -
type Connection struct {
        ID            string                                   `json:"id,omitempty"`
        SourceID      string                                   `json:"sourceId,omitempty"`
        DestinationID string                                   `json:"destinationId,omitempty"`
        IsEnabled     bool                                     `json:"enabled,omitempty"`
        IsDeleted     bool                                     `json:"deleted,omitempty"`
}

// API responses we handle are unmarshalled into this object first -
// Only one field is usually present in a typical API call.
type ApiResponseWrapper struct {
	Source        Source                                   `json:"source,omitempty"`
	Sources       []Source                                 `json:"sources,omitempty"`
	Destination   Destination                              `json:"destination,omitempty"`
	Destinations  []Destination                            `json:"destinations,omitempty"`
	Connection    Connection                               `json:"connection,omitempty"`
	Connections   []Connection                             `json:"connections,omitempty"`
}
