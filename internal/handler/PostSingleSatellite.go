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

// PostSingleSatellite receives a single satellite data, updates it on database and tries to find the ship.
func PostSingleSatellite(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	type DataRequest struct {
		Distance float32  `json:"distance" validate:"required"`
		Message  []string `json:"message" validate:"required"`
	}

	data := DataRequest{}
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	v := validate.Struct(data)
	if ! v.Validate() {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	name, found := req.PathParameters["name"]
	if !found || (name != "sato" && name != "skywalker" && name != "kenobi") {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	var s = satellite.Data{
		Name:     name,
		Distance: data.Distance,
		Message:  data.Message,
	}
	err = db.UpdateSingleSatellite(s)
	if err != nil {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	satellites, err := db.GetAllSatellites()
	if err != nil {
		return internalHttp.ClientError(http.StatusBadRequest)
	}

	x, y, decryptedMessage, err := satellite.FindShip(satellites)
	if err != nil {
		return internalHttp.ClientError(http.StatusNotFound)
	}

	r := &internalHttp.DataResponse{
		Location: satellite.Location{X: x, Y: y},
		Message: decryptedMessage,
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
