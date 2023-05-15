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
	return func(input chan int) chan int {
		var wg sync.WaitGroup
		output := make(chan int, len(input))
		wg.Add(len(input))
		for val := range input {
			go func(v int, w *sync.WaitGroup) {
				fn(v, output, p)
				wg.Done()
			}(val, &wg)
		}
		wg.Wait()
		close(output)
		return output
	}
}
