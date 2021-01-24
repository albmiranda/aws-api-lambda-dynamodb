// Package handler manages a specific HTTP method
package handler

import (
	"go-meli/internal/db"
	internalHttp "go-meli/internal/http"
	"go-meli/internal/satellite"

	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gookit/validate"
)

// PostMultipleSatellites receives all satellite data at once, updates then on database and tries to find the ship.
func PostMultipleSatellites(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	type DataRequest struct {
		Satellite []satellite.Data `json:"satellites" validate:"required"`
	}

	data := new(DataRequest)
	err := json.Unmarshal([]byte(req.Body), data)
	if err != nil {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	v := validate.Struct(data)
	if !v.Validate() {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	err = db.UpdateMultipleSatellites(data.Satellite)
	if err != nil {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	x, y, decryptedMessage, err := satellite.FindShip(data.Satellite)
	if err != nil {
		return internalHttp.ClientError(http.StatusNotFound)
	}

	r := &internalHttp.DataResponse{
		Location: satellite.Location{X: x, Y: y},
		Message:  decryptedMessage,
	}
	response, err := json.Marshal(r)
	if err != nil {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(response),
	}, nil
}
