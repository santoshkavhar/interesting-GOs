# Connecting to Multiple Modbus RTU Stations.

## Introduction
This code communicates only with Delta-PLC's multiple stations concurrently which have mentioned registers marked active. For using other PLCs (e.g Matsubishi) please refer manufacturer site and just add new modbus map.

## Issues Faced And Solutions
- *Permission Denied* : Run with sudo
- *input/output error* : Make sure "COM_ADDRESS" is correct. For Windows this will be like "COM3" and for Linux it will be "/dev/tty..."
- *Permission denied* for ComPort: give permission to the communication port by
`
sudo chmod +766 COM_ADDRESS
` 
command.
- *Windows User* : Make sure you have correct drivers and setting of parity, baud rate etc. Try Restarting once changes are made.
- *serial timeout* : Verify all the parameters provided to Station Handler. 
- *segmentation fault* : This will be cause of cross-compilation. Avoid cross-compilation.
- *Using USB-Convertor*: If you are using USB converter then COM_ADDRESS might be "/dev/ttyUSB0", find it using 
`
ll /sys/class/tty/ttyUSB*
`
command. If no output is shown that means something is wrong with connection.


## NOTE
Station is alias for Slave. 