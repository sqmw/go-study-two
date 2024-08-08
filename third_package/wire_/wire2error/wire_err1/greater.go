package wire_err1

import (
	"errors"
	"fmt"
	"time"
)

/// 定义三层套娃

type Temper int

const (
	TemperGrumpy Temper = iota
)

type Message string

type Greater struct {
	Temper  Temper
	Message Message
}

type Event struct {
	Greater Greater
}

func NewMessage() Message {
	return Message("hi i'm message")
}

func NewGreater(m Message) Greater {
	greater := Greater{Message: m}
	if time.Now().Unix()%2 == 0 {
		greater.Temper = TemperGrumpy
	}
	return greater
}

func NewEvent(g Greater) (Event, error) {
	if g.Temper == TemperGrumpy {
		return Event{}, errors.New("event can't be create with grumpy temp")
	}
	return Event{g}, nil
}

func (g *Greater) Greet() Message {
	if g.Temper == TemperGrumpy {
		return Message("Go away!")
	}
	return g.Message
}

func (e *Event) Trigger() {
	msg := e.Greater.Greet()
	fmt.Println(msg)
}
