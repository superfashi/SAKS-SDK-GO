package entities

import (
	"github.com/stianeikeland/go-rpio"
)

type IC_TM1637 struct {
	Pins     map[string]rpio.Pin
	RealTrue rpio.State
}

func (t *IC_TM1637) StartBus() {
	t.Pins["clk"].Write(t.RealTrue)
	t.Pins["di"].Write(t.RealTrue)
	t.Pins["di"].Write(t.RealTrue ^ 0x01)
	t.Pins["clk"].Write(t.RealTrue ^ 0x01)
}

func (t *IC_TM1637) StopBus() {
	t.Pins["clk"].Write(t.RealTrue ^ 0x01)
	t.Pins["di"].Write(t.RealTrue ^ 0x01)
	t.Pins["clk"].Write(t.RealTrue)
	t.Pins["di"].Write(t.RealTrue)
}

func (t *IC_TM1637) SetBit(bit rpio.State) {
	t.Pins["clk"].Write(t.RealTrue ^ 0x01)
	t.Pins["di"].Write(bit)
	t.Pins["clk"].Write(t.RealTrue)
}

func (t *IC_TM1637) SetByte(data uint8) {
	for i := uint(0); i < 8; i++ {
		t.SetBit(rpio.State((data >> i) & 0x01))
	}
	t.Pins["clk"].Write(t.RealTrue ^ 0x01)
	t.Pins["di"].Write(t.RealTrue)
	t.Pins["clk"].Write(t.RealTrue)
}

func (t *IC_TM1637) SetCommand(command uint8) {
	t.StartBus()
	t.SetByte(command)
	t.StartBus()
}

func (t *IC_TM1637) SetData(address, data uint8) {
	t.StartBus()
	t.SetByte(address)
	t.SetByte(data)
	t.StartBus()
}

func (t *IC_TM1637) Clear() {
	t.SetCommand(0x80)
}