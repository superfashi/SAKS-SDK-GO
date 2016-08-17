package sakshat

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
	"os"
	"os/signal"
)

func SaksGpioInit() {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}

	process := []*rpio.Pin{&IC_TM1637_DI, &IC_TM1637_CLK, &IC_74HC595_DS, &IC_74HC595_SHCP, &IC_74HC595_STCP}
	for p := range(process) {
		process[p].Output()
		process[p].Low()
	}

	process = []*rpio.Pin{&BUZZER, &TACT_RIGHT, &TACT_LEFT, &DIP_SWITCH_1, &DIP_SWITCH_2}
	for p := range(process) {
		process[p].Output()
		process[p].High()
	}

	process = []*rpio.Pin{&TACT_RIGHT, &TACT_LEFT, &DIP_SWITCH_1, &DIP_SWITCH_2}
	for p := range(process) {
		process[p].Input()
		process[p].PullUp()
	}
}

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("Closing pins and terminating program...")
			rpio.Close()
			os.Exit(0)
		}
	}()
	SaksGpioInit()
}