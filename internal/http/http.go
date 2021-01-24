// Package http exports a response information expected by contract
package http

import (
	"go-meli/internal/satellite"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// DataResponse represents the HTTP body response
type DataResponse struct {
	Location satellite.Location `json:"position"`
	Message  string             `json:"messaje"`
}

// ClientError returns an event response
func ClientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}