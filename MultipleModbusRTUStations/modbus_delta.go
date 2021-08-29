package main

const (
	ON  = 0xFF00
	OFF = 0x0000
)

// For Delta Modbus Registers Manual
// Refer : http://www.deltronics.ru/images/manual/DVP-ES2-EX2-SS2-SA2-SX2-SE-TP_PM_EN_20181030.pdf
// Chapter 4 Pages 10-14
// NOTE: You may use other PLCs, just have correct Register mappings
// Below registers must be marked active.

var MODBUS_DELTA map[string]map[int]uint16 = map[string]map[int]uint16{
	// Data Registers
	"D": map[int]uint16{
		23: uint16(0x1017),
	},
	// Coil registers: ON:0xFF00, OFF: 0x0000
	"M": map[int]uint16{
		1: uint16(0x0801),
	},
}
