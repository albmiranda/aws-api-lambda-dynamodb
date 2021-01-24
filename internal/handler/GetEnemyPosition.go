package handler

import (
	"go-meli/internal/db"
	"go-meli/internal/satellite"
	localHttp "go-meli/pkg/http"

	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// GetEnemyPosition TODO: adicionar comentario
func GetEnemyPosition(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// read all satellites from dynamodb
	satellites, err := db.GetAllSatellites()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure read database"),
		}, nil
	}

	msg := satellite.GetMessage(satellites)
	x, y, found := satellite.GetLocation(satellites)
	if !found {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       string(""),
		}, nil
	}

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
