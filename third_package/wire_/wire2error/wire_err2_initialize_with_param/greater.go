package greater3changeinjectsinnature

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

func NewMessage(phrase string) Message {
	if phrase != "" {
		/// 这里的 phrase 是显式的，需要使用 Message 来进行操作
		return Message(phrase)
	}
	/// 这里是一个 字面量，不需要转化
	return "hi i'm message"
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
