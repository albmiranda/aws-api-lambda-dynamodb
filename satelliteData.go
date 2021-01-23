package main


type SatelliteData struct {
	Name string `json:"name"`
	Distance float32 `json:"distance"`
	Message []string `json:"message"`
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
