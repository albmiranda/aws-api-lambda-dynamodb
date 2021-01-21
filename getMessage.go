package main

import (
	"fmt"
)


func GetMessage(sattelites []SatelliteRequest) (msg string) {

	for index := 0; index < len(sattelites); index++ {
		fmt.Println("satellite ", sattelites[index].Name, " has message [", sattelites[index].Message, "]")
	}

	return string("vai corinthians!!!!")
}