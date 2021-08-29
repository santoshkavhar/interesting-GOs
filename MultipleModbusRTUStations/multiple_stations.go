package main

import (
	"github.com/goburrow/modbus"
)

// Wrong Way
// Different Transporter
func ConnectRTUMultipleStations_WrongWay() (err error) {

	_, err = ConnectRTUSingleStation(station_1_ID, nil)
	if err != nil {
		// error already logged
		return
	}
	_, err = ConnectRTUSingleStation(station_2_ID, nil)
	return
}

// Right Way
// Same Transporter
func ConnectRTUMultipleStations_RightWay() (err error) {

	var firstStationHandler *modbus.RTUClientHandler

	firstStationHandler, err = ConnectRTUSingleStation(station_1_ID, nil)
	if err != nil {
		return
	}

	// Connect this way to all other stations
	_, err = ConnectRTUSingleStation(station_2_ID, firstStationHandler)
	return
}
