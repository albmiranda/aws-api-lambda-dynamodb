package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)


func handlerSingleSatellite(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	type DataRequest struct {
		Distance float32 `json:"distance"`
		Message []string `json:"message"`
	}

	data := DataRequest{}
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: string("Failure to open request"),
		  }, nil
	}

	// read query parameter
	name, found := req.PathParameters["name"]
	if (!found ||(name != "sato" && name != "skywalker" && name != "kenobi")) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: string("Failure to open query"),
		  }, nil
	}

	var satellite = SatelliteData{
		Name: name,
		Distance: data.Distance,
		Message: data.Message,
	}

	// update dynamodb
	err = CreateSatellite(satellite)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: string("Failure to open request"),
			}, nil
	}

	// read all satellites from dynamodb
	satellites, err := GetSatellite()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: string("Failure read database"),
			}, nil
	}

	x, y, found := GetLocation(satellites)
	if !found {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body: string(""),
		  }, nil
	}
	msg := GetMessage(satellites)

	response := DataResponse{Location{X:x, Y:y},msg}
	responseJson, err := json.Marshal(&response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: string("Failure to create response"),
		  }, nil
	}

	return events.APIGatewayProxyResponse{
	  StatusCode: http.StatusOK,
	  Body: string(responseJson),
	}, nil
}