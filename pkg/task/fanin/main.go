package task

import "github.com/forrestbthomas/curly-parakeet/pkg/types"

type Task struct {
	Name string
}

func (t Task) GetName() string {
	return t.Name
}

func (t Task) Generator(fn types.TaskWork) types.TaskDefinition {
	output := make(chan int, 1) // can buffer at 1, because it is a fan-in
	return func(input chan int) chan int {
		for val := range input {
			if len(output) == 0 {
				output <- val
				continue
			}
			fn(val, output)
		}
		return output
	}

}
