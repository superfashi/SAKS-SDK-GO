package entities

import (
	"regexp"
	"strings"
	"strconv"
	"log"
	"github.com/stianeikeland/go-rpio"
)

var (
	numberCode = []uint8{0x3f, 0x06, 0x5b, 0x4f, 0x66, 0x6d, 0x7d, 0x07, 0x7f, 0x6f, 0x00, 0x40}
	addressCode = []uint8{0xc0, 0xc1, 0xc2, 0xc3}
)

type DigitalDisplayTM1637 struct {
	IC      *IC_TM1637
	Numbers []string
	IsOn    bool
}

func NewDigitalDisplayTM1637(pins map[string]rpio.Pin, realTrue rpio.State) *DigitalDisplayTM1637 {
	return &DigitalDisplayTM1637{
		IC: &IC_TM1637{
			Pins:     pins,
			RealTrue: realTrue,
		},
	}
}

func (d *DigitalDisplayTM1637) SetNumbers(value string) {
	pattern, _ := regexp.Compile(`[-|#|\d]\.?`)
	matches := pattern.FindAllString(value, -1)
	d.Numbers = []string{}
	for _, i := range matches {
		d.Numbers = append(d.Numbers, i)
	}
}

func (d *DigitalDisplayTM1637) On() {
	d.IC.SetCommand(0x8f)
	d.IsOn = true
}

func (d *DigitalDisplayTM1637) Off() {
	d.IC.Clear()
	d.IsOn = false
}

func (d *DigitalDisplayTM1637) Show(str string) {
	d.SetNumbers(str)
	d.IC.SetCommand(0x44)
	lower := len(d.Numbers)
	if lower > 4 {
		lower = 4
	}
	for i := 0; i < lower; i++ {
		dp := strings.Count(d.Numbers[i], ".") > 0
		num := strings.Replace(d.Numbers[i], ".", "", -1)
		var after int
		switch num {
		case "#":
			after = 10
		case "-":
			after = 11
		default:
			var err error
			after, err = strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
		}
		if dp {
			d.IC.SetData(addressCode[i], numberCode[after] | 0x80)
		} else {
			d.IC.SetData(addressCode[i], numberCode[after])
		}
	}
	d.On()
}