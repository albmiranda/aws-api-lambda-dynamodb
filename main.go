package main

import (
  "net/http"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

func handlers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  if req.Path == "/topsecret" {
    if req.HTTPMethod == "POST" {
      return handleFindTransmitter(req)
    }
  }

  return events.APIGatewayProxyResponse{
    StatusCode: http.StatusMethodNotAllowed,
    Body:       http.StatusText(http.StatusMethodNotAllowed),
  }, nil
}

func main() {
  lambda.Start(handlers)
}