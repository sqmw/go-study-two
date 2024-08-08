package wire1easy

import "fmt"

/// 定义三层套娃

type Message string

type Greater struct {
	Message Message
}

type Event struct {
	Greater Greater
}

func NewMessage() Message {
	return Message("hi i'm message")
}

func NewGreater(m Message) Greater {
	return Greater{m}
}

func NewEvent(g Greater) Event {
	return Event{g}
}

func (g *Greater) Greet() Message {
	return g.Message
}

func (e *Event) Trigger() {
	msg := e.Greater.Greet()
	fmt.Println(msg)
}
