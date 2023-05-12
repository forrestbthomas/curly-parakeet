package task

import (
	"sync"

	"github.com/forrestbthomas/curly-parakeet/pkg/types"
)

type Task struct {
	Name string
}

func (t Task) GetName() string {
	return t.Name
}
func (t Task) Generator(fn types.TaskWork) types.TaskDefinition {
	var wg sync.WaitGroup
	output := make(chan int)
	return func(input chan int) chan int {
		for val := range input {
			wg.Add(1)
			go func(v int) {
				fn(v, output)
				wg.Done()
			}(val)
		}
		go func(w *sync.WaitGroup) {
			w.Wait()
			close(output)
		}(&wg)
		return output
	}
}
