package main

import (
	"log"

	"github.com/goburrow/modbus"
)

// Only Single Station
func NewRTUSingleStation() (err error) {

	var result []byte

	handler := modbus.NewRTUClientHandler(COM_ADDRESS)
	handler.BaudRate = baudRate
	handler.DataBits = dataBits
	handler.Parity = parity
	handler.StopBits = stopBits
	handler.SlaveId = station_1_ID
	handler.Timeout = timeoutMs

	handler.Connect()

	client := modbus.NewClient(handler)

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

	return nil
}
