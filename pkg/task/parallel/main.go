package task

import (
	"sync"

	"github.com/forrestbthomas/curly-parakeet/pkg/task"
)

type Parallel struct {
	state int
}

func (p *Parallel) Get(s string) int {
	switch s {
	case "state":
		return p.state
	}
	return 0
}

func (p *Parallel) Set(s string, v int) {
	switch s {
	case "state":
		p.state = v
	}
}
func (p *Parallel) Generator(fn task.TaskWork) task.TaskDefinition {
	var wg sync.WaitGroup
	output := make(chan int)
	return func(input chan int) chan int {
		for val := range input {
			wg.Add(1)
			go func(v int, w *sync.WaitGroup) {
				fn(v, output, p)
				wg.Done()
			}(val, &wg)
		}
		go func(w *sync.WaitGroup) {
			w.Wait()
			p.Set("state", 0)
		}(&wg)
		return output
	}
}
