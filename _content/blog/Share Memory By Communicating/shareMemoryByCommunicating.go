package main

import (
	"sync"
	"time"
)

/*
Program which polls a list of URLs
*/
func main() {
	// TODO: Fix how to invoke
	/*PollerThreadModel()
	PollerGoModel()*/
}

// 1. Thread model-based
// Resource Data structured in a traditional thread model
type Resource struct {
	url        string
	polling    bool
	lastPolled int64
}
type Resources struct {
	data []*Resource
	lock *sync.Mutex
}

// TODO: Comprehend
// Many would run in separated  thread -- such as?
func PollerThreadModel(res *Resources) {
	for {
		// get the least recently-polled Resource
		// and mark it as being polled
		res.lock.Lock()
		var r *Resource
		for _, v := range res.data {
			if v.polling {
				continue
			}
			if r == nil || v.lastPolled < r.lastPolled {
				r = v
			}
		}
		if r != nil {
			r.polling = true
		}
		res.lock.Unlock()
		if r == nil {
			continue
		}

		// poll the URL

		// update the Resource's polling and lastPolled
		res.lock.Lock()
		r.polling = false
		r.lastPolled = time.Duration.Nanoseconds()
		res.lock.Unlock()
	}
}

// TODO: Comprehend
// 2. Go model-based
/**
Resources polled -- from an -- input channel, -- are sent to -> output channel
*/
func PollerGoModel(in, out chan *Resource) {
	for r := range in {
		// poll the URL

		// send the processed Resource to out
		out <- r
	}
}
