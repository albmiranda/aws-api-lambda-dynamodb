package main

import (
	"encoding/json"
	"net/http"
  
	"github.com/aws/aws-lambda-go/events"
)

func handlerMultipleSatellites(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	type DataRequest struct {
		Satellite []SatelliteData `json:"satellites"`
	}

	data := DataRequest{
		Satellite: []SatelliteData{},
	}
	
	err := json.Unmarshal([]byte(req.Body), &data)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body: string("Failure to open request"),
		  }, nil
	}

	// update dynamodb
	for i := range data.Satellite {
		err := CreateSatellite(data.Satellite[i])
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body: string("Failure to open request"),
			  }, nil
		}
	}

	x, y, found := GetLocation(data.Satellite)
	if !found {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body: string(""),
		  }, nil
	}
	msg := GetMessage(data.Satellite)

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