package main

import (
	"encoding/json"
	"net/http"
  
	"github.com/aws/aws-lambda-go/events"
)


type DataRequest struct {
	Satellite []SatelliteData `json:"satellites"`
}

type SatelliteData struct {
	Name string `json:"name"`
	Distance float32 `json:"distance"`
	Message []string `json:"message"`
}

type DataResponse struct {
	Location Location `json:"position"`
	Message string `json:"messaje"`
}

type Location struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

var satellitesLocation = map[string]Location {
	"kenobi": Location{-500.0, -200.0},
	"skywalker": Location{100.0, -100.0},
	"sato": Location{500.0, 100.0},
}


func handleFindTransmitter(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

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