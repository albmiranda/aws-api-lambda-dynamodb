// Package http exports a response information expected by contract
package http

import (
	"go-meli/internal/satellite"
)

// DataResponse represents the HTTP body response
type DataResponse struct {
	Location satellite.Location `json:"position"`
	Message  string             `json:"messaje"`
}