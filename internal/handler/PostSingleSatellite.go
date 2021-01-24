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

// PostSingleSatellite TODO: adicionar comentario
func PostSingleSatellite(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	type DataRequest struct {
		Distance float32  `json:"distance" validate:"required"`
		Message  []string `json:"message" validate:"required"`
	}

	data := DataRequest{}
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

	// read query parameter
	name, found := req.PathParameters["name"]
	if !found || (name != "sato" && name != "skywalker" && name != "kenobi") {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure to open path parameter"),
		}, nil
	}

	var sat = satellite.Data{
		Name:     name,
		Distance: data.Distance,
		Message:  data.Message,
	}

	// update db
	err = db.UpdateSingleSatellite(sat)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure to update database"),
		}, nil
	}

	// read all satellites from dynamodb
	satellites, err := db.GetAllSatellites()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string("Failure to read database"),
		}, nil
	}

	x, y, found := satellite.GetLocation(satellites)
	if !found {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       string(""),
		}, nil
	}
	msg := satellite.GetMessage(satellites)

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
