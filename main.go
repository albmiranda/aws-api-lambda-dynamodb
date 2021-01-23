package main

import (
  "net/http"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

type DataResponse struct {
  Location Location `json:"position"`
  Message string `json:"messaje"`
}

func handlers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  if req.Path == "/topsecret" {
    if req.HTTPMethod == "POST" {
      return handlerMultipleSatellites(req)
    }
  }

  if req.Resource == "/topsecret_split/{name}" {
    if req.HTTPMethod == "POST" {
      return handlerSingleSatellite(req)
    }
  }

  if req.Resource == "/topsecret_split" {
    if req.HTTPMethod == "GET" {
      return handlerGetEnemyPosition(req)
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