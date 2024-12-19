package main

import (
	//"net/http"

	//"github.com/gin-gonic/gin"
)

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Decrement() {
	c.count--
}

func (c Counter) GetCount() int {
	return c.count
}
