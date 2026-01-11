package main

import (
	"sync"
	"time"
)

// 试下一个支持并发读写，数据过期的map
type Map struct {
	m  map[int]value
	rw sync.RWMutex

	stopChan chan struct{}
	ttl      time.Duration
}

type value struct {
	v int
	t time.Time
}

func New(ttl time.Duration) *Map {
	m := Map{
		m:        map[int]value{},
		rw:       sync.RWMutex{},
		stopChan: make(chan struct{}),
		ttl:      ttl,
	}
	go m.StartClean()
	return &m
}

func (m *Map) Get(i int) (int, bool) {
	m.rw.RLock()
	v, ok := m.m[i]
	defer m.rw.RUnlock()

	if !ok {
		return 0, false
	}

	if time.Now().After(v.t) {
		return 0, false
	}
	return v.v, true
}

func (m *Map) Set(i, v int, t time.Duration) {
	m.rw.Lock()
	m.m[i] = value{v: v, t: time.Now().Add(t)}
	m.rw.Unlock()
}

func (m *Map) Del(i int) {
	m.rw.Lock()
	delete(m.m, i)
	m.rw.Unlock()
}

func (m *Map) Run() {
	m.rw.Lock()
	defer m.rw.Unlock()
	t := time.Now()
	for k, v := range m.m {
		if t.After(v.t) {
			m.Del(k)
		}
	}
}

func (m *Map) StartClean() {
	t := time.NewTicker(m.ttl)
	defer t.Stop()
	for {
		select {
		case <-m.stopChan:
			return
		case <-t.C:
			m.Run()
		}
	}
}

func (m *Map) Stop() {
	once := sync.Once{}
	once.Do(func() { close(m.stopChan) })
}
