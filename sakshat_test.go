package sakshat

import (
	"testing"
	"time"
)

func TestBuzzer(t *testing.T) {
	for x := 0; x < 6; x++ {
		BUZZER.Toggle()
		time.Sleep(time.Second / 5)
	}
}