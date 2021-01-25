// Package db which allows GET and PUT item (satellite.Data) into a database
package db

import (
	"go-meli/internal/satellite"
	database "go-meli/pkg/dynamodb"
)

// GetAllSatellites retrieves all item (satellite.Data) stored in database
func GetAllSatellites() ([]satellite.Data, error) {
	var data, err = database.Scan()
	if err != nil {
		return []satellite.Data{}, err
	}
	if len(data.Items) == 0 {
		return []satellite.Data{}, nil
	}

	var s []satellite.Data
	s, err = database.GetItemSatellite(data)

	return s, err
}

// UpdateSingleSatellite updates a item (satellite.Data) of database
func UpdateSingleSatellite(satellite satellite.Data) error {
	var data, err = database.NewItem(satellite)
	if err != nil {
		return err
	}

	err = database.PutItem(data)
	return err
}

// UpdateMultipleSatellites updates all available itens (satellite.Data) of database
func UpdateMultipleSatellites(satellites []satellite.Data) error {
	for _, s := range satellites {
		err := UpdateSingleSatellite(s)
		if err != nil {
			return err
		}
	}

	return nil
}
