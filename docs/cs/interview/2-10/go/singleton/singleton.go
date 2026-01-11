package main

import "sync"

func main() {
}

type client struct{}

var (
	c *client

	once sync.Once
	mu   sync.Mutex
)

// 饿汉模式
func init() {
	c = &client{}
}

// 懒汉模式 sync.Once
func Get() *client {
	once.Do(func() {
		c = &client{}
	})

	return c
}

// 双重检查
func Get2() *client {
	if c != nil {
		return c
	}
	mu.Lock()
	if c == nil {
		c = &client{}
	}
	mu.Unlock()
	return c
}
