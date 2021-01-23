package main

import (
	"net/http"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func handlerGetEnemyPosition(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

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