package main

import (
	"log"

	"github.com/goburrow/modbus"
)

// Only Single Station
func ConnectRTUSingleStation(stationID byte, oldHandler *modbus.RTUClientHandler) (thisHandler *modbus.RTUClientHandler, err error) {

	var result []byte
	var client modbus.Client

	handler := modbus.NewRTUClientHandler(COM_ADDRESS)
	handler.BaudRate = baudRate
	handler.DataBits = dataBits
	handler.Parity = parity
	handler.StopBits = stopBits
	handler.SlaveId = stationID
	handler.Timeout = timeoutMs

	handler.Connect()

	if oldHandler == nil {
		client = modbus.NewClient(handler)
	} else {
		client = modbus.NewClient2(handler, oldHandler)
	}

	// Switch On M1
	_, err = client.WriteSingleCoil(MODBUS_DELTA["M"][1], ON)
	if err != nil {
		log.Println("Couldn't Write to M-1 for station ", handler.SlaveId, err)
		return
	}

	// Write uint16(100) onto D23
	result, err = client.WriteSingleRegister(MODBUS_DELTA["D"][23], uint16(100))
	if err != nil {
		log.Println("Couldn't Write to D-23 for station ", handler.SlaveId, err)
		return
	}
	log.Println("Result from D-23", result)

	if oldHandler == nil {
		return handler, nil
	}
	return nil, nil
}
