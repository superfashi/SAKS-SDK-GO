package sakshat

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"os"
	"os/signal"
	"github.com/hanbang-wang/SAKS-SDK-GO/entities"
)

var (
	Buzzer             *entities.Buzzer
	LEDRow             *entities.Led74HC595
	DigitalDisplay     *entities.DigitalDisplayTM1637
	Ds18b20            *entities.DS18B20
	TactRow, DipSwitch *entities.TactRow

	TactEventHandler      func(rpio.Pin, bool)
	DipSwitchEventHandler func([]bool)
)

func SaksGpioInit() {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}

	process := []rpio.Pin{IC_TM1637_DI, IC_TM1637_CLK, IC_74HC595_DS, IC_74HC595_SHCP, IC_74HC595_STCP}
	for p := range (process) {
		process[p].Output()
		process[p].Low()
	}

	process = []rpio.Pin{BUZZER, TACT_RIGHT, TACT_LEFT, DIP_SWITCH_1, DIP_SWITCH_2}
	for p := range (process) {
		process[p].Output()
		process[p].High()
	}

	process = []rpio.Pin{TACT_RIGHT, TACT_LEFT, DIP_SWITCH_1, DIP_SWITCH_2}
	for p := range (process) {
		process[p].Input()
		process[p].PullUp()
	}
}

func Clean() {
	LEDRow.Off()
	Buzzer.Off()
	DigitalDisplay.Off()
}

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			log.Println("Closing pins and terminating program...")
			Clean()
			rpio.Close()
			os.Exit(0)
		}
	}()
	SaksGpioInit()
	Buzzer = entities.NewBuzzer(BUZZER, rpio.Low)
	LEDRow = entities.NewLed74HC595(map[string]rpio.Pin{"ds": IC_74HC595_DS, "shcp": IC_74HC595_SHCP, "stcp": IC_74HC595_STCP}, rpio.High)
	Ds18b20 = entities.NewDS18B20(DS18B20)
	DigitalDisplay = entities.NewDigitalDisplayTM1637(map[string]rpio.Pin{"di": IC_TM1637_DI, "clk": IC_TM1637_CLK}, rpio.High)
	TactRow = entities.NewTactRow([]rpio.Pin{TACT_LEFT, TACT_RIGHT}, rpio.Low)
	for _, t := range(TactRow.Tacts) {
		t.Register(OnTactEvent)
	}
	DipSwitch = entities.NewTactRow([]rpio.Pin{DIP_SWITCH_1, DIP_SWITCH_2}, rpio.Low)
	DipSwitch.Register(OnDipSwitchEvent)
}

func OnTactEvent(pin rpio.Pin, status bool) {
	if TactEventHandler != nil {
		TactEventHandler(pin, status)
	}
}

func OnDipSwitchEvent(status []bool) {
	if DipSwitchEventHandler != nil {
		DipSwitchEventHandler(status)
	}
}