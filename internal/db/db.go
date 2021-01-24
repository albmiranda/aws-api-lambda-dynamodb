package db

import (
	"go-meli/internal/satellite"
	database "go-meli/pkg/dynamodb"
)

// GetAllSatellites TODO: adicionar comentario
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

// UpdateSingleSatellite TODO: adicionar comentario
func UpdateSingleSatellite(satellite satellite.Data) error {
	var data, err = database.NewItem(satellite)
	if err != nil {
		return err
	}

	err = database.PutItem(data)
	return err
}


// UpdateMultipleSatellites TODO: adicionar comentario
func UpdateMultipleSatellites(satellites []satellite.Data) error {

	for _, s := range satellites {
		err := UpdateSingleSatellite(s)
		if err != nil {
			return err
		}
	}

	return nil
}