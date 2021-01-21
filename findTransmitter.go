package main

import (
	"net/http"
  
	"github.com/aws/aws-lambda-go/events"
)

func handleFindTransmitter(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	
	return events.APIGatewayProxyResponse{
	  StatusCode: http.StatusOK,
	  Body: string("teste2"),
	}, nil
}