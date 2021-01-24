package http

import (
	"go-meli/internal/satellite"
)

// DataResponse TODO: adicionar comentario
type DataResponse struct {
	Location satellite.Location `json:"position"`
	Message  string             `json:"messaje"`
}