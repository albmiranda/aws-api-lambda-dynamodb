// Package handler manages a specific HTTP method
package handler

import (
	"go-meli/internal/satellite"
	"go-meli/internal/db"
	internalHttp "go-meli/internal/http"

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

	data := DataRequest{
		Satellite: []satellite.Data{},
	}
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure to open request"),
		}, nil
	}

	v := validate.Struct(data)
	if ! v.Validate() {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure to open request - required field"),
		}, nil
	}

	err = db.UpdateMultipleSatellites(data.Satellite)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure to update database"),
		}, nil
	}

	x, y, decryptedMessage, err := satellite.FindShip(data.Satellite)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       string(""),
		}, nil
	}

	r := &internalHttp.DataResponse{
		Location: satellite.Location{X: x, Y: y},
		Message: decryptedMessage,
	}
	response, err := json.Marshal(r)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure to create response"),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(response),
	}, nil
}
