// Package satellite stores struct and methods which handle satellites information
package satellite

import (
	"errors"
)

// FindShip tries to decrypt a message and find the ship based on information of all satellitees
func FindShip(satellites []Data) (x float32, y float32, message string, e error) {
	var found = true
	e = nil

	message = GetMessage(satellites)
	x, y, found = GetLocation(satellites)
	if !found {
		e = errors.New("That is no possible to find the ship")
	}

	return
}