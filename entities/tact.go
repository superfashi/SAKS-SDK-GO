package entities

import (
	"github.com/stianeikeland/go-rpio"
	"time"
	"reflect"
)

type Tact struct {
	Pin       rpio.Pin
	RealTrue  rpio.State
	Status    bool
	observers []func(rpio.Pin, bool)
}

func NewTact(pin rpio.Pin, realTrue rpio.State) *Tact {
	ret := &Tact{
		Pin:      pin,
		RealTrue: realTrue,
	}
	ret.Update()
	go ret.Watching()
	return ret
}

func (t *Tact) Update() {
	t.Status = t.Pin.Read() == t.RealTrue
}

func (t *Tact) IsOn() bool {
	t.Update()
	return t.Status
}

func (t *Tact) RegisterContains(e func(rpio.Pin, bool)) int {
	for index, a := range t.observers {
		if &a == &e {
			return index
		}
	}
	return -1
}

func (t *Tact) Register(observer func(rpio.Pin, bool)) {
	if t.RegisterContains(observer) == -1 {
		t.observers = append(t.observers, observer)
	}
}

func (t *Tact) DeRegister(observer func(rpio.Pin, bool)) {
	pos := t.RegisterContains(observer)
	if pos != -1 {
		t.observers = append(t.observers[:pos], t.observers[pos + 1:]...)
	}
}

func (t *Tact) NotifyObservers() {
	for _, o := range t.observers {
		go o(t.Pin, t.Status)
	}
}

func (t *Tact) Watching() {
	lastStatus := t.Status
	for {
		t.Update()
		if t.Status != lastStatus {
			go t.NotifyObservers()
			lastStatus = t.Status
		}
		time.Sleep(50 * time.Millisecond)
	}
}

type TactRow struct {
	Tacts     []*Tact
	Pins      []rpio.Pin
	RealTrue  rpio.State
	observers []func([]bool)
}

func NewTactRow(pins []rpio.Pin, realTrue rpio.State) *TactRow {
	ret := &TactRow{
		Pins:     pins,
		RealTrue: realTrue,
	}
	for _, p := range pins {
		ret.Tacts = append(ret.Tacts, NewTact(p, realTrue))
	}
	go ret.Watching()
	return ret
}

func (t *TactRow) IsOn(index int) bool {
	if index >= len(t.Tacts) {
		return false
	}
	return t.Tacts[index].IsOn()
}

func (t *TactRow) RowStatus() []bool {
	var ret []bool
	for _, p := range t.Tacts {
		ret = append(ret, p.Status)
	}
	return ret
}

func (t *TactRow) RegisterContains(e func([]bool)) int {
	for index, a := range t.observers {
		if &a == &e {
			return index
		}
	}
	return -1
}

func (t *TactRow) Register(observer func([]bool)) {
	if t.RegisterContains(observer) == -1 {
		t.observers = append(t.observers, observer)
	}
}

func (t *TactRow) DeRegister(observer func([]bool)) {
	pos := t.RegisterContains(observer)
	if pos != -1 {
		t.observers = append(t.observers[:pos], t.observers[pos + 1:]...)
	}
}

func (t *TactRow) NotifyObservers() {
	for _, o := range t.observers {
		go o(t.RowStatus())
	}
}

func (t *TactRow) Watching() {
	lastStatus := t.RowStatus()
	for {
		nowStatus := t.RowStatus()
		if !reflect.DeepEqual(nowStatus, lastStatus) {
			go t.NotifyObservers()
			lastStatus = nowStatus
		}
		time.Sleep(50 * time.Millisecond)
	}
}