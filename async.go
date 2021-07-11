package main

import "sync"

func Parallel(fn func(sync.WaitGroup, chan interface{}, interface{}), vals []interface{}) chan interface{} {
	channel := make(chan interface{}, len(vals))

	var wg sync.WaitGroup
	wg.Add(len(vals))

	for _, val := range vals {
		go fn(wg, channel, val)
	}

	wg.Wait()

	return channel
}
