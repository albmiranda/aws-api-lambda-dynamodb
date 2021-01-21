package main

import (
	"fmt"
)


func GetLocation(sattelites []SatelliteRequest) (x, y float32) {

	for index := 0; index < len(sattelites); index++ {
		fmt.Println("satellite ", sattelites[index].Name, " has location ", satellitesLocation[sattelites[index].Name], " and distancce ", sattelites[index].Distance)
	}

	return 604.3, -3.1
}