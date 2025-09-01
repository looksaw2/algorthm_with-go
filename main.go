package main

import (
	"fmt"
	"time"
)

type Event struct {
	Type    string
	Payload any
}

type EventLoop struct {
	Events chan Event
	done   chan bool
}

// 初始化事件循环
func NewEventLoop() *EventLoop {
	return &EventLoop{
		Events: make(chan Event, 100),
		done:   make(chan bool),
	}
}

// 启动事件循环
func (el *EventLoop) Start() {
	go func() {
		fmt.Printf("事件循环启动")
		for {
			select {
			case Event := <-el.Events:
				fmt.Printf("处理事件 %s , 数据%v\n", Event.Type, Event.Payload)
			case <-el.done:
				fmt.Printf("事件循环终止\n")
				return
			}
		}
	}()
}

// 终止事件循环
func (el *EventLoop) Stop() {
	el.done <- true
}

// 发布事件
func (el *EventLoop) Post(event Event) {
	el.Events <- event
}

// 循环
func main() {
	loop := NewEventLoop()
	loop.Start()
	//模拟事件发布
	go func() {
		for i := 0; i < 5; i++ {
			loop.Post(Event{
				Type:    "message",
				Payload: fmt.Sprintf("ID is %d", i),
			})
			time.Sleep(1 * time.Second)
		}
		loop.Stop()
	}()

	time.Sleep(6 * time.Second)
}
