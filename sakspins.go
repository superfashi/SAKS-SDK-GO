package sakshat

import (
	"github.com/stianeikeland/go-rpio"
)

const (
	IC_74HC595_DS   rpio.Pin = rpio.Pin(6)
	IC_74HC595_SHCP          = rpio.Pin(19)
	IC_74HC595_STCP          = rpio.Pin(13)

	IC_TM1637_DI  = rpio.Pin(25)
	IC_TM1637_CLK = rpio.Pin(5)

	BUZZER = rpio.Pin(12)

	TACT_RIGHT   = rpio.Pin(20)
	TACT_LEFT    = rpio.Pin(16)
	DIP_SWITCH_1 = rpio.Pin(21)
	DIP_SWITCH_2 = rpio.Pin(26)

	IR_SENDER   = rpio.Pin(17)
	IR_RECEIVER = rpio.Pin(9)
	DS18B20     = rpio.Pin(4)
	UART_TXD    = rpio.Pin(14)
	UART_RXD    = rpio.Pin(15)
	I2C_SDA     = rpio.Pin(2)
	I2C_SLC     = rpio.Pin(3)
)
