package satellite

// Data TODO: adicionar comentario
type Data struct {
	Name     string   `json:"name" validate:"required,name"`
	Distance float32  `json:"distance" validate:"required,distance"`
	Message  []string `json:"message" validate:"required,message"`
}

// Location TODO: adicionar comentario
type Location struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

var satellitesLocation = map[string]Location{
	"kenobi":    Location{-500.0, -200.0},
	"skywalker": Location{100.0, -100.0},
	"sato":      Location{500.0, 100.0},
}
