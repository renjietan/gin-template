package utility

import (
	"fmt"
	"sync"
	"time"
)

type TimeoutManager struct {
	timers map[string]*time.Timer
	mu     sync.Mutex
}

func NewTimeoutManager() *TimeoutManager {
	return &TimeoutManager{
		timers: make(map[string]*time.Timer),
	}
}

func (t *TimeoutManager) Set(name string, duration time.Duration, fn func()) {
	fmt.Printf("name-----------%s, duration=%v, fn=%+v\n", name, duration, fn)
	t.mu.Lock()
	if t.timers == nil {
		t.timers = make(map[string]*time.Timer)
	}
	// t.mu.Unlock()
	// t.mu.Lock()
	// if t_timer, exists := t.timers[name]; exists {
	// 	// 已存在同名定时器，停止它
	// 	res := t_timer.Stop()
	// 	fmt.Println("res===============", res)
	// 	delete(t.timers, name)
	// }
	t.StopByName(name)
	t.mu.Unlock()
	var timer *time.Timer
	timer = time.AfterFunc(duration, func() {
		defer func() {
			t.mu.Lock()
			fmt.Println("==============name", name)
			delete(t.timers, name)
			t.mu.Unlock()
		}()
		fn()
	})
	t.mu.Lock()
	t.timers[name] = timer
	t.mu.Unlock()
}

func (t *TimeoutManager) StopAll() {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, timer := range t.timers {
		timer.Stop()
	}
	t.timers = make(map[string]*time.Timer)
}

func (t *TimeoutManager) StopByName(name string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if timer, exists := t.timers[name]; exists {
		b := timer.Stop()
		delete(t.timers, name)
		return b
	}
	return false
}
