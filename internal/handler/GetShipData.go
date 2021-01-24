// Package handler manages a specific HTTP method
package handler

import (
	"go-meli/internal/db"
	"go-meli/internal/satellite"
	internalHttp "go-meli/internal/http"

	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// GetShipData gets the ship position and decrypts the message received by satellites
func GetShipData(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	s, err := db.GetAllSatellites()
	if err != nil {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	x, y, decryptedMessage, err := satellite.FindShip(s)
	if err != nil {
		return internalHttp.ClientError(http.StatusNotFound)
	}

	r := &internalHttp.DataResponse{
		Location: satellite.Location{X: x, Y: y},
		Message: decryptedMessage,
	}
	response, err := json.Marshal(r)
	if err != nil {
		return internalHttp.ClientError(http.StatusInternalServerError)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(response),
	}, nil
}
