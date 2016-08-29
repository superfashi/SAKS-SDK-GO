package sakshat

import (
	"testing"
	"time"
	"strconv"
)

func TestBuzzer(t *testing.T) {
	Buzzer.BeepAction(time.Second / 5, time.Second / 5, 2)
}

func TestLEDRow(t *testing.T) {
	// Test for Index
	for i := uint(0); i < 8; i++ {
		LEDRow.OnForIndex(i)
		time.Sleep(time.Second)
		LEDRow.OffForIndex(i)
	}
	time.Sleep(time.Second)

	// Test for SetRow
	LED1 := [8]bool{false, true, false, true, false, true, false, true}
	LED2 := [8]bool{true, false, true, false, true, false, true, false}
	LEDRow.SetRow(LED1)
	time.Sleep(time.Second)
	LEDRow.SetRow(LED2)
	time.Sleep(time.Second)

	// Test for status acquiring
	LED3 := LEDRow.RowStatus()
	t.Logf("The result of status acquiring: %t", LED2 == LED3)

	// Test for All
	LEDRow.On()
	time.Sleep(time.Second)
	LEDRow.Off()
}

func TestDigitalDisplay(t *testing.T) {
	// normal display test
	cases := []string{"0000", "5678", "-999", "1.2.3.4.", "12.34"}
	for i := range(cases) {
		DigitalDisplay.Show(cases[i])
		time.Sleep(time.Second)
	}

	DigitalDisplay.Off()
}

func TestTemperature(t *testing.T) {
	// Temperature read test
	for i := 0; i < 10; i++ {
		readTemp := Ds18b20.Temperature(0)
		t.Logf("Temperature #%d: %f", i, readTemp)
		DigitalDisplay.Show(strconv.FormatFloat(readTemp, 'g', 4, 64))
		time.Sleep(2 * time.Second)
	}
	DigitalDisplay.Off()
}