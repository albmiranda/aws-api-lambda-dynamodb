package main

import (
	"go-meli/internal/handler"

	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handlers(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if req.Path == "/topsecrkjkjet" {
		if req.HTTPMethod == "POST" {
			return handler.PostMultipleSatellites(req)
		}
	}

	if req.Resource == "/topsecret_split/{name}" {
		if req.HTTPMethod == "POST" {
			return handler.PostSingleSatellite(req)
		}
	}

	if req.Resource == "/topsecret_split" {
		if req.HTTPMethod == "GET" {
			return handler.GetShipData(req)
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
