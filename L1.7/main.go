package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SafeMap[Key comparable, Value any] struct {
	storage map[Key]Value
	mx      sync.RWMutex
}

func (targ *SafeMap[Key, Value]) NewMap() {
	targ.storage = make(map[Key]Value)
}

func (targ *SafeMap[Key, Value]) Set(key Key, value Value) {
	targ.mx.Lock()
	defer targ.mx.Unlock()
	targ.storage[key] = value
}

func (targ *SafeMap[Key, Value]) Get(key Key) Value {
	targ.mx.RLock()
	defer targ.mx.RUnlock()

	if value, hit := targ.storage[key]; hit {
		return value
	}
	var defaultValue Value
	return defaultValue
}

func main() {
	var concurrentMap SafeMap[string, int]
	concurrentMap.NewMap()

	var wg sync.WaitGroup
	ctx, close := context.WithCancel(context.Background())

	ConcurrentWrite := func(key string) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
				concurrentMap.Set(key, concurrentMap.Get(key)+1)
			}
		}

	}
	wg.Add(3)
	go ConcurrentWrite("1")
	go ConcurrentWrite("2")
	go ConcurrentWrite("3")

	time.Sleep(2 * time.Millisecond)
	close()
	wg.Wait()

	fmt.Println("1", concurrentMap.Get("1"))
	fmt.Println("2", concurrentMap.Get("2"))
	fmt.Println("3", concurrentMap.Get("3"))
}
