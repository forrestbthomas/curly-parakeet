package task

import (
	"log"
	"sync"

	"github.com/forrestbthomas/curly-parakeet/pkg/task"
)

type FanOut struct {
	state int
}

func (f *FanOut) Get(s string) int {
	switch s {
	case "state":
		return f.state
	}
	return 0
}

func (f *FanOut) Set(s string, v int) {
	switch s {
	case "state":
		f.state = v
	}
}

func (f *FanOut) Generator(fn task.TaskWork) task.TaskDefinition {
	var wg sync.WaitGroup
	output := make(chan int)
	return func(input chan int) chan int {
		if len(input) == 0 || len(input) > 1 {
			log.Fatal("Cannot FanOut with less than or more than 1 element: has ", len(input), " elements")
		}
		el := <-input
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			fn(el, output, f)
			w.Done()
		}(&wg)
		wg.Wait()
		return output
	}

}
