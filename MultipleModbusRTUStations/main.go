package main

import (
	"fmt"
)

func main() {

	// Connect according to choice
	err := ConnectToRTU()
	if err != nil {
		panic(err)
	}
}

func ConnectToRTU() (err error) {
	var choice string

	fmt.Println(`Enter Your choice:
	1. Connect to single station.
	2. Connect to 2 stations- The wrong way.
	3. Connect to 2 stations- The right way.`)

	fmt.Scanln(&choice)

	switch choice {
	case "1":
		_, err = ConnectRTUSingleStation(station_1_ID, nil)
	case "2":
		err = ConnectRTUMultipleStations_WrongWay()
	case "3":
		err = ConnectRTUMultipleStations_RightWay()
	default:
		return fmt.Errorf("please verify your choice")
	}
	return
}
