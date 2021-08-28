package main

import (
	"time"
)

// NOTE: This configurations can be input from user.
const (
	COM_ADDRESS  = "/dev/ttyS4"
	baudRate     = 9600
	dataBits     = 8
	parity       = "O"
	stopBits     = 1
	station_1_ID = byte(1)
	station_2_ID = byte(2)
	timeoutMs    = 200 * time.Millisecond
)
