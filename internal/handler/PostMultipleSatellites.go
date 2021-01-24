package handler

import (
	"go-meli/internal/satellite"
	"go-meli/internal/db"
	localHttp "go-meli/pkg/http"

	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gookit/validate"
)

// PostMultipleSatellites TODO: adicionar comentario
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

	// update db
	err = db.UpdateMultipleSatellites(data.Satellite)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure to update database"),
		}, nil
	}

	x, y, found := satellite.GetLocation(data.Satellite)
	if !found {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       string(""),
		}, nil
	}
	msg := satellite.GetMessage(data.Satellite)

	response := localHttp.DataResponse{satellite.Location{X: x, Y: y}, msg}
	responseJSON, err := json.Marshal(&response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure to create response"),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(responseJSON),
	}, nil
}
