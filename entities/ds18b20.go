package entities

import (
	"strings"
	"strconv"
	"log"
	"path/filepath"
	"os"
	"bufio"
	"time"
	"github.com/stianeikeland/go-rpio"
)

type DS18B20 struct {
	Pin rpio.Pin
}

func NewDS18B20(pin rpio.Pin) *DS18B20 {
	return &DS18B20{Pin: pin}
}

func (d *DS18B20) GetDeviceFile(index int) string {
	baseDir := `/sys/bus/w1/devices/`
	result, err := filepath.Glob(baseDir + "28*")
	if err != nil {
		log.Fatal(err)
	}
	if len(result) <= index {
		return ""
	} else {
		return result[index] + "/w1_slave"
	}
}

func (d *DS18B20) ReadTempRaw(index int) []string {
	var ret []string
	df := d.GetDeviceFile(index)
	if df == "" {
		return ret
	}
	file, err := os.Open(df)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

func (d *DS18B20) ReadTemp(index int) float64 {
	var ret float64 = -128.0
	tr := d.ReadTempRaw(index)
	if len(tr) == 0 {
		return ret
	}
	lines := tr
	afterTrim := strings.TrimSpace(lines[0])
	for afterTrim[len(afterTrim) - 3:] != "YES" {
		time.Sleep(200 * time.Millisecond)
		tr := d.ReadTempRaw(index)
		if len(tr) == 0 {
			return ret
		}
		afterTrim = strings.TrimSpace(tr[0])
		lines = tr
	}
	equalsPos := strings.Index(lines[1], "t=")
	if equalsPos != -1 {
		tempString := lines[1][equalsPos + 2:]
		var err error
		ret, err = strconv.ParseFloat(tempString, 64)
		if err != nil {
			log.Fatal(err)
		}
		ret = ret / 1000.0
	}
	return ret
}

func (d *DS18B20) Temperature(index int) float64 {
	return d.ReadTemp(index)
}