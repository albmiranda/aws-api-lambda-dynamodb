// Package satellite stores struct and methods which handle satellites information
package satellite

// Data represents satellite information
type Data struct {
	Name     string   `json:"name" validate:"required,name"`
	Distance float32  `json:"distance" validate:"required,distance"`
	Message  []string `json:"message" validate:"required,message"`
}

// Location represents ship/satellite point in a cartesian system
type Location struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

var satellitesLocation = map[string]Location{
	"kenobi":    Location{-500.0, -200.0},
	"skywalker": Location{100.0, -100.0},
	"sato":      Location{500.0, 100.0},
}
