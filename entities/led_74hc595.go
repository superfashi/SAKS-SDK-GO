package entities

import "github.com/stianeikeland/go-rpio"

type Led74HC595 struct {
	IC *IC_74HC595
}

func NewLed74HC595(pins map[string]rpio.Pin, realTrue rpio.State) *Led74HC595 {
	return &Led74HC595{
		IC: &IC_74HC595{
			Pins:     pins,
			RealTrue: realTrue,
		},
	}
}

func (d *Led74HC595) IsOn(index uint) bool {
	if index > 7 {
		return false
	}
	return (d.IC.Data >> index) & 0x01 != 0
}

func (d *Led74HC595) RowStatus() [8]bool {
	var r [8]bool
	for i := uint(0); i < 8; i++ {
		r[i] = d.IsOn(i)
	}
	return r
}

func (d *Led74HC595) On() {
	d.IC.SetData(0xff)
}

func (d *Led74HC595) Off() {
	d.IC.Clear()
}

func (d *Led74HC595) OnForIndex(index uint) {
	d.IC.SetData(d.IC.Data | (0x01 << index))
}

func (d *Led74HC595) OffForIndex(index uint) {
	arr := []uint8{0xfe, 0xfd, 0xfb, 0xf7, 0xef, 0xdf, 0xbf, 0x7f}
	d.IC.SetData(d.IC.Data & arr[index])
}

func (d *Led74HC595) SetRow(status [8]bool) {
	for i, stat := range status {
		if stat {
			d.OnForIndex(uint(i))
		} else {
			d.OffForIndex(uint(i))
		}
	}
}