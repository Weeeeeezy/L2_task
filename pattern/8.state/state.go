package main

import "fmt"

type State interface {
	FirstHandler()
	SecondHandler()
}

type Context struct {
	currentState State
}

func (c *Context) SetState(state State) {
	c.currentState = state
}

func (c *Context) FirstHandler() {
	c.currentState.FirstHandler()
}

func (c *Context) SecondHandler() {
	c.currentState.SecondHandler()
}

type ConcreteStateA struct {
	ctx *Context
}

func (c *ConcreteStateA) FirstHandler() {
	fmt.Println("First handler of State A")
}

func (c *ConcreteStateA) SecondHandler() {
	fmt.Println("Second handler of State A")
	fmt.Println("Change object state")
	c.ctx.SetState(&ConcreteStateB{ctx: c.ctx})
}

type ConcreteStateB struct {
	ctx *Context
}

func (c *ConcreteStateB) FirstHandler() {
	fmt.Println("First handler of State B")
}

func (c *ConcreteStateB) SecondHandler() {
	fmt.Println("Second handler of State B")
	fmt.Println("Change object state")
	c.ctx.SetState(&ConcreteStateA{ctx: c.ctx})
}

func main() {
	ctx := new(Context)
	ctx.SetState(&ConcreteStateA{ctx: ctx})
	ctx.FirstHandler()
	ctx.SecondHandler()
	ctx.FirstHandler()
	ctx.SecondHandler()
	ctx.FirstHandler()

}