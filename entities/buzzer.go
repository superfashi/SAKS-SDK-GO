package entities

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

type Buzzer struct {
	Pin      rpio.Pin
	RealTrue rpio.State
	IsOn     bool
}

func NewBuzzer(pin rpio.Pin, realTrue rpio.State) *Buzzer {
	return &Buzzer{
		Pin:      pin,
		RealTrue: realTrue,
	}
}

func (b *Buzzer) On() {
	b.Pin.Write(b.RealTrue)
	b.IsOn = true
}

func (b *Buzzer) Off() {
	b.Pin.Write(b.RealTrue ^ 0x01)
	b.IsOn = false
}

func (b *Buzzer) Beep(t time.Duration) {
	b.On()
	time.Sleep(t)
	b.Off()
}

func (b *Buzzer) BeepAction(t, sleep time.Duration, times int) {
	for i := 0; i < times; i++ {
		b.Beep(t)
		time.Sleep(sleep)
	}
}