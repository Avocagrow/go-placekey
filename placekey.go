// Package goplacekey provices bindings for interacting with the Placekey REST APIs.
package goplacekey

// Public constants
const (
	PlacekeyURL                    = "https://api.placekey.io"
	DefaultMaxNetworkRetries int64 = 2
)

type Client struct {
	Key string
}

type PlacekeyQuery struct {
	Lat        float64 `json:"latitude,omitempty"`
	Lng        float64 `json:"longitude,omitempty"`
	LocName    string  `json:"location_name,omitempty"`
	StreetAddr string  `json:"street_address,omitempty"`
	City       string  `json:"city,omitempty"`
	Region     string  `json:"region,omitempty"`
	PostalCode string  `json:"postal_code,omitempty"`
}

type PlacekeyOptions struct {
	StrictAddrMatch bool `json:"strict_addrss_match"`
	StrictNameMatch bool `json:"strict_name_match"`
}

type PlacekeyRequest struct {
	Query   PlacekeyQuery   `json:"query"`
	Options PlacekeyOptions `json:"options"`
}

type PlacekeyResponse struct {
	QueryID  string `json:"query_id,omitempty"`
	Placekey string `json:"placekey"`
}

func (c Client) GetByLatLng() {
}
